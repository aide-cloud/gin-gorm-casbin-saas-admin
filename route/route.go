package route

import (
	"github.com/gin-gonic/gin"
	"saas-admin/controller"
)

func SetupRoutes(r *gin.Engine) {
	// 系统管理接口
	system := r.Group("/system")
	{
		system.POST("/roles", controller.CreateSystemRole)
		system.POST("/tenants", controller.CreateTenant)
		system.POST("/tenants/:id/apis", controller.UpdateTenantApis)
	}

	// 租户管理接口
	tenant := r.Group("/tenant")
	tenant.Use(middleware.RequireTenant())
	{
		tenant.POST("/roles", controller.CreateTenantRole)
		tenant.POST("/members", controller.AddTenantMember)
	}

	// 公共接口
	r.POST("/login", controller.Login)
}
