package main

import (
	"fmt"

	"saas-admin/utils"
)

func init() {
	db, err := utils.InitGORM("./gorm.db")
	if err != nil {
		panic(err)
	}
	_, err = utils.InitCasbinModel(db)
	if err != nil {
		panic(err)
	}
}

func main() {
	casbin := utils.GetCasbin()

	// ===================== 1. 为租户 tenant1 创建角色并分配权限 =====================
	// 管理员角色（admin）权限
	_, _ = casbin.AddPolicy("admin", "tenant1", "/api/v1/tenant/create", "*")
	_, _ = casbin.AddPolicy("admin", "tenant1", "/api/v1/tenant/update", "*")
	_, _ = casbin.AddPolicy("admin", "tenant1", "/api/v1/tenant/delete", "*")
	_, _ = casbin.AddPolicy("admin", "tenant1", "/api/v1/tenant/list", "*")
	_, _ = casbin.AddPolicy("admin", "tenant1", "/api/v1/tenant/get", "*")

	// 普通用户角色（user）权限
	_, _ = casbin.AddPolicy("user", "tenant1", "/api/v1/tenant/list", "GET")
	_, _ = casbin.AddPolicy("user", "tenant1", "/api/v1/tenant/get", "GET")

	// ===================== 2. 将用户关联到角色（指定租户域） =====================
	// 用户 user1 在 tenant1 下拥有 admin 角色
	_, _ = casbin.AddGroupingPolicy("user1", "admin", "tenant1")

	// 用户 user2 在 tenant1 下拥有 user 角色
	_, _ = casbin.AddGroupingPolicy("user2", "user", "tenant1")

	// ===================== 3. 验证权限 =====================
	// 检查 user1 在 tenant1 是否有权创建租户
	ok, _ := casbin.Enforce("user1", "tenant1", "/api/v1/tenant/create", "POST")
	fmt.Println("user1 是否有权创建租户:", ok) // 输出 true

	// 检查 user2 在 tenant1 是否有权删除租户
	ok, _ = casbin.Enforce("user2", "tenant1", "/api/v1/tenant/delete", "DELETE")
	fmt.Println("user2 是否有权删除租户:", ok) // 输出 false
}
