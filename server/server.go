package server

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"sync"
	"time"

	"connectrpc.com/connect"
	"connectrpc.com/grpcreflect"
	queryv1 "github.com/noble-assets/jester/gen/query/v1"
	"github.com/noble-assets/jester/gen/query/v1/queryv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

var l *VaaList

type VaaList struct {
	mu   sync.Mutex
	list []string
}

func InitVaaList() *VaaList {
	l = &VaaList{}
	return l
}

func (v *VaaList) Add(item string) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.list = append(v.list, item)
}

func (v *VaaList) GetThenClearAll() []string {
	v.mu.Lock()
	defer v.mu.Unlock()
	copyList := make([]string, len(v.list))
	copy(copyList, v.list)
	v.list = nil // clear list
	return copyList
}

//

var _ queryv1connect.QueryServiceHandler = &JesterServer{}

type JesterServer struct{}

func (s *JesterServer) GetVaas(
	ctx context.Context,
	req *connect.Request[queryv1.GetVaasRequest],
) (*connect.Response[queryv1.GetVaasResponse], error) {
	vaas := l.GetThenClearAll()
	res := connect.NewResponse(&queryv1.GetVaasResponse{
		Dollar: &queryv1.Dollar{
			Vaas: vaas,
		},
	})
	return res, nil
}

// StartServer initializes and starts the HTTP/2 gRPC server
func StartServer(ctx context.Context, log *slog.Logger, serverAddress string) {
	server := &JesterServer{}
	mux := http.NewServeMux()
	path, handler := queryv1connect.NewQueryServiceHandler(server)
	mux.Handle(path, handler)

	// enables reflection for gRPC support
	mux.Handle(grpcreflect.NewHandlerV1(
		grpcreflect.NewStaticReflector(queryv1connect.QueryServiceName),
	))

	srv := &http.Server{
		Addr:        serverAddress,
		Handler:     h2c.NewHandler(mux, &http2.Server{}),
		ReadTimeout: 45 * time.Second,
	}

	go func() {
		log.Info("started gRPC server", "address", serverAddress)
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
