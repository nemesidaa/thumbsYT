package service

import (
	"github.com/nemesidaa/thumbsYT/internal/loader"
	store "github.com/nemesidaa/thumbsYT/internal/storage/store/sqlstore"
	pb "github.com/nemesidaa/thumbsYT/proto/gen/service"
)

const (
	// DefaultPort is the default port for the service
	DefaultPort   = 5252
	StatusSuccess = "success"
	StatusError   = "error"
)

type GRPCServer struct {
	pb.UnimplementedMainstreamServer
	// Thumbloader & YTcfg here
	Loader *loader.Loader

	DBCfg *store.Storage
}

func NewGRPCServer() *GRPCServer {

	return &GRPCServer{
		Loader: loader.NewLoader(),
		DBCfg:  store.NewStorage(),
	}
}
