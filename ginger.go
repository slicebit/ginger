package ginger

import (
	"github.com/slicebit/ginger/app"
	"github.com/slicebit/ginger/config"
	"github.com/slicebit/ginger/db"
)

// NewConfig creates a new configuration
func NewConfig(cfgFile string, cfg interface{}) error {
	return config.New(cfgFile, cfg)
}

// NewApp creates a new gin app given config file
func NewApp(host string, port int, env string) (*app.App, error) {
	return app.New(host, port, env)
}

// NewDB creates a new db context given driver & dsn
func NewDB(driver string, dsn string) (*db.Context, error) {
	return db.New(driver, dsn)
}
