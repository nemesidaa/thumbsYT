package txbroker

import (
	"context"

	store "github.com/nemesidaa/thumbsYT/internal/storage/store/sqlstore"
)

const (
	StatusFatal   = "fatal"
	StatusError   = "error"
	StatusSuccess = "success"
)

// Attempt to realize effective-once delivery guarantee to client
type TxLoadLazyStack struct {
	// Finalized ctxes from parts of application
	ServCtx    context.Context
	UndoneChan chan TxLoadLazyStackComponent

	// Retries
	MaxRetriesCounter int8

	// MemCached fatal requests contains there, with their retries count
	MemCached map[TxLoadLazyStackComponent]struct{}

	// IdealCaching
	IdealCaching bool

	// Storage
	Storage *store.Storage
}

type TxLoadLazyStackComponent struct {
	ProcCtx        context.Context
	RequestID      string
	RetriesCounter int8
}

func NewLazyBroker(n int, maxRetries int8, idealCaching bool, db *store.Storage) *TxLoadLazyStack {
	return &TxLoadLazyStack{
		UndoneChan:        make(chan TxLoadLazyStackComponent, n),
		MaxRetriesCounter: maxRetries,
		IdealCaching:      idealCaching,
		Storage:           db,
	}
}

func (s *TxLoadLazyStack) Close() {
	close(s.UndoneChan)
}
