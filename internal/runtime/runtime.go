package next

import (
	"github.com/spf13/viper"
	"github.com/stevexnicholls/next/persist"
)

// NewRuntime creates a new application level runtime that encapsulates the shared services for this application
func NewRuntime() (*Runtime, error) {
	db, err := persist.NewBoltStore(viper.GetString("db_path"), viper.GetString("db_bucket"))
	if err != nil {
		return nil, err
	}

	return &Runtime{
		db: db,
	}, nil
}

// Runtime encapsulates the shared services for this application
type Runtime struct {
	db persist.Store
}

// DB returns the persistent store
func (r *Runtime) DB() persist.Store {
	return r.db
}

// Close closes the database associated with the runtime
func (r *Runtime) Close() {
	r.db.Close()
}

// Config returns the viper config for this application
// func (r *Runtime) Config() *viper.Viper {
// 	return r.cfg
// }

// // Tracer returns the root tracer, this is typically the only one you need
// func (r *Runtime) Tracer() tracing.Tracer {
// 	return r.app.Tracer()
// }

// // Logger gets the root logger for this application
// func (r *Runtime) Logger() logrus.FieldLogger {
// 	return r.app.Logger()
// }

// // NewLogger creates a new named logger for this application
// func (r *Runtime) NewLogger(name string, fields logrus.Fields) logrus.FieldLogger {
// 	return r.app.NewLogger(name, fields)
// }
