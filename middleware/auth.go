package middleware

import (
	"strconv"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"saas-admin/model"
)

func CasbinMiddleware(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取用户信息
		user, _ := c.Get("user")
		userId := user.(model.User).ID

		// 获取租户上下文（从请求头或路径参数）
		tenantId := c.GetHeader("X-Tenant-ID")

		// 获取请求路径
		path := c.Request.URL.Path

		// 组合参数
		sub := strconv.FormatUint(uint64(userId), 10)
		dom := tenantId
		if dom == "" {
			dom = "system"
		}

		// 检查权限
		ok, _ := e.Enforce(sub, dom, path)

		if !ok {
			c.AbortWithStatusJSON(403, gin.H{"error": "Forbidden"})
			return
		}
		c.Next()
	}
}
