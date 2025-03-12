package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"saas-admin/model"
	"saas-admin/utils"
)

// 创建租户角色
func CreateTenantRole(c *gin.Context) {
	tenantID := c.MustGet("tenant_id").(uint)

	var input struct {
		Name string
		Apis []uint
	}
	c.BindJSON(&input)

	db := utils.GetDB()

	// 验证API是否属于租户授权列表
	var tenant model.Tenant
	db.Preload("Apis").First(&tenant, tenantID)

	var validApis []model.Api
	db.Model(&tenant).Association("Apis").Find(&validApis, input.Apis)

	if len(validApis) != len(input.Apis) {
		c.JSON(400, gin.H{"error": "包含未授权的API"})
		return
	}

	// 创建角色
	role := model.TenantRole{
		TenantID:  tenantID,
		Name:      input.Name,
		IsBuiltIn: false,
		Apis:      validApis,
	}
	db.Create(&role)

	casbin := utils.GetCasbin()
	// 添加Casbin策略
	for _, api := range validApis {
		_, _ = casbin.AddPolicy(strconv.FormatUint(uint64(role.ID), 10),
			strconv.FormatUint(uint64(tenantID), 10),
			api.Path)
	}

	c.JSON(200, role)
}

// 给用户分配系统角色
func AssignSystemRole(c *gin.Context) {
	var input struct {
		UserID uint
		RoleID uint
	}
	c.BindJSON(&input)

	casbin := utils.GetCasbin()
	// 添加Casbin分组
	_, _ = casbin.AddGroupingPolicy(strconv.FormatUint(uint64(input.UserID), 10), strconv.FormatUint(uint64(input.RoleID), 10), "system")

	c.JSON(200, gin.H{"message": "分配成功"})
}
