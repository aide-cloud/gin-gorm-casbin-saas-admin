package model

import "gorm.io/gorm"

// User 系统用户
type User struct {
	gorm.Model
	Username    string `gorm:"uniqueIndex"`
	Password    string
	SystemRoles []SystemRole `gorm:"many2many:user_system_roles;"`
}

// SystemRole 系统角色（硬编码）
type SystemRole struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex"`
	Description string
	Apis        []Api `gorm:"many2many:system_role_apis;"`
}

// Tenant 租户
type Tenant struct {
	gorm.Model
	Name string `gorm:"uniqueIndex"`
	Apis []Api  `gorm:"many2many:tenant_apis;"`
}

// TenantMember 租户成员
type TenantMember struct {
	gorm.Model
	UserID   uint
	TenantID uint
	Roles    []TenantRole `gorm:"many2many:member_roles;"`
}

// TenantRole 租户角色（包含内置和自定义）
type TenantRole struct {
	gorm.Model
	TenantID  uint
	Name      string
	IsBuiltIn bool  // 标识是否是内置角色
	Apis      []Api `gorm:"many2many:tenant_role_apis;"`
}

// Api API列表
type Api struct {
	gorm.Model
	Path string `gorm:"uniqueIndex;size:255"`
}
