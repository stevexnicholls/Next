package backup

import (
	"context"
	"encoding/binary"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/stevexnicholls/next/internal/runtime"
	"github.com/stevexnicholls/next/persist"
	log "github.com/stevexnicholls/next/logger"
	"github.com/stevexnicholls/next/restapi/operations/backup"
)

// curl --header "x-api-key: " http://localhost:3000/api/v1alpha/backup -o backup.store

type Backup struct {
	rt *next.Runtime
}

// New returns a new Backup struct
func New(rt *next.Runtime) *Backup {
	return &Backup{rt: rt}
}

func (b *Backup) BackupGet(ctx context.Context, params backup.BackupGetParams) middleware.Responder {
	c, err := b.rt.DB().Backup()

	if err != nil {
		log.Info(swag.String(err.Error()))
		if err == persist.ErrNotFound {
			return backup.NewBackupGetNotFound()
		}
		return backup.NewBackupGetDefault(0)//.WithPayload(modelsError(err))
	}

	log.Infof("get backup")

	return backup.NewBackupGetOK().WithPayload(c)
}

// Itob returns an 8-byte big endian representation of v.
func Itob(v int64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}