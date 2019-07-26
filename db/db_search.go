package db

import (
	"code-generator/load"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"strings"
	"time"
)

type Table struct {
	TableName    string `json:"tableName"`
	ModuleName   string
	FileName     string
	ClassName    string
	Engine       string   `json:"engine"`
	TableComment string   `json:"tableComment"`
	CreateTime   UnixTime `json:"createTime"`
}

type UnixTime time.Time

func (t UnixTime) MarshalJSON() ([]byte, error) {
	format := time.Time(t).Format("2006/1/2 15:04:05")
	format = "\"" + format + "\""
	return []byte(format), nil
}

func (table *Table) Parse() {
	table.TableName = strings.ToLower(table.TableName)
	splits := strings.Split(table.TableName, "_")
	table.ModuleName = splits[0]
	splits = splits[1:]
	for i, str := range splits {
		if i != 0 {
			splits[i] = strings.ToUpper(string(str[0])) + string(str[1:])
		}
	}
	table.FileName = strings.Join(splits, "")
	table.ClassName = strings.ToUpper(string(table.FileName[0])) + table.FileName[1:]

}

type Column struct {
	ColumnName      string `json:"columnName"`
	FieldName       string
	DataType        string `json:"dataType"`
	JavaType        string
	ColumnComment   string `json:"columnComment"`
	ColumnKey       string `json:"columnKey"`
	Extra           string `json:"extra"`
	NeedShow        bool   //是否需要展示
	NeedAdd         bool   // 是否需要添加
	NeedFilter      bool   // 是否需要过滤
	ShowMode        int    // 展示方式
	DictionaryKey   string // 数据字典key
	DictionaryLabel string // 数据字典label
	DictionaryValue string // 数据字典value
	IsJoinColumn    bool   //是否关联字段
}

type Dictionary struct {
	Id         string
	TenantId   string
	CodeType   string
	CodeValue  string
	CodeText   string
	CodeName   string
	IsCommon   int
	State      int
	OrderNum   float32
	CreateBy   string
	CreateTime time.Time
	UpdateBy   string
	UpdateTime time.Time
	Deleted    int
}

func (Dictionary) TableName() string {
	return "SYS_DICTIONARY"
}

func (dict *Dictionary) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("ID", uuid)
}

func (column *Column) Parse() {
	splits := strings.Split(strings.ToLower(column.ColumnName), "_")
	column.FieldName = splits[0]
	splits = splits[1:]
	for i, str := range splits {
		splits[i] = strings.ToUpper(string(str[0])) + string(str[1:])
	}
	column.FieldName += strings.Join(splits, "")
	javaType, err := load.Types.GetValue("javatype", column.DataType)
	column.JavaType = javaType
	if err != nil {
		column.JavaType = "Object"
	}
}

type page struct {
	index  int //开始页
	offset int //开始条数
	limit  int //条数
}

func NewPage(index int, limit int) *page {
	return &page{index: index, offset: (index - 1) * limit, limit: limit}
}

func QueryTotal(tableName string) int {
	db := DB.Table("information_schema.tables").Select("COUNT(*)").Where("table_schema = (select database())")
	if "" != tableName {
		db = db.Where("table_name like concat('%', ?, '%')", tableName)
	}
	var total int
	db.Count(&total)
	return total
}

func QueryTable(tableName string) *Table {
	table := new(Table)
	DB.Table("information_schema.tables").
		Select("table_name tableName, engine, table_comment tableComment, create_time createTime").
		Where("table_schema = (select database()) and table_name = ?", tableName).
		Row().Scan(&table.TableName, &table.Engine, &table.TableComment, &table.CreateTime)
	table.Parse()
	return table
}

func QueryColumns(tableName string) *[]*Column {
	columns := make([]*Column, 0)
	rows, e := DB.Table("information_schema.columns").
		Select("column_name ColumnName, data_type dataType, column_comment columnComment, column_key columnKey, extra").
		Where("table_name = ? and table_schema = (select database())", tableName).Order("ordinal_position").Rows()
	if nil == e {
		for rows.Next() {
			column := new(Column)
			rows.Scan(&column.ColumnName, &column.DataType, &column.ColumnComment, &column.ColumnKey, &column.Extra)
			column.Parse()
			columns = append(columns, column)
		}
	} else {
		panic(e)
	}
	defer rows.Close()
	return &columns
}

func QueryList(tableName string, index int, limit int) (*[]*Table, int) {
	var count int
	db := DB.Table("information_schema.tables").
		Select("table_name tableName, engine, table_comment tableComment, create_time createTime").
		Where("table_schema = (select database())")
	if "" != tableName {
		db = db.Where("table_name like concat('%', ?, '%')", tableName)
	}
	db.Count(&count)
	if index != -1 && limit != -1 {
		page := NewPage(index, limit)
		db = db.Offset(page.offset).Limit(page.limit)
	}
	rows, e := db.Rows()
	tables := make([]*Table, 0)
	if nil == e {
		for rows.Next() {
			table := new(Table)
			rows.Scan(&table.TableName, &table.Engine, &table.TableComment, &table.CreateTime)
			tables = append(tables, table)
		}
	} else {
		goto CloseRows
		panic(e)
	}
	defer rows.Close()
CloseRows:
	rows.Close()
	return &tables, count
}
