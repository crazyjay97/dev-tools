package gen

import (
	"archive/zip"
	"code-generator/internal/app/base"
	"code-generator/internal/app/db"
	"code-generator/internal/common/utils"
	"github.com/flosch/pongo2"
	"net/http"
	"strings"
	"time"
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
	tpls := base.Config.Tpl
	for _, module := range config.Modules {
		columns := db.QueryColumns(module.TableName)
		table := db.QueryTable(module.TableName)
		addColumns, searchColumns, pkColumn := filterColumns(columns, &module.ColumnSetting, config)
		listColumns := (*columns)[0:]
		appendColumn(&listColumns, &module.JoinTables)
		hasBigDecimal, hasDate, hasTime := searchSpecialType(&listColumns)
		hasJoinColumn := len(module.JoinTables) > 0
		data := map[string]interface{}{
			"columns":       columns,
			"pkColumn":      pkColumn,
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
			"authorName":    config.AuthorName,
			"emailAddress":  config.EmailAddress,
			"mainPath":      config.MainPath,
			"genTime":       time.Now().Format("2006-01-02 15:04:05"),
			"hasBigDecimal": hasBigDecimal,
			"hasDate":       hasDate,
			"hasTime":       hasTime,
			"hasJoinColumn": hasJoinColumn,
		}
		for _, tpl := range tpls {
			bytes, _ := utils.GetFileInProject("asset/tpl/" + tpl.Name + ".tpl")
			template, _ := pongo2.FromString(string(bytes))
			rs, _ := template.Execute(data)
			fW, _ := zipW.Create(getPath(tpl, table.ModuleName, table.FileName, table.ClassName))
			fW.Write([]byte(rs))
		}
	}
	defer func() {
		zipW.Close()
	}()
}

func searchSpecialType(columns *[]*db.Column) (bool, bool, bool) {
	hasBigDecimal := false
	hasDate := false
	hasTime := false
	for _, column := range *columns {
		switch column.JavaType {
		case "BigDecimal":
			hasBigDecimal = true
		case "Date":
			hasDate = true
		case "Time":
			hasTime = true
		}
	}
	return hasBigDecimal, hasDate, hasTime
}

func appendColumn(columns *[]*db.Column, joinTables *[]JoinTable) {
	newColumns := make([]*db.Column, 0)
	for _, table := range *joinTables {
		table.parse()
		realColumns := db.QueryColumns(table.TableName)
		var currentColumn db.Column
		for _, column := range *realColumns {
			if column.ColumnName == table.JoinColumn {
				currentColumn = *column
			}
		}
		newColumns = append(newColumns,
			&db.Column{IsJoinColumn: true,
				NeedShow:      true,
				ColumnName:    table.Alias,
				FieldName:     table.FieldName,
				ColumnComment: table.Description,
				Extra:         currentColumn.Extra,
				DataType:      currentColumn.DataType,
				JavaType:      currentColumn.JavaType})
	}
	*columns = append(newColumns, *columns...)
}

func filterColumns(columns *[]*db.Column, columnSetting *[]ColumnSetting, config *Config) (*[]*db.Column, *[]*db.Column, *db.Column) {
	addColumns := make([]*db.Column, 0)
	searchColumns := make([]*db.Column, 0)
	var pkColumn db.Column
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
				if "" != setting.DictionaryLabel {
					dictionaryKeyAndLabel := strings.Split(setting.DictionaryLabel, ":")
					column.DictionaryKey = dictionaryKeyAndLabel[0]
					if "" != setting.DictionaryValue {
						//todo
						//genDictionary(dictionaryKeyAndLabel[0], dictionaryKeyAndLabel[1], setting.DictionaryValue, *config)
					}
				}
				if column.ColumnKey == "PRI" {
					pkColumn = *column
				}
				if setting.NeedAdd {
					addColumns = append(addColumns, column)
				}
				if setting.NeedFilter {
					searchColumns = append(searchColumns, column)
				}
			}
		}
	}
	return &addColumns, &searchColumns, &pkColumn
}

func genDictionary(key, label, dictionaries string, config Config) {
	dictList := strings.Split(dictionaries, ",")
	for _, dict := range dictList {
		valueAndDesc := strings.Split(dict, ":")
		count := 0
		db.DB.Model(&db.Dictionary{}).Where(&db.Dictionary{CodeType: key, CodeValue: valueAndDesc[0]}).Count(&count)
		if count == 0 {
			dictionary := db.Dictionary{CodeType: key, CodeName: label, CodeValue: valueAndDesc[0],
				CodeText: valueAndDesc[1], CreateBy: config.AuthorName, CreateTime: time.Now()}
			db.DB.Create(dictionary)
		}
	}

}

func getPath(tpl *base.Tpl, moduleName string, pageName string, className string) string {
	mainPath := "code"
	fileName := tpl.FileName
	if tpl.AppendFileName {
		fileName = pageName + tpl.FileName
	}
	if tpl.AppendClassName {
		fileName = className + tpl.FileName
	}
	if !tpl.NeedModule {
		pageName = ""
	}
	return strings.Join([]string{
		mainPath,
		tpl.Root,
		moduleName,
		pageName,
		tpl.CustomModule,
		fileName,
	}, "/")
}
