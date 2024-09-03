package txbroker

import (
	"context"

	"github.com/nemesidaa/thumbsYT/internal/loader"
	"github.com/nemesidaa/thumbsYT/internal/storage/model"
	pb "github.com/nemesidaa/thumbsYT/proto/gen/service"
)

// func
func (s *TxLoadLazyStack) CreateTask(ProcCtx context.Context, RequestID string) {
	s.UndoneChan <- TxLoadLazyStackComponent{
		ProcCtx:        ProcCtx,
		RequestID:      RequestID,
		RetriesCounter: 0,
	}
}

func (s *TxLoadLazyStack) ManageTask() (*pb.LoadResponse, error) {
	current := <-s.UndoneChan
	if current.RetriesCounter < s.MaxRetriesCounter {
		if current.ProcCtx.Value(loader.LoaderDone{}) != "" {
			thumb, ctx, err := s.Storage.Thumb().GetByID(context.WithValue(context.Background(), loader.LoaderDone{}, current.RequestID), current.RequestID)
			v := ctx.Value(model.FailedThumbTX{}).(string)
			if err != nil {
				s.UndoneChan <- current
				current.RetriesCounter++
			}
			if s.IdealCaching && v != "" {
				s.Storage.Thumb().Save(context.WithValue(context.Background(), loader.LoaderDone{}, current.RequestID), current.RequestID, "hqdefault")
			}
		}
	} else {
		return &pb.LoadResponse{DBID: "", RawData: "", Status: StatusFatal}, nil
	}
}
