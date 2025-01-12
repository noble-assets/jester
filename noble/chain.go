package noble

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	nodev1beta1 "cosmossdk.io/api/cosmos/base/node/v1beta1"
	"github.com/avast/retry-go/v4"
	wormholev1 "github.com/noble-assets/wormhole/api/v1"
	"github.com/spf13/viper"
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

// InitializeNoble initializes a new Noble instance and verifies the gRPC connection.
// The intent behind this is to have this command run during cobras `PreRunE` or
// `PersistentPreRunE`.
// The returned *Noble pointer should be added to the app state.
func InitializeNoble(ctx context.Context, log *slog.Logger, gRPCurl string) (*Noble, error) {
	noble := newNoble(gRPCurl)

	if err := noble.newGrpcClient(); err != nil {
		return nil, err
	}

	noble.nodeServiceClient = noble.NewNodeServiceClient()
	noble.WormholeClient = noble.NewWormholeClient()

	if viper.GetBool(FlagSkipHealth) {
		log.Warn("skipping Noble gRPC health check")
		return noble, nil
	}

	err := retry.Do(func() error {
		ok, err := noble.HealthyGRPC(ctx)
		if !ok {
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
		return nil, fmt.Errorf("unhealthy Noble gRPC endpoint: %v", err)
	}

	log.Info("verified healthy Noble gRPC endpoint")

	return noble, nil
}

// newNoble initializes a new Noble instance
func newNoble(gRPCurl string) *Noble {
	return &Noble{
		Config: &Config{
			gRPCurl: gRPCurl,
		},
	}
}

// newGrpcClient creates a new gRPC client for the Noble instance
func (n *Noble) newGrpcClient() (err error) {
	n.gRPCClient, err = grpc.NewClient(n.Config.gRPCurl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to create Noble gRPC client: %w", err)
	}
	return nil
}

// CloseClients closes the Noble gRPC client connection
func (n *Noble) CloseClients() {
	if n.gRPCClient != nil {
		n.gRPCClient.Close()
	}
}

// NewNodeServiceClient creates a new NodeServiceClient for the Noble instance
func (n *Noble) NewNodeServiceClient() nodev1beta1.ServiceClient {
	return nodev1beta1.NewServiceClient(n.gRPCClient)
}

// NewWormholeClient creates a new WormholeClient for the Noble instance
func (n *Noble) NewWormholeClient() wormholev1.QueryClient {
	return wormholev1.NewQueryClient(n.gRPCClient)
}

// HealthyGRPC checks if the gRPC client is healthy by making a simple request.
// It returns true if the client is healthy, otherwise false.
func (n *Noble) HealthyGRPC(ctx context.Context) (bool, error) {
	_, err := n.nodeServiceClient.Status(ctx, &nodev1beta1.StatusRequest{})
	return err == nil, err
}
