package service

import (
	"time"

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
	pb.MainstreamServer
	// Thumbloader & Cache here
	Loader *loader.Loader

	StorageString string

	ConnTimeout time.Duration
	// LazyBroker *txbroker.TxLoadLazyStack
}

func NewGRPCServer(config *config.ServerConfig) (*GRPCServer, error) {
	err := store.InitStorage(config.DBName)
	if err != nil {
		return nil, err
	}
	return &GRPCServer{
		Loader:        loader.NewLoader(config.Resolution),
		StorageString: config.DBName,
		ConnTimeout:   time.Duration(config.DefaultDBTimeout) * time.Second,
		// LazyBroker: txbroker.NewLazyBroker(config.BrokerCapacity, config.MaxBrokerRetriesCounter, config.IdealCaching, db),
	}, nil
}

func (s *GRPCServer) ConnStorage() (*store.Storage, func() error, error) {
	// // Handling ctx outside cause of rearming timeout:
	// var cancelFunc context.CancelFunc
	// if _, ok := ctx.Deadline(); !ok {
	// 	ctx, cancelFunc = context.WithDeadline(ctx, time.Now().Add(s.ConnTimeout))
	// } else {
	// 	cancelFunc = func() {}
	// }

	// DB connection:
	db, closefunc, err := store.NewConn(s.StorageString)
	if err != nil {
		// cancelFunc() // Canceling context in case of error
		return nil, nil, err
	}

	// Return in normal case:
	return db, closefunc, nil
}
