package domain

import (
	"time"

	"github.com/ingot-cloud/ingot-go/pkg/framework/core/model/types"
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
