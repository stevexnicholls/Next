package kv

import (
	"context"
	"encoding/binary"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/stevexnicholls/next/internal/runtime"
	"github.com/stevexnicholls/next/models"
	"github.com/stevexnicholls/next/persist"
	"github.com/stevexnicholls/next/restapi/operations/kv"
	log "github.com/stevexnicholls/next/logger"
)

type Kv struct {
	// data map[int64]*models.KeyValue
	rt *next.Runtime
}

// New retursn a new Kv struct
func New(rt *next.Runtime) *Kv {
	return &Kv{rt: rt}
}

// ValueGet handles getting a value
func (k *Kv) ValueGet(ctx context.Context, params kv.ValueGetParams) middleware.Responder {
	key := params.Key

	c, err := k.rt.DB().Get(key)

	if err != nil {
		log.Info(swag.String(err.Error()))
		if err == persist.ErrNotFound {
			return kv.NewValueGetNotFound() //kv.NewGetEntryNotFound().WithPayload(modelsError(err))
		}
		return kv.NewValueGetDefault(0).WithPayload(modelsError(err))
	}

	log.Infof("get key: %v value: %v", key, c.Value)

	return kv.NewValueGetOK().WithPayload(c)
}

func (k *Kv) ValueUpdate(ctx context.Context, params kv.ValueUpdateParams) middleware.Responder {
	keyvalue := params.Keyvalue

	if err := k.rt.DB().Update(keyvalue); err != nil {
		// if err == persist.ErrVersionMismatch {
		// 	return kv.NewPutEntryConflict().WithXRequestID(rid).WithPayload(modelsError(err))
		// }
		if err == persist.ErrGone {
			return kv.NewValueUpdateNotFound() //NewUpdateValueGone() //.WithXRequestID(rid).WithPayload(modelsError(errors.New("entry was deleted")))
		}
		if err == persist.ErrNotFound {
			return kv.NewValueUpdateNotFound() //.WithXRequestID(rid).WithPayload(modelsError(err))
		}
		return kv.NewValueUpdateDefault(0) //.NewPutEntryDefault(0).WithXRequestID(rid).WithPayload(modelsError(err))
	}

	return kv.NewValueUpdateCreated().WithPayload(keyvalue)
}

func (*Kv) KeyDelete(ctx context.Context, params kv.KeyDeleteParams) middleware.Responder {
	panic("implement me")
}

func (*Kv) KeyList(cts context.Context, params kv.KeyListParams) middleware.Responder {
	panic("implement me")
}

func modelsError(err error) *models.Error {
	return &models.Error{
		Message: swag.String(err.Error()),
	}
}

// Itob returns an 8-byte big endian representation of v.
func Itob(v int64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
