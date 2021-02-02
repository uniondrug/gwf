// author: wsfuyibing <websearch@163.com>
// date: 2021-02-02

package tests

import (
	"testing"

	"gwf/xapp"
)

type ExampleRequest struct {
	Age int `validate:"required,gte=14,lte=60" label:"年龄"`
	Name string `validate:"required,min=3,max=12"`
}

func TestValidate(t *testing.T) {

	s := &ExampleRequest{
		Age: 19,
		Name:"chinese",
	}

	if err := xapp.Validate(s); err != nil {
		t.Errorf("validate error: %s.", err)
		return
	}

	t.Logf("validate succeed.")

}
