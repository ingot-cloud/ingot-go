package impl

import (
	"github.com/ingot-cloud/ingot-go/internal/app/model/dto"
	"github.com/ingot-cloud/ingot-go/internal/app/model/enums"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/model/types"
)

// UserDetail 服务实现
type UserDetail struct {
}

// GetUserAuthDetails 获取用户详情信息
func (u *UserDetail) GetUserAuthDetails(tenantID types.ID, params dto.UserDetailsDto) (*dto.UserAuthDetails, error) {
	// todo 查询该用户所有角色，查询授权应用，以及对应的 authType

	// TODO 目前临时返回固定测试值
	return &dto.UserAuthDetails{
		ID:       1,
		DeptID:   1,
		TenantID: 1,
		Username: params.UniqueCode,
		Password: "{noop}" + params.UniqueCode,
		Status:   enums.UserStatusEnable,
		AuthType: "standard",
		Roles:    []string{"admin"},
	}, nil
}
