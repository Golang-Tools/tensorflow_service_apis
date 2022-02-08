package tensorflow_service_apis

import (
	"time"
)

//CtxOptions 设置ctx行为的选项
type CtxOptions struct {
	UntilEnd bool
	Timeout  time.Duration
}

var DefaultCtxOpts = CtxOptions{}

// CtxOption configures how we set up the connection.
type CtxOption interface {
	Apply(*CtxOptions)
}

// func (emptyOption) apply(*Options) {}
type funcCtxOption struct {
	f func(*CtxOptions)
}

func (fo *funcCtxOption) Apply(do *CtxOptions) {
	fo.f(do)
}

func newFuncCtxOption(f func(*CtxOptions)) *funcCtxOption {
	return &funcCtxOption{
		f: f,
	}
}

//UntilEnd NewCtx方法的参数,用于设置ctx为不会超时
func UntilEnd() CtxOption {
	return newFuncCtxOption(func(o *CtxOptions) {
		o.UntilEnd = true
	})
}

//WithTimeout NewCtx方法的参数,用于设置ctx为指定的超时时长
func WithTimeout(timeout time.Duration) CtxOption {
	return newFuncCtxOption(func(o *CtxOptions) {
		o.Timeout = timeout
	})
}
