package dto

import "github.com/bitwormhole/starter/util"

// Auth ...
type Auth struct {
	BaseDTO

	Account   string      `json:"account"`   // 用户的标识，可以是 UserID, UserName, Email, Phone ... 等
	Mechanism string      `json:"mechanism"` // 登录机制
	Secret    util.Base64 `json:"secret"`    // 可以是，但不局限于密码等，取决于具体的机制

	// 以下表示返回结果
	Success  bool   `json:"success"`  // 表示已登录成功
	WantMore bool   `json:"wantmore"` // 指出需要更多的步骤
	Want     string `json:"want"`     // 指出需要进一步提供的证明
}
