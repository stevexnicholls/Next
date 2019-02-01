package health

import (
	"context"

	"github.com/go-openapi/runtime/middleware"
	"github.com/stevexnicholls/next/restapi/operations/health"
)

type Health struct{}

func (*Health) HealthGet(ctx context.Context, params health.HealthGetParams) middleware.Responder {
	return health.NewHealthGetOK()
}
