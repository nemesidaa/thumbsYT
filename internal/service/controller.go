package service

import (
	"log"
	"net"

	"github.com/nemesidaa/thumbsYT/internal/config"
	pb "github.com/nemesidaa/thumbsYT/proto/gen/service"
	"google.golang.org/grpc"
)

func SafeExecution() {
	if r := recover(); r != nil {
		log.Println("Recovered from:", r)
	}
}

func ListenAndServe(cfg *config.ServerConfig) error {
	server := grpc.NewServer()
	logical, err := NewGRPCServer(cfg)
	if err != nil {
		return err
	}
	// defer logical.LazyBroker.Close()
	pb.RegisterMainstreamServer(server, logical)
	log.Println("Server Registered!")
	l, err := net.Listen("tcp", Addr(cfg.ServerHost, cfg.ServerPort))
	if err != nil {
		return err
	}
	log.Printf("Listening on %s", l.Addr().String())
	if err := server.Serve(l); err != nil {
		return err
	}

	return nil
}
