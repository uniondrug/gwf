// author: wsfuyibing <websearch@163.com>
// date: 2021-02-02

package tests

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/uniondrug/go-iris/xapp"
)

type Data1 struct {
	Key   string
	Value string
}

func TestWithData(t *testing.T) {
	x := xapp.WithData(nil, Data1{})
	s, err := json.Marshal(x)
	if err != nil {
		t.Errorf("with data error: %s.", err)
		return
	}
	t.Logf("with data result: %s", s)
}

func TestWithError(t *testing.T) {
	x := xapp.WithError(nil, 1, errors.New("Example error"))
	s, err := json.Marshal(x)
	if err != nil {
		t.Errorf("with error: %s.", err)
		return
	}
	t.Logf("with error result: %s", s)
}

func TestWithList(t *testing.T) {
	x := xapp.WithList(nil, []Data1{Data1{}, Data1{}})
	s, err := json.Marshal(x)
	if err != nil {
		t.Errorf("with list error: %s.", err)
		return
	}
	t.Logf("with list result: %s", s)
}
