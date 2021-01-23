package service

import (
	"context"

	"github.com/ingot-cloud/ingot-go/internal/app/model/dao"
	"github.com/ingot-cloud/ingot-go/internal/app/model/domain"
	"github.com/ingot-cloud/ingot-go/internal/app/model/dto"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/model/types"
)

// Permission 权限
type Permission struct {
	RoleDao          *dao.Role
	RoleAuthorityDao *dao.RoleAuthority
	AuthorityDao     *dao.Authority
	UserDao          *dao.User
	RoleUserDao      *dao.RoleUser
}

// GetRolePolicy 获取所有角色策略
func (p *Permission) GetRolePolicy(ctx context.Context) (*dto.RolePolicys, error) {
	roleList, err := p.RoleDao.List(ctx, dto.QueryCondition{})
	if err != nil {
		return nil, err
	}

	count := len(*roleList)
	// var wait sync.WaitGroup
	// wait.Add(count)
	// ch := make(chan int, count)

	rolePolicys := make([]*dto.RolePolicy, 0, count)

	for _, role := range *roleList {
		ids, err := p.RoleAuthorityDao.GetRoleAuthorityIDs(ctx, role.ID)
		if err != nil {
			return nil, err
		}
		authoritys, err := p.AuthorityDao.GetAuthoritysWithIDs(ctx, dto.QueryCondition{
			IDs: *ids,
		})
		if err != nil {
			return nil, err
		}

		result := make(map[types.ID]domain.SysAuthority)
		p.deepAuthority(ctx, authoritys, &result)

		authorityList := make([]*domain.SysAuthority, 0, len(result))
		for _, authority := range result {
			authorityList = append(authorityList, &authority)
		}

		rolePolicys = append(rolePolicys, &dto.RolePolicy{
			RoleID:        role.ID,
			TenantID:      role.TenantID,
			AuthorityList: authorityList,
		})
	}

	result := dto.RolePolicys(rolePolicys)

	return &result, nil
}

func (p *Permission) deepAuthority(ctx context.Context, list *domain.SysAuthoritys, result *map[types.ID]domain.SysAuthority) error {
	for _, authority := range *list {
		if _, ok := (*result)[authority.ID]; !ok {
			(*result)[authority.ID] = *authority
		}
		childAuthoritys, err := p.AuthorityDao.GetChildWithPID(ctx, dto.QueryCondition{
			ID: authority.ID,
		})
		if err != nil {
			return err
		}
		if len(*childAuthoritys) == 0 {
			continue
		}
		p.deepAuthority(ctx, childAuthoritys, result)
	}
	return nil
}

// GetUserPolicy 获取用户策略
func (p *Permission) GetUserPolicy(ctx context.Context) (*dto.UserPolicys, error) {
	userList, err := p.UserDao.List(ctx, dto.QueryCondition{})
	if err != nil {
		return nil, err
	}

	userPolicys := make([]*dto.UserPolicy, 0, len(*userList))

	for _, user := range *userList {
		ids, err := p.RoleUserDao.GetUserRoleIDs(ctx, user.ID)
		if err != nil {
			return nil, err
		}
		userPolicys = append(userPolicys, &dto.UserPolicy{
			UserID:   user.ID,
			TenantID: user.TenantID,
			RoleList: *ids,
		})
	}

	result := dto.UserPolicys(userPolicys)

	return &result, nil
}
