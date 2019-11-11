package base

import (
	"code-generator/asset"
	"code-generator/internal/common/utils"
	"encoding/json"
	"github.com/Unknwon/goconfig"
	"log"
	"os"
	"path/filepath"
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
	configJson, err := utils.GetFileInProject("configs/config.json")
	if err != nil {
		restoreDependency()
		configJson, err = utils.GetFileInProject("configs/config.json")
	}
	typesIni, _ := utils.GetFileInProject("configs/types.ini")
	Types, _ = goconfig.LoadFromData(typesIni)
	err = json.Unmarshal(configJson, &Config)
	if nil != err {
		panic(err)
	}
}

func restoreDependency() {
	log.Println("Init Static File")
	files := []string{"asset", "configs"}
	AppPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		err = asset.RestoreAssets(AppPath, file)
	}
}
