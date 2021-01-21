package provider

import (
	"context"
	"fmt"
	"github.com/ingot-cloud/ingot-go/internal/app/common/log"
	"github.com/ingot-cloud/ingot-go/internal/app/model/dao"
	"github.com/ingot-cloud/ingot-go/internal/app/model/dto"
	"github.com/ingot-cloud/ingot-go/internal/app/model/enums"

	casbinModel "github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"github.com/google/wire"
)

// CasbinAdapterSet inject persist.Adapter
var CasbinAdapterSet = wire.NewSet(wire.Struct(new(CasbinAdapter), "*"), wire.Bind(new(persist.Adapter), new(*CasbinAdapter)))

// CasbinAdapter casbin适配器
type CasbinAdapter struct {
	UserDao          *dao.User
	AuthorityDao     *dao.Authority
	RoleDao          *dao.Role
	RoleUserDao      *dao.RoleUser
	RoleAuthorityDao *dao.RoleAuthority
}

// LoadPolicy loads all policy rules from the storage.
func (c *CasbinAdapter) LoadPolicy(model casbinModel.Model) error {
	ctx := context.Background()
	err := c.loadRolePolicy(ctx, model)
	if err != nil {
		log.WithContext(ctx).Errorf("Load casbin role policy error: %s", err.Error())
		return err
	}

	err = c.loadUserPolicy(ctx, model)
	if err != nil {
		log.WithContext(ctx).Errorf("Load casbin user policy error: %s", err.Error())
		return err
	}

	return nil
}

// SavePolicy saves all policy rules to the storage.
func (c *CasbinAdapter) SavePolicy(model casbinModel.Model) error {
	return nil
}

// AddPolicy adds a policy rule to the storage.
// This is part of the Auto-Save feature.
func (c *CasbinAdapter) AddPolicy(sec string, ptype string, rule []string) error {
	return nil
}

// RemovePolicy removes a policy rule from the storage.
// This is part of the Auto-Save feature.
func (c *CasbinAdapter) RemovePolicy(sec string, ptype string, rule []string) error {
	return nil
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
// This is part of the Auto-Save feature.
func (c *CasbinAdapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	return nil
}

func (c *CasbinAdapter) loadRolePolicy(ctx context.Context, model casbinModel.Model) error {
	result, err := c.AuthorityDao.RoleAuthority(ctx, dto.QueryStatusParams{Status: enums.StatusEnabled})
	if err != nil {
		return nil
	}

	roleAuthorityList := result.List
	if len(roleAuthorityList) == 0 {
		return nil
	}

	for _, item := range roleAuthorityList {
		line := fmt.Sprintf("p,%s,%s,%s", item.RoleID, item.Path, "[GET|POST|PUT|DELETE|HEAD]")
		persist.LoadPolicyLine(line, model)
	}

	return nil
}

func (c *CasbinAdapter) loadUserPolicy(ctx context.Context, model casbinModel.Model) error {
	userResult, err := c.UserDao.List(ctx, dto.UserQueryParams{Status: enums.StatusEnabled})
	if err != nil {
		return nil
	}

	users := userResult.List
	if len(users) == 0 {
		return nil
	}

	roleUserResult, err := c.RoleUserDao.List(ctx, dto.RoleUserQueryParams{})
	if err != nil {
		return nil
	}

	roleUsers := roleUserResult.List
	if len(roleUsers) == 0 {
		return nil
	}

	// roleUsers 改变结构为 userId => [roleId, roleId]
	mapUserIDRoleIds := make(map[string][]string)
	for _, item := range roleUsers {
		mapUserIDRoleIds[item.UserID] = append(mapUserIDRoleIds[item.UserID], item.RoleID)
	}

	for _, user := range users {
		userID := user.ID
		if roleIds, ok := mapUserIDRoleIds[userID]; ok {
			for _, roleID := range roleIds {
				line := fmt.Sprintf("g,%s,%s", userID, roleID)
				persist.LoadPolicyLine(line, model)
			}
		}
	}

	return nil
}
