package provider

import (
	"github.com/ingot-cloud/ingot-go/internal/app/config"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/persist"
)

// BuildCasbin Casbin
func BuildCasbin(adapter persist.Adapter) (*casbin.SyncedEnforcer, func(), error) {
	cfg := config.CONFIG.Casbin

	e, err := casbin.NewSyncedEnforcer(cfg.ModelPath)
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
