package service

import (
	"context"

	"log"

	pb "github.com/nemesidaa/thumbsYT/proto/gen/service"
)

func (s *GRPCServer) Load(ctx context.Context, req *pb.LoadRequest) (*pb.LoadResponse, error) {
	log.Printf("Received: %s\n", req)
	var data []byte
	store, closefunc, err := s.ConnStorage()
	if err != nil {
		return &pb.LoadResponse{DBID: "nil", RawData: []byte{}, Status: StatusError}, err
	}
	defer closefunc()
	cachedVal, _, err := store.Thumb().GetByID(ctx, req.VideoID)
	if cachedVal == nil || err != nil {
		log.Printf("Caching: %s\n", req)
		data, _, err := s.Loader.Load(req.VideoID, req.Resolution, ctx)
		if err != nil {
			//log.Printf("Failed to load: %s", err)
			return &pb.LoadResponse{DBID: "nil", RawData: []byte{}, Status: StatusError}, err
		}
		log.Printf("Loaded: %s\n", data[:100])
		go store.Thumb().Save(ctx, req.VideoID, string(data), req.Resolution)
	} else {
		log.Printf("Cached: %s\n", req.VideoID)
		data = []byte(cachedVal.Data)
	}

	// ctx for scalability
	log.Printf("Done: %s\n", req.ServiceID)
	return &pb.LoadResponse{DBID: req.VideoID, RawData: data, Status: StatusSuccess}, nil
}
