// author: wsfuyibing <websearch@163.com>
// date: 2021-02-01

package xapp

import (
	"github.com/kataras/iris/v12"
)

// 返回数据结构.
func WithData(ctx iris.Context, data interface{}) interface{} {
	return nil
}

// 返回错误结构.
func WithError(ctx iris.Context, code int, err error) interface{} {
	return nil
}

// 输出列表结构.
func WithList(ctx iris.Context, list interface{}) interface{} {
	return nil
}
