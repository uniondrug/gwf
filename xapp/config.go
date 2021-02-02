// author: wsfuyibing <websearch@163.com>
// date: 2021-02-02

package xapp

import (
	"io/ioutil"
	"sync"

	"gopkg.in/yaml.v2"

	"github.com/uniondrug/gwf/xlog"
)

const (
	DefaultAddr    = ":8080"
	DefaultName    = "sketch"
	DefaultVersion = "0.0"
)

var Config *config

type config struct {
	Addr    string `yaml:"Addr"`
	Name    string `yaml:"Name"`
	Version string `yaml:"Version"`
}

func (o *config) LoadYaml(path string) error {
	data, err := ioutil.ReadFile(path)
	// return if read file error.
	if err != nil {
		return err
	}
	// return if parse yaml error.
	if err = yaml.Unmarshal(data, o); err != nil {
		return err
	}
	// parse config.
	xlog.Infof("load config from %s.", path)
	return nil
}

func (o *config) onInit() {
	for _, path := range []string{"./config/app.yaml", "../config/app.yaml"} {
		if err := o.LoadYaml(path); err == nil {
			break
		}
	}
	if o.Addr == "" {
		o.Addr = DefaultAddr
	}
	if o.Name == "" {
		o.Name = DefaultName
	}
	if o.Version == "" {
		o.Version = DefaultVersion
	}

}

func init() {
	new(sync.Once).Do(func() {
		xlog.Info("initialize golang framework application.")
		Config = new(config)
		Config.onInit()
	})
}
