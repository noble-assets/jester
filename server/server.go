// SPDX-License-Identifier: Apache-2.0
//
// Copyright 2025 NASD Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"time"

	"connectrpc.com/connect"
	"connectrpc.com/grpcreflect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"jester.noble.xyz/metrics"
	"jester.noble.xyz/state"

	api "jester.noble.xyz/api"
)

var _ api.QueryServiceHandler = &JesterGrpcServer{}

type JesterGrpcServer struct {
	log           *slog.Logger
	Metrics       *metrics.PrometheusMetrics
	mux           *http.ServeMux
	serverAddress string
	vaaList       *state.VaaList
}

func NewJesterGrpcServer(log *slog.Logger, m *metrics.PrometheusMetrics, mux *http.ServeMux, serverAddress string, vaaList *state.VaaList) *JesterGrpcServer {
	return &JesterGrpcServer{
		log:           log.With("server", "grpc"),
		Metrics:       m,
		mux:           mux,
		serverAddress: serverAddress,
		vaaList:       vaaList,
	}
}

// Start initializes and starts Jesters HTTP/2 gRPC server.
// This should be run in a goroutine.
func (s *JesterGrpcServer) StartServer(ctx context.Context) error {
	path, handler := api.NewQueryServiceHandler(s)
	s.mux.Handle(path, handler)

	// enables reflection for gRPC support
	s.mux.Handle(grpcreflect.NewHandlerV1(
		grpcreflect.NewStaticReflector(api.QueryServiceName),
	))

	srv := &http.Server{
		Addr:        s.serverAddress,
		Handler:     h2c.NewHandler(s.mux, &http2.Server{}),
		ErrorLog:    slog.NewLogLogger(s.log.Handler(), slog.LevelError),
		ReadTimeout: 45 * time.Second,
		BaseContext: func(net.Listener) context.Context {
			return ctx
		},
	}

	errChan := make(chan error, 1)
	go func() {
		s.log.Info("starting server", "address", s.serverAddress)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errChan <- err
			close(errChan)
		}
	}()

	select {
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := srv.Shutdown(shutdownCtx); err != nil {
			return fmt.Errorf("shutdown error: %w", err)
		}
		return nil
	case err := <-errChan:
		return fmt.Errorf("server error: %w", err)
	}
}

func (s *JesterGrpcServer) GetVoteExtension(
	ctx context.Context,
	req *connect.Request[api.GetVoteExtensionRequest],
) (*connect.Response[api.GetVoteExtensionResponse], error) {
	vaas := s.vaaList.GetThenClearAll()

	res := connect.NewResponse(&api.GetVoteExtensionResponse{
		Dollar: &api.Dollar{
			Vaas: vaas,
		},
	})

	s.Metrics.IncGetVoteExtensionCounter()

	return res, nil
}
