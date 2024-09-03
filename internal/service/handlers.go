package service

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	pb "github.com/nemesidaa/thumbsYT/proto/gen/service"
)

func (s *GRPCServer) Give(ctx context.Context, req *pb.GiveRequest) (*pb.GiveResponse, error) {

	return &pb.GiveResponse{Status: StatusSuccess}, nil
}

func (s *GRPCServer) Load(ctx context.Context, req *pb.LoadRequest) (*pb.LoadResponse, error) {
	ctx, err := s.Loader.Load(req.VideoID, req.Resolution, ctx)
	if err != nil {
		log.Warnf("Failed to load: %s", err)
		return &pb.LoadResponse{DBID: StatusError}, err
	}

	return &pb.LoadResponse{DBID: "123"}, nil
}
