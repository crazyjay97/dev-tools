package gen

import (
	"archive/zip"
	"code-generator/db"
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
	template, _ := pongo2.FromFile("./tpl/index.tpl")
	zipW := zip.NewWriter(w)
	for _, module := range config.Modules {
		columns := db.QueryColumns(module.TableName)
		table := db.QueryTable(module.TableName)
		rs, _ := template.Execute(map[string]interface{}{
			"columns": columns,
			"table":   table,
		})
		fW, _ := zipW.Create(getPath(table.ModuleName, table.FileName, "index.vue"))
		fW.Write([]byte(rs))
	}
	defer func() {
		zipW.Close()
	}()
}

func getPath(moduleName, pageName, fileName string) string {
	mainPath := "code"
	return strings.Join([]string{
		mainPath,
		moduleName,
		pageName,
		fileName,
	}, "/")
}
