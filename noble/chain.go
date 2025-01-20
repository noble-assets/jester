package noble

import (
	"context"
	"log/slog"

	nodev1beta1 "cosmossdk.io/api/cosmos/base/node/v1beta1"
	wormholev1 "github.com/noble-assets/wormhole/api/v1"
	"google.golang.org/grpc"
)

type Noble struct {
	Config     *Config
	gRPCClient *grpc.ClientConn

	nodeServiceClient nodev1beta1.ServiceClient
	WormholeClient    wormholev1.QueryClient
}

type Config struct {
	gRPCurl string
}

// NewNoble initializes a new Noble instance and verifies the gRPC connection.
// The intent behind this is to have this command run during cobras `PreRunE` or
// `PersistentPreRunE`.
// The returned *Noble pointer should be added to the app state.
func NewNoble(ctx context.Context, log *slog.Logger, gRPCurl string, skipHealth bool) (noble *Noble, err error) {
	noble = &Noble{
		Config: &Config{
			gRPCurl: gRPCurl,
		},
	}

	noble.gRPCClient, err = newGrpcClient(gRPCurl)
	if err != nil {
		return nil, err
	}

	noble.nodeServiceClient = newNodeServiceClient(noble.gRPCClient)
	noble.WormholeClient = newWormholeClient(noble.gRPCClient)

	if skipHealth {
		log.Warn("skipping Noble gRPC health check")
		return noble, nil
	}

	if err = noble.grpcHealthCheck(ctx, log); err != nil {
		return nil, err
	}

	log.Info("verified healthy Noble gRPC endpoint")
	return noble, nil
}
