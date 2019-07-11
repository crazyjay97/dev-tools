package gen

import (
	"archive/zip"
	"code-generator/db"
	"code-generator/load"
	"github.com/flosch/pongo2"
	"net/http"
	"strings"
)

type Config struct {
	Modules               []Module
	MainPath              string
	PackageName           string
	AuthorName            string
	EmailAddress          string
	AutoSettingModuleName bool
	ModuleName            string
	RemovePrefix          bool
}

type Module struct {
	TableName    string      //表名
	SearchFields []string    //查询字段
	AddFields    []string    //新增字段
	JoinTables   []JoinTable //关联表
}

type JoinTable struct {
	TableName   string //表名
	SelfColumn  string //关联字段
	JoinColumn  string //被关联字段
	Alias       string //别名
	Description string //描述
}

func Gen(config *Config, w http.ResponseWriter) {
	zipW := zip.NewWriter(w)
	tpls := load.Config.Tpl
	for _, module := range config.Modules {
		columns := db.QueryColumns(module.TableName)
		table := db.QueryTable(module.TableName)
		data := map[string]interface{}{
			"columns":   columns,
			"table":     table,
			"addFields": module.AddFields,
		}
		for _, tpl := range tpls {
			template, _ := pongo2.FromFile("./tpl/" + tpl.Name + ".tpl")
			rs, _ := template.Execute(data)
			fW, _ := zipW.Create(getPath(tpl.Root, table.ModuleName, table.FileName, tpl.FileName, tpl.NeedModule))
			fW.Write([]byte(rs))
		}
	}
	defer func() {
		zipW.Close()
	}()
}

func getPath(root, moduleName, pageName, fileName string, needPageModule bool) string {
	mainPath := "code"
	if !needPageModule {
		fileName = ""
	}
	return strings.Join([]string{
		mainPath,
		root,
		moduleName,
		pageName,
		fileName,
	}, "/")
}
