package health

import (
	"context"

	"github.com/go-openapi/runtime/middleware"
	"github.com/stevexnicholls/next/restapi/operations/health"
)

type Health struct{}

func (*Health) HealthGet(ctx context.Context, params health.GetHealthParams) middleware.Responder {
	return health.NewGetHealthOK()
}
