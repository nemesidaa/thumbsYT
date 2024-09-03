package service

import (
	"context"

	"log"

	pb "github.com/nemesidaa/thumbsYT/proto/gen/service"
)

func (s *GRPCServer) Load(ctx context.Context, req *pb.LoadRequest) (*pb.LoadResponse, error) {
	data, ctx, err := s.Loader.Load(req.VideoID, req.Resolution, ctx)
	if err != nil {
		log.Printf("Failed to load: %s", err)
		return &pb.LoadResponse{DBID: "nil", RawData: "nil", Status: StatusError}, err
	}

	_, _, err = s.Storage.Thumb().Save(ctx, req.VideoID, string(data), req.Resolution)
	if err != nil {
		log.Printf("Failed to cache: %s", err)
		return &pb.LoadResponse{DBID: "nil", RawData: "nil", Status: StatusError}, err
	}
	// ctx for scalability
	log.Printf("Loaded: %s", req.VideoID)
	return &pb.LoadResponse{DBID: req.VideoID, RawData: string(data), Status: StatusSuccess}, nil
}
