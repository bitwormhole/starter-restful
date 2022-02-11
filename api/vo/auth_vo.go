package vo

import "github.com/bitwormhole/starter-restful/api/dto"

// Auth ...
type Auth struct {
	BaseVO
	Auth dto.Auth `json:"auth"`
}
