package enums

// DeptScope 部门Scope
type DeptScope string

const (
	// DeptScopeCurrent 当前Scope
	DeptScopeCurrent DeptScope = "0"
	// DeptScopeCurrentChild 当前Scope以及子部门Scope
	DeptScopeCurrentChild DeptScope = "1"
)

// UserDetailsModeEnum 获取用户详情模式
type UserDetailsModeEnum string

// 模式
const (
	Password UserDetailsModeEnum = "password"
	Social   UserDetailsModeEnum = "social"
)

// UserStatusEnum 用户状态
type UserStatusEnum string

// 用户状态
const (
	// 可用
	UserStatusEnable UserStatusEnum = "0"
	// 锁定
	UserStatusLock UserStatusEnum = "9"
)
