package service

import (
	"context"

	"log"

	pb "github.com/nemesidaa/thumbsYT/proto/gen/service"
)

func (s *GRPCServer) Load(ctx context.Context, req *pb.LoadRequest) (*pb.LoadResponse, error) {
	log.Printf("Received: %s\n", req)
	var data []byte
	cachedVal, _, err := s.Storage.Thumb().GetByID(ctx, req.VideoID)
	if cachedVal == nil || err != nil {
		log.Printf("Caching: %s\n", req)
		data, ctx, err := s.Loader.Load(req.VideoID, req.Resolution, ctx)
		if err != nil {
			log.Printf("Failed to load: %s", err)
			return &pb.LoadResponse{DBID: "nil", RawData: []byte{}, Status: StatusError}, err
		}
		log.Printf("Loaded: %s\n", req.VideoID)
		go s.Storage.Thumb().Save(ctx, req.VideoID, string(data), req.Resolution)
	} else {
		log.Printf("Cached: %s\n", req.VideoID)
		data = []byte(cachedVal.Data)
	}

	// ctx for scalability
	log.Printf("Done: %s\n", req.ServiceID)
	return &pb.LoadResponse{DBID: req.VideoID, RawData: data, Status: StatusSuccess}, nil
}
