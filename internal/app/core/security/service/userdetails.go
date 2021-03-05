package service

import (
	"github.com/ingot-cloud/ingot-go/internal/app/core/security/user"
	"github.com/ingot-cloud/ingot-go/internal/app/model/dto"
	"github.com/ingot-cloud/ingot-go/internal/app/service"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/authority"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/userdetails"
)

// UserDetails 服务
type UserDetails struct {
	UserDetailService service.UserDetail
}

// LoadUserByUsername 加载指定 username 的用户
func (u *UserDetails) LoadUserByUsername(username string) (userdetails.UserDetails, error) {

	// TODO 临时测试
	params := dto.UserDetailsDto{
		UniqueCode: username,
	}
	details, err := u.UserDetailService.GetUserAuthDetails(1, params)
	if err != nil {
		return nil, err
	}
	coreUser := userdetails.NewUserAllParams(username, details.Password, authority.CreateAuthorityList(details.Roles), true, true, true, true)
	return user.NewIngotUser(details.ID, details.DeptID, details.TenantID, details.AuthType, coreUser), nil
}
