package service

import (
	"net"

	pb "github.com/nemesidaa/thumbsYT/proto/gen/service"
	"google.golang.org/grpc"
)

func ListenAndServe(port int) error {
	server := grpc.NewServer()
	logical := NewGRPCServer()
	pb.RegisterMainstreamServer(server, logical)

	l, err := net.Listen("tcp", Port(port))
	if err != nil {
		return err
	}

	if err := server.Serve(l); err != nil {
		return err
	}

	return nil
}
