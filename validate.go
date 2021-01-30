// author: wsfuyibing <websearch@163.com>
// date: 2021-01-29

package gwf

// 校验结构体字段值.
// 限结构体类型入参, 并且字段有 `validate:""` 限定.
//
//   type ExampleRequest struct {
//       TopicName string `json:"topic_name" validate:"required,min=3,max=16" label:"主题名"`
//   }
func Validate(req interface{}) error {
	return nil
}
