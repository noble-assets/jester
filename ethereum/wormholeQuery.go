package ethereum

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/avast/retry-go/v4"
	"jester.noble.xyz/server"
)

type QueryData struct {
	WormHoleChainID uint16
	Emitter         string
	Sequence        uint64
	txHash          string // logging purposes only
}

type WormholeResp struct {
	VaaBytes []byte `json:"vaaBytes"`
}

var (
	errTooManyRequests = fmt.Errorf("%d:%s", http.StatusTooManyRequests, http.StatusText(http.StatusTooManyRequests))
	errServerError     = fmt.Errorf(">= %d:%s", http.StatusInternalServerError, "server error")
	errNotFound        = fmt.Errorf("%d:%s - waiting for wormhole to pickup tx", http.StatusNotFound, http.StatusText(http.StatusNotFound))
)

// StartQueryWorker starts a worker to query Wormhole VAA's.
// Once found, the VAA is added to the vaaList which is queryable
// via GRPC
func StartQueryWorker(
	ctx context.Context, log *slog.Logger,
	wormholeApiUrl string,
	dequeued *QueryData,
	vaaList *server.VaaList,
) {
	resp, err := fetchVaa(ctx, log, wormholeApiUrl, dequeued.WormHoleChainID, dequeued.Sequence, dequeued.Emitter, dequeued.txHash)
	if err != nil {
		log.Error("wormhole VAA query failed", "error", err)
	}

	log.Info("found VAA", "txHash", dequeued.txHash)
	vaaList.Add(resp.VaaBytes)
}

// fetchVaa sends a GET request with retry logic to the wormhole API
//
// `txHash` is used for logging purposes only
func fetchVaa(
	ctx context.Context, log *slog.Logger,
	wormholeApiUrl string,
	chainID uint16,
	seq uint64,
	emitter, txHash string,
) (WormholeResp, error) {
	chainIdStr := strconv.FormatUint(uint64(chainID), 10)
	seqStr := strconv.FormatUint(seq, 10)
	url := fmt.Sprintf("%s/%s/%s/%s", wormholeApiUrl, chainIdStr, emitter, seqStr)

	var wormholeResp WormholeResp

	fistAttempt := time.Now()
	retryAttemps := 50

	err := retry.Do(
		func() error {
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				return fmt.Errorf("failed to create request: %w", err)
			}

			req.Header.Set("accept", "application/json")

			client := &http.Client{Timeout: 10 * time.Second}
			resp, err := client.Do(req)
			if err != nil {
				return fmt.Errorf("request failed: %w", err)
			}
			defer resp.Body.Close()

			switch {
			case resp.StatusCode >= http.StatusInternalServerError:
				return errServerError
			case resp.StatusCode == http.StatusTooManyRequests:
				return errTooManyRequests
			case resp.StatusCode == http.StatusNotFound: // likely waiting for wormhole to pick up tx
				return errNotFound
			case resp.StatusCode != http.StatusOK:
				return retry.Unrecoverable(fmt.Errorf("non-retryable error: code %d, code: %s",
					resp.StatusCode, http.StatusText(resp.StatusCode),
				))
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return fmt.Errorf("failed to read wormhole response body: %w", err)
			}

			if err := json.Unmarshal(body, &wormholeResp); err != nil {
				return fmt.Errorf("failed to unmarshall wormholeResp json: %w", err)
			}

			return nil
		},
		retry.RetryIf(func(err error) bool {
			switch err {
			case errServerError, errTooManyRequests, errNotFound:
				return true
			default:
				return false
			}
		}),
		// adjust Attempts and Delay to ensure we don't give up querying
		// wormhole too soon
		retry.Attempts(uint(retryAttemps)),
		retry.Delay(30*time.Second),
		retry.Context(ctx),
		retry.DelayType(retry.FixedDelay),
		retry.Delay(30*time.Second),
		retry.OnRetry(func(attempt uint, err error) {
			since := time.Since(fistAttempt).Round(time.Second)
			log.Info("retry: VAA lookup", "attempt", fmt.Sprintf(
				"%d/%d", attempt+1, retryAttemps), "seq", seq, "error", err, "since-first-attempt", since, "txHash", txHash,
			)
		}),
	)
	if err != nil {
		return WormholeResp{}, err
	}

	return wormholeResp, nil
}
