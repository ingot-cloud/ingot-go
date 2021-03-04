package provider

import (
	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/internal/app/core/security/service"
)

// SecurityClientDetailsService 服务实现
var SecurityClientDetailsService = wire.Struct(new(service.ClientDetails), "*")

// SecurityUserDetailsService 服务实现
var SecurityUserDetailsService = wire.Struct(new(service.UserDetails), "*")

// CustomSecurityAll 自定义
var CustomSecurityAll = wire.NewSet(
	SecurityClientDetailsService,
	SecurityUserDetailsService,
)
