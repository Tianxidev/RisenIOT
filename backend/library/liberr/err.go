package liberr

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

// ErrIsNil 判断错误是否为空
func ErrIsNil(ctx context.Context, err error, msg ...string) {
	if !g.IsNil(err) {
		if len(msg) > 0 {
			g.Log().Error(ctx, err.Error())
			panic(msg[0])
		} else {
			panic(err.Error())
		}
	}
}

// ValueIsNil 判断值是否为空
func ValueIsNil(ctx context.Context, value interface{}, msg string) {
	if g.IsNil(value) {
		g.Log().Error(ctx, msg)
		panic(msg)
	}
}

// ValueIsTrue 判断值是否为真
func ValueIsTrue(ctx context.Context, value bool, msg string) {
	if value {
		g.Log().Error(ctx, msg)
		panic(msg)
	}
}

// ValueIsFalse 判断值是否为假
func ValueIsFalse(ctx context.Context, value bool, msg string) {
	if !value {
		g.Log().Error(ctx, msg)
		panic(msg)
	}
}
