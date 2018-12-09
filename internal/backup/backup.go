package backup

import (
	"context"

	"github.com/go-openapi/runtime/middleware"
	"github.com/stevexnicholls/next/restapi/operations/backup"
)

type Backup struct{}

func (*Backup) BackupGet(ctx context.Context, params backup.BackupGetParams) middleware.Responder {
	panic("implement me")
}
