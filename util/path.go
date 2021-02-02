// author: wsfuyibing <websearch@163.com>
// date: 2021-02-02

package util

import (
	"os"
	"strings"
)

type Path struct {
	abs  string
	dirs []string
}

func NewPath() *Path {
	p := &Path{}
	p.dirs = []string{
		"app/controllers",
		"app/logics",
		"app/middlewares",
		"app/models",
		"app/services",
		"config",
	}
	return p
}

func (o *Path) Build() {
	if o.abs == "" {
		o.abs, _ = os.Getwd()
	}
	for _, path := range o.dirs {
		o.create(path)
	}
}

func (o *Path) SetAbsolute(abs string) *Path {
	o.abs = abs
	return o
}

func (o *Path) create(path string) {
	p := o.abs
	for _, n := range strings.Split(path, "/") {
		p += "/" + n
		if err := os.Mkdir(p, os.ModePerm); err != nil {
			println("error: ", err.Error())
			continue
		}
		println("Create: ", p)
	}
}
