package service

import (
	"context"
	"strings"
	"time"

	"github.com/ingot-cloud/ingot-go/internal/app/core/security"
	"github.com/ingot-cloud/ingot-go/internal/app/core/security/password"
	"github.com/ingot-cloud/ingot-go/internal/app/model/dao"
	"github.com/ingot-cloud/ingot-go/internal/app/model/domain"
	"github.com/ingot-cloud/ingot-go/internal/app/model/dto"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/errors"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/model/enums"
)

// Auth service
type Auth struct {
	UserDao         *dao.User
	RoleUserDao     *dao.RoleUser
	RoleAppDao      *dao.RoleApp
	Auth            security.Authentication
	PasswordEncoder password.Encoder
}

// VerifyUserInfo 验证用户信息
func (a *Auth) VerifyUserInfo(ctx context.Context, params dto.LoginParams) (*domain.User, []string, error) {
	user, err := a.UserDao.One(ctx, params.Username)
	if err != nil {
		return nil, nil, errors.ErrUserInvalid
	}

	if !a.PasswordEncoder.Matches(params.Password, user.Password) {
		return nil, nil, errors.ErrUserInvalid
	}

	if strings.Compare(user.Status, string(enums.StatusDisabled)) == 0 {
		return nil, nil, errors.ErrUserDisabled
	}

	roleUserResult, err := a.RoleUserDao.GetUserRole(ctx, user.ID)
	if err != nil {
		return nil, nil, err
	}

	roleAppList, err := a.RoleAppDao.GetAppRole(ctx, params.AppID)
	if err != nil {
		return nil, nil, err
	}

	var roles []string
	for _, roleUser := range roleUserResult.List {
		if !a.appIncludeRole(roleAppList, roleUser.RoleID) {
			return nil, nil, errors.ErrUserAppForbidden
		}
		roles = append(roles, roleUser.RoleID)
	}

	return user, roles, nil
}

func (a *Auth) appIncludeRole(arr []domain.RoleApp, target string) bool {
	for _, roleApp := range arr {
		if roleApp.RoleID == target {
			return true
		}
	}
	return false
}

// GenerateToken 生成Token
func (a *Auth) GenerateToken(ctx context.Context, user security.User) (*dto.LoginResult, error) {
	accessToken, err := a.Auth.GenerateToken(ctx, user)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	expire := time.Unix(accessToken.GetExpiration(), 0).Sub(time.Now())

	// store token
	a.Auth.GetStore().Store(ctx, accessToken.GetValue(), expire)

	return &dto.LoginResult{
		Username:    user.Username,
		Role:        user.Role,
		AccessToken: accessToken.GetValue(),
		TokenType:   accessToken.GetTokenType(),
		Expiration:  accessToken.GetExpiration(),
	}, nil
}

// RevokeToken 撤销Token
func (a *Auth) RevokeToken(ctx context.Context, user *security.User, token string) error {
	return a.Auth.RevokeToken(ctx, token)
}
