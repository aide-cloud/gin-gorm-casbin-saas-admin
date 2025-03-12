package utils

import (
	_ "embed"
	"sync"

	"github.com/casbin/casbin/v2"
	casbinModel "github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

//go:embed rbac_model.conf
var rbacModelConf string

var (
	rbacOnce sync.Once
	enforcer *casbin.SyncedEnforcer
)

// GetCasbin ...
func GetCasbin() *casbin.SyncedEnforcer {
	return enforcer
}

// InitCasbinModel 初始化 casbin 模型
func InitCasbinModel(db *gorm.DB) (*casbin.SyncedEnforcer, error) {
	if enforcer != nil {
		return enforcer, nil
	}
	var err error

	rbacOnce.Do(func() {
		var adapter *gormadapter.Adapter
		var rbacModel casbinModel.Model
		adapter, err = gormadapter.NewAdapterByDB(db)
		if err != nil {
			return
		}
		rbacModel, err = casbinModel.NewModelFromString(rbacModelConf)
		if err != nil {
			return
		}
		enforcer, err = casbin.NewSyncedEnforcer(rbacModel, adapter)
		if err != nil {
			return
		}

		// 加载策略
		if err = enforcer.LoadPolicy(); err != nil {
			return
		}
	})

	return enforcer, nil
}
