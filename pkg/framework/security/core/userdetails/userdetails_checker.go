package userdetails

// Checker 检查加载 UserDetails 的状态
type Checker interface {
	// 检测用户状态
	Check(user UserDetails) error
}
