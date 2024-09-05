package loader

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

type LoaderDone struct{}
type LoaderFailed struct{}

func (l *Loader) Load(id, resolution string, ctx context.Context) ([]byte, context.Context, error) {
	if _, ok := AllAvailableResolutions[Resolution(resolution)]; !ok {
		resolution = l.Resolution
	}

	// Forming URL
	url := fmt.Sprintf("https://i.ytimg.com/vi/%s/%s.jpg", id, resolution)
	// https://i.ytimg.com/vi/%s/%s.jpg https://img.youtube.com/vi/%s/%s.jpg
	// Sending Get request
	resp, err := http.Get(url)
	if err != nil {
		return nil, context.WithValue(ctx, LoaderFailed{}, id), err
	}
	defer resp.Body.Close()
	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, context.WithValue(ctx, LoaderFailed{}, id), err
	}
	return rawBody, context.WithValue(ctx, LoaderDone{}, id), nil
}
