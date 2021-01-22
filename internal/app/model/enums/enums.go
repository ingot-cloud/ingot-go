package enums

// DeptScope 部门Scope
type DeptScope string

const (
	// DeptScopeCurrent 当前Scope
	DeptScopeCurrent DeptScope = "0"
	// DeptScopeCurrentChild 当前Scope以及子部门Scope
	DeptScopeCurrentChild DeptScope = "1"
)
