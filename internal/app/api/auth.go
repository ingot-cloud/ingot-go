package api

import (
	"github.com/ingot-cloud/ingot-go/internal/app/model/dto"
	"github.com/ingot-cloud/ingot-go/internal/app/service"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/errors"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/web/ingot"
	ginwrapper "github.com/ingot-cloud/ingot-go/pkg/framework/core/wrapper/gin"
	"github.com/ingot-cloud/ingot-go/pkg/framework/log"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"

	"github.com/gin-gonic/gin"
)

// Auth api
type Auth struct {
	AuthService *service.Auth
}

// Apply api配置
func (a *Auth) Apply(app *ingot.Router) {
	router := app.Group("auth")
	router.POST("/login", a.Login)
	router.POST("/logout", a.Logout)
}

// Login 登录
// path: /api/auth/login method: post
func (a *Auth) Login(ctx *gin.Context) (interface{}, error) {
	var params dto.LoginParams
	if err := ginwrapper.ParseJSON(ctx, &params); err != nil {
		return nil, err
	}

	// if params.AppID == "" {
	// 	response.FailureWithError(ctx, errors.IllegalArgument("AppID can not be empty"))
	// 	return
	// }

	context := ctx.Request.Context()

	userResult, roles, err := a.AuthService.VerifyUserInfo(context, params)
	if err != nil {
		return nil, err
	}

	user := &security.User{
		ID:       userResult.ID,
		Username: userResult.Username,
		Roles:    roles,
	}

	ginwrapper.SetUser(ctx, user)

	token, err := a.AuthService.GenerateToken(context, *user)
	if err != nil {
		return nil, err
	}

	context = log.NewUserIDContext(context, user.ID)
	log.WithContext(context).Infof("[%s] Login Success", user.Username)
	return token, nil
}

// Logout 登出
// path: /api/auth/logout method: post
func (a *Auth) Logout(ctx *gin.Context) (interface{}, error) {

	token := ginwrapper.GetBearerToken(ctx)

	if token == "" {
		return nil, errors.ErrUnauthorized
	}

	user, ok := ginwrapper.GetUser(ctx)
	if !ok {
		return nil, errors.ErrUnauthorized
	}

	context := ctx.Request.Context()
	err := a.AuthService.RevokeToken(context, user, token)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
