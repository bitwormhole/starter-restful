package vo

import (
	"time"

	"github.com/bitwormhole/starter/util"
)

// BaseVO 定义基本的视图对象
type BaseVO struct {
	Message       string    `json:"message"`
	Error         string    `json:"error"`
	Status        int       `json:"status"`
	Time          util.Time `json:"time"`
	TimeFormatted time.Time `json:"timef"`
}
