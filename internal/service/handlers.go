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
		log.Printf("Caching: %s\n", req.VideoID)
		data, _, err = s.Loader.Load(req.VideoID, req.Resolution, ctx)
		if err != nil {
			//log.Printf("Failed to load: %s", err)
			return &pb.LoadResponse{DBID: "nil", RawData: []byte{}, Status: StatusError}, err
		}

		_, _, err = store.Thumb().Save(ctx, req.VideoID, data, req.Resolution)
		if err != nil {
			return &pb.LoadResponse{DBID: "nil", RawData: []byte{}, Status: StatusError}, err
		}
	} else {
		log.Printf("Cached: %s\n", req.VideoID)
		log.Printf("Data: %d\n", len(cachedVal.Data))
		data = []byte(cachedVal.Data)
	}

	// ctx for scalability
	log.Printf("Loaded file with length: %d\n", len(data))

	log.Printf("Done: %s\n", req.ServiceID)
	return &pb.LoadResponse{DBID: req.VideoID, RawData: data, Status: StatusSuccess}, nil
}
