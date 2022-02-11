package vo

import (
	"time"

	"github.com/bitwormhole/starter/util"
)

// Meta 定义VO 的元数据
type Meta struct {
	Message       string    `json:"message"`
	Error         string    `json:"error"`
	Status        int       `json:"status"`
	Time          util.Time `json:"time"`
	TimeFormatted time.Time `json:"timef"`
}

// ViewObject 抽象的VO接口
type ViewObject interface {
	GetMeta() *Meta
}

// BaseVO 定义基本的视图对象
type BaseVO struct {
	Meta
}

func (inst *BaseVO) _Impl() ViewObject {
	return inst
}

// GetMeta ...
func (inst *BaseVO) GetMeta() *Meta {
	return &inst.Meta
}
