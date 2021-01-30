// author: wsfuyibing <websearch@163.com>
// date: 2021-01-29

package gwf

import (
	"github.com/kataras/iris/v12"
)

// 返回给IRIS, 标准的Data结构.
func WithData(ctx iris.Context, data interface{}) interface{} {
	return iris.Map{
		"errno":    0,
		"error":    "",
		"dataType": "OBJECT",
		"data":     data,
	}
}

// 返回给IRIS, 标准的错误返回结构.
func WithError(ctx iris.Context, code int, err error) interface{} {
	return iris.Map{
		"errno":    code,
		"error":    err.Error(),
		"dataType": "OBJECT",
		"data":     iris.Map{},
	}
}

// 返回给IRIS, 标准的数据列表.
func WithList(ctx iris.Context, data interface{}) interface{} {
	return iris.Map{
		"errno":    0,
		"error":    "",
		"dataType": "LIST",
		"data":     data,
	}
}
