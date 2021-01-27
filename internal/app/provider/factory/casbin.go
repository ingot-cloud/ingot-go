package factory

import (
	"github.com/ingot-cloud/ingot-go/internal/app/config"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/persist"
)

// NewCasbin Casbin
func NewCasbin(options *config.Options, adapter persist.Adapter) (*casbin.SyncedEnforcer, func(), error) {
	e, err := casbin.NewSyncedEnforcer(options.CasbinModelFile)
	if err != nil {
		return nil, nil, err
	}

	e.EnableLog(true)

	err = e.InitWithModelAndAdapter(e.GetModel(), adapter)
	if err != nil {
		return nil, nil, err
	}

	e.EnableEnforce(true)

	return e, func() {}, nil
}
