package server

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"connectrpc.com/connect"
	"connectrpc.com/grpcreflect"
	wormholev1 "github.com/noble-assets/wormhole/api/v1"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"jester.noble.xyz/state"

	api "jester.noble.xyz/api"
)

var _ api.QueryServiceHandler = &JesterGrpcServer{}

type JesterGrpcServer struct {
	log            *slog.Logger
	serverAddress  string
	vaaList        *state.VaaList
	wormholeClient wormholev1.QueryClient
}

func NewJesterGrpcServer(serverAddress string, log *slog.Logger, vaaList *state.VaaList, wormholeClient wormholev1.QueryClient) *JesterGrpcServer {
	return &JesterGrpcServer{
		log:            log,
		serverAddress:  serverAddress,
		vaaList:        vaaList,
		wormholeClient: wormholeClient,
	}
}

// Start initializes and starts Jesters HTTP/2 gRPC server.
// This should be run in a goroutine.
func (s *JesterGrpcServer) Start(ctx context.Context) {
	log := s.log

	mux := http.NewServeMux()
	path, handler := api.NewQueryServiceHandler(s)
	mux.Handle(path, handler)

	// enables reflection for gRPC support
	mux.Handle(grpcreflect.NewHandlerV1(
		grpcreflect.NewStaticReflector(api.QueryServiceName),
	))

	srv := &http.Server{
		Addr:        s.serverAddress,
		Handler:     h2c.NewHandler(mux, &http2.Server{}),
		ReadTimeout: 45 * time.Second,
	}

	go func() {
		log.Info("started Jester's gRPC server", "address", s.serverAddress)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(fmt.Sprintf("server error: %v", err))
		}
	}()

	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Error("server shutdown error", "error", err)
	}
}

func (s *JesterGrpcServer) GetVoteExtension(
	ctx context.Context,
	req *connect.Request[api.GetVoteExtensionRequest],
) (*connect.Response[api.GetVoteExtensionResponse], error) {
	log := s.log
	vaas := s.vaaList.GetThenClearAll()

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
