package noble

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	nodev1beta1 "cosmossdk.io/api/cosmos/base/node/v1beta1"
	"github.com/avast/retry-go/v4"
	wormholev1 "github.com/noble-assets/wormhole/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

// SkipHealth is a flag used for testing purposes only
const FlagSkipHealth = "skip-health"

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

	noble.nodeServiceClient = NewNodeServiceClient(noble.gRPCClient)
	noble.WormholeClient = NewWormholeClient(noble.gRPCClient)

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

// newGrpcClient creates a new gRPC client for the Noble instance
func newGrpcClient(url string) (*grpc.ClientConn, error) {
	gRPCClient, err := grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to create Noble gRPC client: %w", err)
	}
	return gRPCClient, nil
}

// NewNodeServiceClient creates a new NodeServiceClient for the Noble instance
func NewNodeServiceClient(cc *grpc.ClientConn) nodev1beta1.ServiceClient {
	return nodev1beta1.NewServiceClient(cc)
}

// NewWormholeClient creates a new WormholeClient for the Noble instance
func NewWormholeClient(cc grpc.ClientConnInterface) wormholev1.QueryClient {
	return wormholev1.NewQueryClient(cc)
}

// CloseClients closes the Noble gRPC client connection
func (n *Noble) CloseClients() {
	if n.gRPCClient != nil {
		n.gRPCClient.Close()
	}
}

// grpcHealthCheck checks the health of the Noble gRPC endpoint
func (n *Noble) grpcHealthCheck(ctx context.Context, log *slog.Logger) error {
	err := retry.Do(func() error {
		_, err := n.nodeServiceClient.Status(ctx, &nodev1beta1.StatusRequest{})
		if err != nil {
			return err
		}
		return nil
	},
		retry.Context(ctx),
		retry.Delay(10*time.Second),
		retry.DelayType(retry.FixedDelay),
		retry.OnRetry(func(attempt uint, err error) {
			log.Warn("noble gRPC health check failed. Retrying... (every 10 sec)", "attempt", fmt.Sprintf("%d/%d", attempt+1, 10), "err", err)
		}),
	)
	if err != nil {
		return fmt.Errorf("unhealthy Noble gRPC endpoint: %v", err)
	}

	return nil
}
