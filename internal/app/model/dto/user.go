package dto

import (
	"github.com/ingot-cloud/ingot-go/internal/app/model/enums"
	"time"
)

// Users for user array
type Users []*User

// User dto
type User struct {
	ID            string     `json:"id"`
	Username      string     `json:"username"`
	Password      string     `json:"password"`
	RealName      string     `json:"realName"`
	Phone         string     `json:"phone"`
	Status        string     `json:"status"`
	Remark        string     `json:"remark"`
	LastLoginTime *time.Time `json:"last_login_time"`
	Creator       string     `json:"creator"`
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     time.Time  `json:"updateAt"`
}

// UserQueryParams params
type UserQueryParams struct {
	Username string             `json:"username" form:"username"`
	Status   enums.CommonStatus `json:"status" form:"status"`
}

// UserPageQueryParams params
type UserPageQueryParams struct {
	Pagination
	UserQueryParams
}

// UserQueryResult response
type UserQueryResult struct {
	List Users
}

// StaffLoginInfo 员工登录信息
type StaffLoginInfo struct {
	UserID       string `json:"userID"`
	EnterpriseID string `json:"enterpriseID"`
	Username     string `json:"username"`
	Password     string `json:"password"`
}

// StaffParams 员工
type StaffParams struct {
	EnterpriseID string `json:"enterpriseID"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	RoleID       string `json:"roleID"`
	Remark       string `json:"remark"`
}

// StaffQueryParams 查询员工参数
type StaffQueryParams struct {
	EnterpriseID string             `json:"enterpriseID" form:"enterpriseID"`
	Number       string             `json:"number" form:"number"`
	Username     string             `json:"username" form:"username"`
	Status       enums.CommonStatus `json:"status" form:"status"`
	Pagination
}

// StaffLoginInfoQuery 员工登录信息查询参数
type StaffLoginInfoQuery struct {
	EnterpriseID string `json:"enterpriseID" form:"enterpriseID"`
}

// FixPasswordParams 修改密码参数
type FixPasswordParams struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}
