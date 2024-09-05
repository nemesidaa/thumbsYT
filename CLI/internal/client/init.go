package client

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/nemesidaa/thumbsYT/CLI/internal/config"
	pb "github.com/nemesidaa/thumbsYT/proto/gen/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Closefunc func() error

type Client struct {
	pb.MainstreamClient

	Resolution string
	ServerAddr string
	Timeout    int

	ServiceID string
}

func NewClient(cfg *config.ClientConfig) (*Client, Closefunc) {
	addr := fmt.Sprintf("%s:%d", cfg.ServerHost, cfg.ServerPort)
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials())) // change if needed
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return &Client{
		MainstreamClient: pb.NewMainstreamClient(conn),
		ServerAddr:       addr,
		Timeout:          cfg.Timeout,
		Resolution:       cfg.Resolution,
		ServiceID:        uuid.New().String(),
	}, conn.Close
}

type Resolution string

var AllAvailableResolutions map[Resolution]struct{} = map[Resolution]struct{}{
	"sddefault": {},
	"hqdefault": {},
	"mqdefault": {},
	"lqdefault": {},
	"hq720":     {},
}
