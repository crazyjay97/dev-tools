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
		addColumns := filterColumns(columns, &module.AddFields)
		searchColumns := filterColumns(columns, &module.SearchFields)
		listColumns := (*columns)[0:]
		appendColumn(&listColumns, &module.JoinTables)
		data := map[string]interface{}{
			"columns":       columns,
			"table":         table,
			"addFields":     module.AddFields,
			"addColumns":    addColumns,
			"searchColumns": searchColumns,
			"listColumns":   listColumns, //查询列表,包含需要关联查的字段
			"joinTables":    module.JoinTables,
			"moduleName":    table.ModuleName,
			"fileName":      table.FileName,
			"className":     table.ClassName,
			"packageName":   config.PackageName,
		}
		for _, tpl := range tpls {
			template, _ := pongo2.FromFile("./tpl/" + tpl.Name + ".tpl")
			rs, _ := template.Execute(data)
			fW, _ := zipW.Create(getPath(tpl.Root, table.ModuleName, table.FileName, tpl.FileName, tpl.NeedModule, tpl.AppendFileName))
			fW.Write([]byte(rs))
		}
	}
	defer func() {
		zipW.Close()
	}()
}

func appendColumn(columns *[]*db.Column, joinTables *[]JoinTable) {
	for _, table := range *joinTables {
		realColumns := db.QueryColumns(table.TableName)
		var currentColumn db.Column
		for _, column := range *realColumns {
			if column.ColumnName == table.JoinColumn {
				currentColumn = *column
			}
		}
		*columns = append(*columns, &db.Column{ColumnName: table.SelfColumn, FieldName: table.Alias, ColumnComment: table.Description, Extra: currentColumn.Extra, DataType: currentColumn.DataType})
	}
}

func filterColumns(columns *[]*db.Column, addFields *[]string) *[]*db.Column {
	newColumn := make([]*db.Column, 0)
	for _, column := range *columns {
		for _, field := range *addFields {
			if field == column.ColumnName {
				newColumn = append(newColumn, column)
			}
		}
	}
	return &newColumn
}

func getPath(root, moduleName, pageName, fileName string, needPageModule bool, appendFileName bool) string {
	mainPath := "code"
	if appendFileName {
		fileName = pageName + fileName
	}
	if !needPageModule {
		pageName = ""
	}
	return strings.Join([]string{
		mainPath,
		root,
		moduleName,
		pageName,
		fileName,
	}, "/")
}
