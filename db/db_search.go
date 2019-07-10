package db

type Table struct {
	TableName    string `json:"tableName"`
	Engine       string `json:"engine"`
	TableComment string `json:"tableComment"`
	CreateTime   string `json:"createTime"`
}

type Column struct {
	ColumnName    string `json:"columnName"`
	DataType      string `json:"dataType"`
	ColumnComment string `json:"columnComment"`
	ColumnKey     string `json:"columnKey"`
	Extra         string `json:"extra"`
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
