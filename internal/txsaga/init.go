package txsaga

import "context"

type TxLoadSaga struct {
	// Finalized ctxes from parts of application
	LoadCtx context.Context
	DBCtx   context.Context

	Finalisator context.Context
}
