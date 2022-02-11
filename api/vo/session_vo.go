package vo

import "github.com/bitwormhole/starter-restful/api/dto"

// Session ...
type Session struct {
	BaseVO
	Session dto.Session `json:"session"`
}
