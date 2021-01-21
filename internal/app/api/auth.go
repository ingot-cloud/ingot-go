package api

import (
	"github.com/ingot-cloud/ingot-go/internal/app/model/dto"
	"github.com/ingot-cloud/ingot-go/internal/app/service"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/errors"
	ginwrapper "github.com/ingot-cloud/ingot-go/pkg/framework/core/wrapper/gin"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/wrapper/response"
	"github.com/ingot-cloud/ingot-go/pkg/framework/log"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"

	"github.com/gin-gonic/gin"
)

// Auth api
type Auth struct {
	AuthService *service.Auth
}

// Login 登录
// path: /api/auth/login method: post
func (a *Auth) Login(ctx *gin.Context) {
	var params dto.LoginParams
	if err := ginwrapper.ParseJSON(ctx, &params); err != nil {
		response.FailureWithError(ctx, err)
		return
	}

	if params.AppID == "" {
		response.FailureWithError(ctx, errors.IllegalArgument("AppID can not be empty"))
		return
	}

	context := ctx.Request.Context()

	userResult, roles, err := a.AuthService.VerifyUserInfo(context, params)
	if err != nil {
		response.FailureWithError(ctx, err)
		return
	}

	user := &security.User{
		ID:       userResult.ID,
		Username: userResult.Username,
		Role:     roles,
	}

	ginwrapper.SetUser(ctx, user)

	token, err := a.AuthService.GenerateToken(context, *user)
	if err != nil {
		response.FailureWithError(ctx, err)
		return
	}

	context = log.NewUserIDContext(context, user.ID)
	log.WithContext(context).Infof("[%s] Login Success", user.Username)
	response.OK(ctx, token)
}

// Logout 登出
// path: /api/auth/logout method: post
func (a *Auth) Logout(ctx *gin.Context) {

	token := ginwrapper.GetToken(ctx)

	if token == "" {
		response.FailureWithError(ctx, errors.ErrUnauthorized)
		return
	}

	user, ok := ginwrapper.GetUser(ctx)
	if !ok {
		response.FailureWithError(ctx, errors.ErrUnauthorized)
		return
	}

	context := ctx.Request.Context()
	err := a.AuthService.RevokeToken(context, user, token)
	if err != nil {
		response.FailureWithError(ctx, err)
		return
	}

	response.OK(ctx, &response.D{})
}
