package service

import (
	"context"

	"github.com/ingot-cloud/ingot-go/internal/app/model/dao"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/userdetails"
)

// UserDetails 服务
type UserDetails struct {
	UserDao *dao.User
}

// LoadUserByUsername 加载指定 username 的用户
func (u *UserDetails) LoadUserByUsername(username string) (userdetails.UserDetails, error) {
	user, err := u.UserDao.One(context.TODO(), username)
	if err != nil {
		return nil, err
	}

	// todo 查询角色

	return nil, nil
}
