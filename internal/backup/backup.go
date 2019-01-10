package backup

import (
	"context"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/stevexnicholls/next/internal/runtime"
	"github.com/stevexnicholls/next/persist"
	"github.com/stevexnicholls/next/restapi/operations/backup"
	log "github.com/stevexnicholls/next/logger"
)

type Backup struct{
	rt *next.Runtime
}

// New returns a new Kv struct
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
