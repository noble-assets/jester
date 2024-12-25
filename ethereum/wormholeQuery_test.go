package ethereum

import (
	"bytes"
	"context"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFetchVaa(t *testing.T) {
	ctx := context.TODO()

	var logBuffer bytes.Buffer
	testHandler := slog.NewJSONHandler(&logBuffer, nil)
	logger := slog.New(testHandler)

	wormholeApiUrl := "https://api.testnet.wormscan.io/v1/signed_vaa"
	res, err := fetchVaa(ctx, logger, wormholeApiUrl, uint16(10002), uint64(45756), "0x7B1bD7a6b4E61c2a123AC6BC2cbfC614437D0470", "")
	require.NoError(t, err)

	expectedVAA := "AQAAAAABAPel1AcBA57rIzaTw70Qqlta9SxhuBYByiTv3viGqwgfFq4Wfx/EN0Mb8D71aTIwBz36NUmI98Q2fCEQyFlFSqQAZ1vRXAAAAAAnEgAAAAAAAAAAAAAAAHsb16a05hwqEjrGvCy/xhRDfQRwAAAAAAAAsrwPAScUAAAAAAAAAAAAAAAAKcvx4HFm0xRGMHrgeZn6bRYiOZAAAADjmUX/EAAAAAAAAAAAAAAAABt64ZSyDFVbnZmcg190zc42pnp0AAAAAAAAAAAAAAAAG3rhlLIMVVudmZyDX3TNzjamenQAmwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABrAAAAAAAAAAAAAAAAlO0ORGvBexrFFbj2bBk6ZU0driQAWZlOVFQGAAAAAAAAJw8AAAAAAAAAAAAAAAAMlBrZTKSlLtrqvyA7Yb3RgHzuwAAAAAAAAAAAAAAAAJTtDkRrwXsaxRW49mwZOmVNHa4kJxQACAAAAO++YLGeAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAD0JAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAGpv1m2ycUAAAAAAAAAAAAAAAAlO0ORGvBexrFFbj2bBk6ZU0driQAAAAAAAAAAAAAAAB6ClOEd3b36UzDV0KXGssiF7DbgQAAAAAAAAAAAAAAAHoKU4R3dvfpTMNXQpcayyIXsNuBAAAAAAAAAAAAAAAAKcvx4HFm0xRGMHrgeZn6bRYiOZAA"
	require.Equal(t, expectedVAA, res.VaaBytes)
}
