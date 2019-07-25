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
	TableName     string   //表名
	SearchFields  []string //查询字段
	AddFields     []string //新增字段
	ColumnSetting []ColumnSetting
	JoinTables    []JoinTable //关联表
}

type ColumnSetting struct {
	Column          string //字段名
	ColumnDesc      string //介绍
	NeedShow        bool   //是否需要展示
	NeedAdd         bool   // 是否需要添加
	NeedFilter      bool   // 是否需要过滤
	ShowMode        int    // 展示方式
	DictionaryLabel string // 数据字典key
	DictionaryValue string // 数据字典value
}

type JoinTable struct {
	TableName    string //表名
	SelfColumn   string //关联字段
	JoinColumn   string //被关联字段
	SearchColumn string //别名
	Alias        string //别名
	FieldName    string //别名
	Description  string //描述
}

func (joinTable *JoinTable) parse() {
	splits := strings.Split(strings.ToLower(joinTable.Alias), "_")
	joinTable.FieldName = splits[0]
	splits = splits[1:]
	for i, str := range splits {
		splits[i] = strings.ToUpper(string(str[0])) + string(str[1:])
	}
	joinTable.FieldName += strings.Join(splits, "")
}

func Gen(config *Config, w http.ResponseWriter) {
	zipW := zip.NewWriter(w)
	tpls := load.Config.Tpl
	for _, module := range config.Modules {
		columns := db.QueryColumns(module.TableName)
		table := db.QueryTable(module.TableName)
		addColumns, searchColumns := filterColumns(columns, &module.ColumnSetting)
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
		table.parse()
		realColumns := db.QueryColumns(table.TableName)
		var currentColumn db.Column
		for _, column := range *realColumns {
			if column.ColumnName == table.JoinColumn {
				currentColumn = *column
			}
		}
		*columns = append(*columns, &db.Column{NeedShow: true, ColumnName: table.Alias, FieldName: table.FieldName, ColumnComment: table.Description, Extra: currentColumn.Extra, DataType: currentColumn.DataType})
	}
}

func filterColumns(columns *[]*db.Column, columnSetting *[]ColumnSetting) (*[]*db.Column, *[]*db.Column) {
	addColumns := make([]*db.Column, 0)
	searchColumns := make([]*db.Column, 0)
	for _, column := range *columns {
		for _, setting := range *columnSetting {
			if setting.Column == column.ColumnName {
				column.ColumnComment = setting.ColumnDesc
				column.NeedShow = setting.NeedShow
				column.NeedAdd = setting.NeedAdd
				column.NeedFilter = setting.NeedFilter
				column.ShowMode = setting.ShowMode
				column.DictionaryLabel = setting.DictionaryLabel
				column.DictionaryValue = setting.DictionaryValue
				if setting.NeedAdd {
					addColumns = append(addColumns, column)
				}
				if setting.NeedFilter {
					searchColumns = append(searchColumns, column)
				}
			}
		}
	}
	return &addColumns, &searchColumns
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
