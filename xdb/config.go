// author: wsfuyibing <websearch@163.com>
// date: 2021-01-30

package xdb

import (
	"io/ioutil"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
	"xorm.io/xorm"
	"xorm.io/xorm/names"

	"gwf/xlog"
)

// 包配置.
type configuration struct {
	Driver      string   `yaml:"driver"`
	Dsn         []string `yaml:"dsn"`
	MaxIdle     int      `yaml:"max-idle"`
	MaxOpen     int      `yaml:"max-open"`
	MaxLifetime int      `yaml:"max-lifetime"`
	ShowSQL     bool     `yaml:"show-sql"`
	Mapper      string   `yaml:"mapper"`
	engines     *xorm.EngineGroup
}

// 从配置文件加载配置.
func (o *configuration) LoadYaml(path string) error {
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
	xlog.Infof("[GWF][SERVICE] load config from %s.", path)
	o.parse()
	return nil
}

// 在包初始化时调用.
func (o *configuration) onInit() {
	for _, path := range []string{"./config/service.yaml", "../config/service.yaml"} {
		err := o.LoadYaml(path)
		if err == nil {
			break
		}
	}
}

// 解析X ORM引擎.
func (o *configuration) parse() {
	// prepare.
	var err error
	if o.engines, err = xorm.NewEngineGroup(o.Driver, o.Dsn); err != nil {
		panic(err)
	}
	// initialize options.
	xlog.Infof("[GWF][SERVICE] assign %s driver with %d dsn, max idles is %d, max open files is %d.", o.Driver, len(o.Dsn), o.MaxIdle, o.MaxOpen)
	o.engines.SetConnMaxLifetime(time.Duration(o.MaxLifetime) * time.Second)
	o.engines.SetMaxIdleConns(o.MaxIdle)
	o.engines.SetMaxOpenConns(o.MaxOpen)
	o.engines.ShowSQL(o.ShowSQL)
	if o.Mapper == "same" {
		o.engines.SetColumnMapper(names.SameMapper{})
	} else if o.Mapper == "snake" {
		o.engines.SetColumnMapper(names.SnakeMapper{})
	}
}
