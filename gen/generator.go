package gen

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

func Gen(config *Config) {

}
