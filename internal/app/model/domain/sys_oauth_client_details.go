package domain

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/ingot-cloud/ingot-go/pkg/framework/core/model/types"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/authority"
)

// SysOauthClientDetails OAuth2 Client
type SysOauthClientDetails struct {
	ID                    types.ID `gorm:"primary_key;size:20"`
	Version               int64
	TenantID              int
	ClientID              string
	ClientSecret          string
	ResourceID            string
	ResourceIDs           string
	Scope                 string
	AuthorizedGrantTypes  string
	WebServerRedirectURI  string
	Authorities           string
	AccessTokenValidity   int
	RefreshTokenValidity  int
	AdditionalInformation string
	Autoapprove           string
	AuthType              string
	Type                  string
	Status                string
	Remark                string
	CreatedAt             time.Time
	UpdatedAt             time.Time
	DeletedAt             *time.Time
}

// TableName 表名
func (*SysOauthClientDetails) TableName() string {
	return "sys_oauth_client_details"
}

// 实现 clientdetails.ClientDetails 接口

// GetClientID 获取ClientID
func (c *SysOauthClientDetails) GetClientID() string {
	return c.ClientID
}

// GetResourceIDs 获取可以访问的资源ID
func (c *SysOauthClientDetails) GetResourceIDs() []string {
	if c.ResourceIDs == "" {
		return nil
	}
	return strings.Split(c.ResourceIDs, ",")
}

// IsSecretRequired 验证此客户端是否需要秘钥
func (c *SysOauthClientDetails) IsSecretRequired() bool {
	return c.ClientSecret != ""
}

// GetClientSecret 获取客户端秘钥
func (c *SysOauthClientDetails) GetClientSecret() string {
	return c.ClientSecret
}

// IsScoped 是否需要验证 scope
func (c *SysOauthClientDetails) IsScoped() bool {
	if c.Scope == "" {
		return false
	}
	return len(strings.Split(c.Scope, ",")) != 0
}

// GetScope 获取 scope
func (c *SysOauthClientDetails) GetScope() []string {
	if c.Scope == "" {
		return nil
	}
	return strings.Split(c.Scope, ",")
}

// GetAuthorizedGrantTypes 客户端的授权类型
func (c *SysOauthClientDetails) GetAuthorizedGrantTypes() []string {
	if c.AuthorizedGrantTypes == "" {
		return nil
	}
	return strings.Split(c.AuthorizedGrantTypes, ",")
}

// GetRegisteredRedirectURI 获取授权码模式预定义的重定向uri
func (c *SysOauthClientDetails) GetRegisteredRedirectURI() []string {
	if c.WebServerRedirectURI == "" {
		return nil
	}
	return strings.Split(c.WebServerRedirectURI, ",")
}

// GetAuthorities 获取授予客户端的权限
func (c *SysOauthClientDetails) GetAuthorities() []core.GrantedAuthority {
	if c.Authorities == "" {
		return nil
	}
	authorities := strings.Split(c.Authorities, ",")
	return authority.CreateAuthorityList(authorities)
}

// GetAccessTokenValiditySeconds 客户端访问令牌有效时间，单位秒
func (c *SysOauthClientDetails) GetAccessTokenValiditySeconds() int {
	return c.AccessTokenValidity
}

// GetRefreshTokenValiditySeconds 客户端刷新令牌有效时间，单位秒
func (c *SysOauthClientDetails) GetRefreshTokenValiditySeconds() int {
	return c.RefreshTokenValidity
}

// IsAutoApprove 指定scope是否需要用户授权批准，如果不需要用户批准则返回ture
func (c *SysOauthClientDetails) IsAutoApprove(scope string) bool {
	if c.Autoapprove == "" {
		return false
	}
	for _, auto := range strings.Split(c.Autoapprove, ",") {
		if auto == "true" || scope == auto {
			return true
		}
	}
	return false
}

// GetAdditionalInformation 客户端的额外附加信息
func (c *SysOauthClientDetails) GetAdditionalInformation() map[string]interface{} {
	info := make(map[string]interface{})
	json.Unmarshal([]byte(c.AdditionalInformation), &info)
	return info
}
