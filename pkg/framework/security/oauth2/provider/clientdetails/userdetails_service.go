package clientdetails

import (
	"reflect"

	"github.com/ingot-cloud/ingot-go/pkg/framework/log"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/userdetails"
)

// UserDetailsService client实现
type UserDetailsService struct {
	ClientDetailsService Service
}

// NewClientDetailsUserDetailsService 实例化
func NewClientDetailsUserDetailsService(service Service) *UserDetailsService {
	return &UserDetailsService{
		ClientDetailsService: service,
	}
}

// LoadUserByUsername 加载指定 username 的用户
func (s *UserDetailsService) LoadUserByUsername(username string) (userdetails.UserDetails, error) {
	log.Errorf("LoadUserByUsername clientDetailsService=%s", reflect.TypeOf(s.ClientDetailsService))
	clientDetails, err := s.ClientDetailsService.LoadClientByClientID(username)
	if err != nil {
		return nil, err
	}

	return userdetails.NewUser(username, clientDetails.GetClientSecret(), clientDetails.GetAuthorities()), nil
}
