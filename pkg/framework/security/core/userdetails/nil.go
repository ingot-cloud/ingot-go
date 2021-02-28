package userdetails

type NilUserDetailsService struct{}

// LoadUserByUsername 加载指定 username 的用户
func (*NilUserDetailsService) LoadUserByUsername(username string) (UserDetails, error) {
	return nil, nil
}
