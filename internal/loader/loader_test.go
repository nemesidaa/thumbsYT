package loader_test

import (
	"context"
	"errors"
	"testing"

	ldr "github.com/nemesidaa/thumbsYT/internal/loader"
)

var (
	ErrWithoutCtx = errors.New("without context")
)

func TestLoader(t *testing.T) {
	//Rickroll!!!
	loader := ldr.NewLoader("hqdefault")
	_, ctx, err := loader.Load("dQw4w9WgXcQ", "hqdefault", context.Background())
	if err != nil {
		t.Error(err)
	}
	if _, ok := ctx.Value(ldr.LoaderDone{}).(string); !ok {
		t.Error(ErrWithoutCtx)
	}

}

func BenchmarkLoader(b *testing.B) {
	loader := ldr.NewLoader("hqdefault")
	_, ctx, err := loader.Load("dQw4w9WgXcQ", "hqdefault", context.Background())
	if err != nil {
		b.Error(err)
	}
	if _, ok := ctx.Value(ldr.LoaderDone{}).(string); !ok {
		b.Error(ErrWithoutCtx)
	}
}
