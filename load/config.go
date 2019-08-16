package load

import (
	"code-generator/asset"
	"encoding/json"
	"flag"
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
	Name            string
	Root            string
	FileName        string
	NeedModule      bool
	AppendFileName  bool
	AppendClassName bool
	CustomModule    string
}

func init() {
	initDependency := flag.Bool("init", true, "init dependency")
	flag.Parse()
	if *initDependency {
		files := []string{"dist", "tpl", "types.ini"}
		bytes, e := ioutil.ReadFile("./config.json")
		if nil != e || bytes == nil {
			files = append(files, "config.json")
		}
		for _, file := range files {
			asset.RestoreAssets("./", file)
		}
	}
	Types, _ = goconfig.LoadConfigFile("types.ini")
	data, _ := ioutil.ReadFile("./config.json")
	err := json.Unmarshal(data, &Config)
	if nil != err {
		panic(err)
	}
}
