// author: wsfuyibing <websearch@163.com>
// date: 2021-02-06

package xredis

import (
	"io/ioutil"
	"time"

	"github.com/gomodule/redigo/redis"
	"gopkg.in/yaml.v2"

	"github.com/uniondrug/gwf/xlog"
)

// Redis配置.
type configuration struct {
	Addr      string `yaml:"addr"`
	Index     int    `yaml:"index"`
	Password  string `yaml:"password"`
	MinIdle   int    `yaml:"min-idle"`
	MaxIdle   int    `yaml:"max-idle"`
	MaxActive int    `yaml:"max-active"`
	Timeout   int    `yaml:"timeout"`
	pools     *redis.Pool
}

// Load configuration from yaml file.
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
	xlog.Infof("load config from %s.", path)
	o.parse()
	return nil
}

// init config with default file.
func (o *configuration) onInit() {
	for _, path := range []string{"./config/redis.yaml", "../config/redis.yaml"} {
		err := o.LoadYaml(path)
		if err == nil {
			break
		}
	}
}

// parse marshal.
func (o *configuration) parse() {
	// Redis连接池.
	o.pools = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				"tcp", o.Addr,
				redis.DialPassword(o.Password),
				redis.DialDatabase(o.Index),
			)
		},
		MaxIdle:         o.MaxIdle,
		MaxActive:       o.MaxActive,
		IdleTimeout:     time.Duration(o.Timeout) * time.Second,
		Wait:            false,
		MaxConnLifetime: time.Duration(o.Timeout) * time.Second,
	}
}
