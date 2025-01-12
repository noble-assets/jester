package server

import (
	"context"
	"encoding/hex"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"sync"
	"time"

	"connectrpc.com/connect"
	"connectrpc.com/grpcreflect"
	wormholev1 "github.com/noble-assets/wormhole/api/v1"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	api "jester.noble.xyz/api"
)

var l *VaaList

type VaaList struct {
	mu   sync.Mutex
	list [][]byte
}

func InitVaaList() *VaaList {
	l = &VaaList{}
	return l
}

func (v *VaaList) Add(item []byte) {
	if item == nil {
		return
	}
	v.mu.Lock()
	defer v.mu.Unlock()
	v.list = append(v.list, item)
}

func (v *VaaList) GetThenClearAll() [][]byte {
	v.mu.Lock()
	defer v.mu.Unlock()
	copyList := make([][]byte, len(v.list))
	copy(copyList, v.list)
	v.list = nil // clear list
	return copyList
}

//

var _ api.QueryServiceHandler = &JesterServer{}

type JesterServer struct {
	log            *slog.Logger
	wormholeClient wormholev1.QueryClient
}

func (s *JesterServer) GetVoteExtension(
	ctx context.Context,
	req *connect.Request[api.GetVoteExtensionRequest],
) (*connect.Response[api.GetVoteExtensionResponse], error) {
	log := s.log
	vaas := l.GetThenClearAll()

	var vaasToKeep [][]byte
	for _, vaa := range vaas {
		gRPCRes, err := s.wormholeClient.ExecutedVAA(ctx, &wormholev1.QueryExecutedVAA{
			Input:     hex.EncodeToString(vaa),
			InputType: "",
		})
		if err != nil {
			log.Error("executedVaa query from Noble over gRPC", "error", err)
			continue
		}
		if !gRPCRes.Executed {
			vaasToKeep = append(vaasToKeep, vaa)
		}
	}

	res := connect.NewResponse(&api.GetVoteExtensionResponse{
		Dollar: &api.Dollar{
			Vaas: vaasToKeep,
		},
	})

	return res, nil
}

// StartServer initializes and starts Jesters HTTP/2 gRPC server
func StartServer(ctx context.Context, log *slog.Logger, serverAddress string, wormholeClient wormholev1.QueryClient) {
	server := &JesterServer{
		log:            log,
		wormholeClient: wormholeClient,
	}
	mux := http.NewServeMux()
	path, handler := api.NewQueryServiceHandler(server)
	mux.Handle(path, handler)

	// enables reflection for gRPC support
	mux.Handle(grpcreflect.NewHandlerV1(
		grpcreflect.NewStaticReflector(api.QueryServiceName),
	))

	srv := &http.Server{
		Addr:        serverAddress,
		Handler:     h2c.NewHandler(mux, &http2.Server{}),
		ReadTimeout: 45 * time.Second,
	}

	go func() {
		log.Info("started Jester's gRPC server", "address", serverAddress)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("server error", "err", err)
			os.Exit(1)
		}
	}()

	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Error("server shutdown error", "error", err)
	}
}
