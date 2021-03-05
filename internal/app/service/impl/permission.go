package impl

import (
	"context"
	"sync"
	"time"

	"github.com/ingot-cloud/ingot-go/internal/app/model/dao"
	"github.com/ingot-cloud/ingot-go/internal/app/model/domain"
	"github.com/ingot-cloud/ingot-go/internal/app/model/dto"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/model/types"
	"github.com/ingot-cloud/ingot-go/pkg/framework/log"
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
	var sw sync.WaitGroup
	sw.Add(count)
	finishChan := make(chan interface{}, 1)
	errorChan := make(chan error, count)

	rolePolicys := make([]*dto.RolePolicy, 0, count)

	log.Debugf("开始获取角色策略")
	startNanosecond := time.Now().Nanosecond()

	for _, role := range *roleList {
		go func(role *domain.SysRole) {
			ids, err := p.RoleAuthorityDao.GetRoleAuthorityIDs(ctx, role.ID)
			if err != nil {
				errorChan <- err
				return
			}
			authoritys, err := p.AuthorityDao.GetAuthoritysWithIDs(ctx, dto.QueryCondition{
				IDs: *ids,
			})
			if err != nil {
				errorChan <- err
				return
			}

			result := make(map[types.ID]*domain.SysAuthority)
			p.deepAuthority(ctx, authoritys, &result)

			authorityList := make([]*domain.SysAuthority, 0, len(result))
			for _, authority := range result {
				authorityList = append(authorityList, authority)
			}

			rolePolicys = append(rolePolicys, &dto.RolePolicy{
				RoleID:        role.ID,
				TenantID:      role.TenantID,
				AuthorityList: authorityList,
			})

			sw.Done()
		}(role)
	}

	go func() {
		sw.Wait()
		log.Debugf("角色策略获取完成，用时%d毫秒", (time.Now().Nanosecond()-startNanosecond)/1e6)
		finishChan <- struct{}{}
		close(finishChan)
	}()

	select {
	case err := <-errorChan:
		return nil, err
	case <-finishChan:
		close(errorChan)
		result := dto.RolePolicys(rolePolicys)
		return &result, nil
	}
}

func (p *Permission) deepAuthority(ctx context.Context, list *[]*domain.SysAuthority, result *map[types.ID]*domain.SysAuthority) error {
	for _, authority := range *list {
		if _, ok := (*result)[authority.ID]; !ok {
			(*result)[authority.ID] = authority
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

	count := len(*userList)
	var sw sync.WaitGroup
	sw.Add(count)
	finishChan := make(chan interface{}, 1)
	errorChan := make(chan error, count)

	userPolicys := make([]*dto.UserPolicy, 0, count)

	log.Debugf("开始获取用户策略")
	startNanosecond := time.Now().Nanosecond()

	for _, user := range *userList {
		go func(user *domain.SysUser) {
			ids, err := p.RoleUserDao.GetUserRoleIDs(ctx, user.ID)
			if err != nil {
				errorChan <- err
				return
			}
			userPolicys = append(userPolicys, &dto.UserPolicy{
				UserID:   user.ID,
				TenantID: user.TenantID,
				RoleList: *ids,
			})
			sw.Done()
		}(user)
	}

	go func() {
		sw.Wait()
		log.Debugf("用户策略获取完成，用时%d毫秒", (time.Now().Nanosecond()-startNanosecond)/1e6)
		finishChan <- struct{}{}
		close(finishChan)
	}()

	select {
	case err := <-errorChan:
		return nil, err
	case <-finishChan:
		close(errorChan)
		result := dto.UserPolicys(userPolicys)
		return &result, nil
	}
}
