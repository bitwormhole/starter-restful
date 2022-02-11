package dto

import "github.com/bitwormhole/starter/util"

// Session ...
type Session struct {
	BaseDTO

	Email       string `json:"email"`    // 用户的邮件地址（可能为空）
	Phone       string `json:"phone"`    // 用户的电话号码（可能为空）
	Avatar      string `json:"avatar"`   // 用户的头像（HTTP-URL）
	DisplayName string `json:"nickname"` // 用户的昵称
	ID          string `json:"userid"`   // 用户ID
	Name        string `json:"username"` // 用户名

	CreatedAt     util.Time `json:"createdat"`     // 会话的创建时间，通常等于登录的时间
	UpdatedAt     util.Time `json:"updatedat"`     // 会话的更新时间
	Authenticated bool      `json:"authenticated"` // 表示用户已登录
}
