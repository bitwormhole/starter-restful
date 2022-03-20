package vo

import (
	"time"

	"github.com/bitwormhole/starter/util"
)

// Meta 定义VO 的元数据
type Meta struct {
	Error         string            `json:"error"`
	Message       string            `json:"message"`
	Params        map[string]string `json:"params"`
	Status        int               `json:"status"`
	Time          util.Time         `json:"time"`
	TimeFormatted time.Time         `json:"timef"`
}

////////////////////////////////////////////////////////////////////////////////

// ViewObject 抽象的VO接口
type ViewObject interface {
	GetMeta() *Meta
}

////////////////////////////////////////////////////////////////////////////////

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

// Param 取参数的值
func (inst *BaseVO) Param(name string) string {
	t := inst.Params
	if t == nil {
		return ""
	}
	return t[name]
}

// SetParam 设置参数
func (inst *BaseVO) SetParam(name, value string) {
	t := inst.Params
	if t == nil {
		t = make(map[string]string)
		inst.Params = t
	}
	t[name] = value
}

////////////////////////////////////////////////////////////////////////////////
