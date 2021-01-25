package userdetails

// Service 用于加载 UserDetails
type Service interface {
	// 加载指定 username 的用户
	LoadUserByUsername(username string) (UserDetails, error)
}
