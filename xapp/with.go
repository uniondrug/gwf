// author: wsfuyibing <websearch@163.com>
// date: 2021-02-01

package xapp

import (
	"github.com/kataras/iris/v12"
)

// 返回数据结构.
func WithData(ctx iris.Context, data interface{}) interface{} {
	return iris.Map{
		"errno":    0,
		"error":    "",
		"data":     data,
		"dataType": "OBJECT",
	}
}

// 返回错误结构.
func WithError(ctx iris.Context, code int, err error) interface{} {
	return iris.Map{
		"errno":    code,
		"error":    err.Error(),
		"data":     iris.Map{},
		"dataType": "ERROR",
	}
}

// 输出列表结构.
func WithList(ctx iris.Context, list interface{}) interface{} {
	return iris.Map{
		"errno":    0,
		"error":    "",
		"data":     list,
		"dataType": "LIST",
	}
}
