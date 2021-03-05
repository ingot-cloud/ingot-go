package impl

import (
	"context"
	"fmt"

	"github.com/ingot-cloud/ingot-go/internal/app/service"
	"github.com/ingot-cloud/ingot-go/pkg/framework/log"

	casbinModel "github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
)

// CasbinAdapterService casbin适配器
type CasbinAdapterService struct {
	PermissionService service.Permission
}

// LoadPolicy loads all policy rules from the storage.
func (c *CasbinAdapterService) LoadPolicy(model casbinModel.Model) error {
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
func (c *CasbinAdapterService) SavePolicy(model casbinModel.Model) error {
	return nil
}

// AddPolicy adds a policy rule to the storage.
// This is part of the Auto-Save feature.
func (c *CasbinAdapterService) AddPolicy(sec string, ptype string, rule []string) error {
	return nil
}

// RemovePolicy removes a policy rule from the storage.
// This is part of the Auto-Save feature.
func (c *CasbinAdapterService) RemovePolicy(sec string, ptype string, rule []string) error {
	return nil
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
// This is part of the Auto-Save feature.
func (c *CasbinAdapterService) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	return nil
}

// 加载角色策略. p,角色,租户,资源,操作
func (c *CasbinAdapterService) loadRolePolicy(ctx context.Context, model casbinModel.Model) error {
	policys, err := c.PermissionService.GetRolePolicy(ctx)
	if err != nil {
		return nil
	}

	for _, item := range *policys {
		for _, authority := range item.AuthorityList {
			line := fmt.Sprintf("p,%s,%d,%s,%s", item.RoleID, item.TenantID, authority.Path, authority.Method) //"[GET|POST|PUT|DELETE|HEAD]"
			persist.LoadPolicyLine(line, model)
		}
	}

	return nil
}

// 用户角色关联策略. g,用户,角色,租户
func (c *CasbinAdapterService) loadUserPolicy(ctx context.Context, model casbinModel.Model) error {
	policys, err := c.PermissionService.GetUserPolicy(ctx)
	if err != nil {
		return nil
	}

	for _, item := range *policys {
		for _, role := range item.RoleList {
			line := fmt.Sprintf("g,%s,%s,%d", item.UserID, role, item.TenantID)
			persist.LoadPolicyLine(line, model)
		}
	}

	return nil
}
