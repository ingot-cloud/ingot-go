package service

import (
	"github.com/ingot-cloud/ingot-go/internal/app/model/dto"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/model/types"
)

// UserDetail 服务
type UserDetail interface {
	// 获取用户详情信息
	GetUserAuthDetails(tenantID types.ID, params dto.UserDetailsDto) (*dto.UserAuthDetails, error)
}
