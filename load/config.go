package load

import (
	"encoding/json"
	"io/ioutil"
)

var Config JsonConfig

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
	data, _ := ioutil.ReadFile("./config.json")
	err := json.Unmarshal(data, &Config)
	if nil != err {
		panic(err)
	}
}
