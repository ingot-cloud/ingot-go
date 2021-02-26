package dao

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/userdetails"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/crypto/password"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/errors"
)

// AuthenticationProvider Dao 提供者
type AuthenticationProvider struct {
	PasswordEncoder          password.Encoder
	UserDetailsService       userdetails.Service
	UserCache                userdetails.UserCache
	PreAuthenticationChecks  userdetails.Checker
	PostAuthenticationChecks userdetails.Checker
}

// NewProvider 实例化
func NewProvider(encoder password.Encoder, service userdetails.Service, cache userdetails.UserCache, preChecker userdetails.Checker, postChecker userdetails.Checker) *AuthenticationProvider {
	return &AuthenticationProvider{
		PasswordEncoder:          encoder,
		UserDetailsService:       service,
		UserCache:                cache,
		PreAuthenticationChecks:  preChecker,
		PostAuthenticationChecks: postChecker,
	}
}

// Authenticate 身份验证
func (p *AuthenticationProvider) Authenticate(auth core.Authentication) (core.Authentication, error) {
	username := p.determineUsername(auth)
	cacheWasUsed := true
	// Supports 方法已经确定了该 auth 为 UsernamePasswordAuthenticationToken
	userAuth, _ := auth.(*authentication.UsernamePasswordAuthenticationToken)
	user, err := p.UserCache.GetUserFromCache(username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		cacheWasUsed = false
		user, err = p.retrieveUser(username, userAuth)
		if err != nil {
			return nil, err
		}
	}
	// 前置检测
	err = p.PreAuthenticationChecks.Check(user)
	if err != nil {
		return nil, err
	}
	err = p.additionalAuthenticationChecks(user, userAuth)
	if err != nil {
		// 如果检测的 user 不是缓存的，那么直接抛出异常
		if !cacheWasUsed {
			return nil, err
		}
		// 如果检测的 user 是从缓存中获取的，那么重新获取最新数据进行检查
		user, err = p.retrieveUser(username, userAuth)
		if err != nil {
			return nil, err
		}
		err = p.PreAuthenticationChecks.Check(user)
		if err != nil {
			return nil, err
		}
		err = p.additionalAuthenticationChecks(user, userAuth)
		if err != nil {
			return nil, err
		}
	}
	// 后置检测
	err = p.PostAuthenticationChecks.Check(user)
	if err != nil {
		return nil, err
	}
	if !cacheWasUsed {
		p.UserCache.PutUserInCache(user)
	}

	return p.createSuccessAuthentication(user, auth, user)
}

// Supports 该身份验证提供者是否支持指定的认证信息
func (p *AuthenticationProvider) Supports(auth interface{}) bool {
	_, ok := auth.(*authentication.UsernamePasswordAuthenticationToken)
	return ok
}

func (p *AuthenticationProvider) determineUsername(auth core.Authentication) string {
	if auth.GetPrincipal() == nil {
		return "NONE_PROVIDED"
	}

	return auth.GetName(auth)
}

func (p *AuthenticationProvider) retrieveUser(username string, auth *authentication.UsernamePasswordAuthenticationToken) (userdetails.UserDetails, error) {
	loadedUser, err := p.UserDetailsService.LoadUserByUsername(username)
	if err != nil {
		return nil, err
	}
	if loadedUser == nil {
		return nil, errors.UsernameNotFound("Failed to find user: ", username)
	}
	return loadedUser, nil
}

func (p *AuthenticationProvider) additionalAuthenticationChecks(userDetails userdetails.UserDetails, auth *authentication.UsernamePasswordAuthenticationToken) error {
	if auth.GetCredentials() == "" {
		return errors.BadCredentials("Bad credentials")
	}
	presentedPassword := auth.GetCredentials()
	ok, err := p.PasswordEncoder.Matches(presentedPassword, userDetails.GetPassword())
	if err != nil {
		return err
	}
	if !ok {
		return errors.BadCredentials("Bad credentials")
	}
	return nil
}

func (p *AuthenticationProvider) createSuccessAuthentication(principal interface{}, auth core.Authentication, user userdetails.UserDetails) (core.Authentication, error) {
	result := authentication.NewAuthenticatedUsernamePasswordAuthToken(principal, auth.GetCredentials(), user.GetAuthorities())
	result.SetDetails(auth.GetDetails())
	return result, nil
}
