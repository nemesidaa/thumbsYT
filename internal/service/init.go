package service

import (
	"github.com/nemesidaa/thumbsYT/internal/config"
	"github.com/nemesidaa/thumbsYT/internal/loader"
	store "github.com/nemesidaa/thumbsYT/internal/storage/store/sqlstore"
	pb "github.com/nemesidaa/thumbsYT/proto/gen/service"
)

const (
	// DefaultPort is the default port for the service
	DefaultPort   = 5252
	StatusSuccess = "success"
	StatusError   = "error"
	StatusFatal   = "fatal"
)

type GRPCServer struct {
	pb.UnimplementedMainstreamServer
	// Thumbloader & YTcfg here
	Loader *loader.Loader

	Storage *store.Storage

	// LazyBroker *txbroker.TxLoadLazyStack
}

func NewGRPCServer(config *config.ServerConfig) *GRPCServer {
	db := store.NewStorage()
	return &GRPCServer{
		Loader:  loader.NewLoader(),
		Storage: db,

		// LazyBroker: txbroker.NewLazyBroker(config.BrokerCapacity, config.MaxBrokerRetriesCounter, config.IdealCaching, db),
	}
}
