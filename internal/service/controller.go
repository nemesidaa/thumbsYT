package service

import (
	"net"

	"github.com/nemesidaa/thumbsYT/internal/config"
	pb "github.com/nemesidaa/thumbsYT/proto/gen/service"
	"google.golang.org/grpc"
)

func ListenAndServe(cfg *config.ServerConfig) error {
	server := grpc.NewServer()
	logical := NewGRPCServer(cfg)
	// defer logical.LazyBroker.Close()
	pb.RegisterMainstreamServer(server, logical)

	l, err := net.Listen("tcp", Port(cfg.ServerPort))
	if err != nil {
		return err
	}

	if err := server.Serve(l); err != nil {
		return err
	}

	return nil
}
