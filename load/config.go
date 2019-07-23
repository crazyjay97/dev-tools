package load

import (
	"encoding/json"
	"github.com/Unknwon/goconfig"
	"io/ioutil"
)

var Config JsonConfig

var Types *goconfig.ConfigFile

type JsonConfig struct {
	Port int
	Db   *Db
	Tpl  []*Tpl
}
type Db struct {
	Dialect     string
	Url         string
	MaxIdleCons int
	MaxOpenCon  int
	PrintLog    bool
}
type Tpl struct {
	Name           string
	Root           string
	FileName       string
	NeedModule     bool
	AppendFileName bool
}

func init() {
	Types, _ = goconfig.LoadConfigFile("types.ini")
	data, _ := ioutil.ReadFile("./config.json")
	err := json.Unmarshal(data, &Config)
	if nil != err {
		panic(err)
	}
}
