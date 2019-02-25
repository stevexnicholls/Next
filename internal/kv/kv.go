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

// Kv key value struct 
type Kv struct {
	// data map[int64]*models.KeyValue
	rt *next.Runtime
}

// New returns a new Kv struct
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

	log.Infof("get { key: %v, value: %v }", key, c.Value)

	return kv.NewValueGetOK().WithPayload(c)
}

// ValueUpdate handles updating a key
func (k *Kv) ValueUpdate(ctx context.Context, params kv.ValueUpdateParams) middleware.Responder {
	keyvalue := params.Keyvalue

	if err := k.rt.DB().Update(keyvalue); err != nil {
		log.Info(swag.String(err.Error()))
		if err == persist.ErrGone {
			return kv.NewValueUpdateNotFound().WithPayload(modelsError(err))
		}
		if err == persist.ErrNotFound {
			return kv.NewValueUpdateNotFound().WithPayload(modelsError(err))
		}
		return kv.NewValueUpdateDefault(500).WithPayload(modelsError(err))
	}

	log.Infof("update { key: %v, value: %v }", keyvalue.Key, keyvalue.Value)

	return kv.NewValueUpdateCreated().WithPayload(keyvalue)
}

// KeyDelete handles deleting a key
func (*Kv) KeyDelete(ctx context.Context, params kv.KeyDeleteParams) middleware.Responder {
	return kv.NewKeyDeleteDefault(501)
}

// KeyList returns a full list of keys and values
func (k *Kv) KeyList(ctx context.Context, params kv.KeyListParams) middleware.Responder {

	c, err := k.rt.DB().View()

	if err != nil {
		log.Info(swag.String(err.Error()))
		if err == persist.ErrNotFound {
			return kv.NewKeyListDefault(500).WithPayload(modelsError(err))
		}
		return kv.NewKeyListDefault(500).WithPayload(modelsError(err))
	}

	for _, kv := range c {
		log.Infof("list { key: %+v, value: %+v }", kv.Key, kv.Value)
	}

	return kv.NewKeyListOK().WithPayload(c)
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
