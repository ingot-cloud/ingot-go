package service

import (
	"context"

	"github.com/ingot-cloud/ingot-go/internal/app/model/dto"
)

// Permission 权限服务
type Permission interface {
	// 获取所有角色策略
	GetRolePolicy(ctx context.Context) (*dto.RolePolicys, error)
	// 获取用户策略
	GetUserPolicy(ctx context.Context) (*dto.UserPolicys, error)
}
