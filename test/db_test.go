package test

import (
	"code-generator/internal/app"
	"fmt"
	"testing"
)

func Test_queryTotal(t *testing.T) {
	//total := db.QueryTotal("")
	total := app.QueryTotal("sys")
	t.Log(total)
}

func Test_QueryTable(t *testing.T) {
	table := app.QueryTable("sys_user")
	t.Log(table)
}

func Test_QueryColumns(t *testing.T) {
	columns := app.QueryColumns("sys_user")
	t.Log(len(*columns))
}

func Test_QueryList(t *testing.T) {
	table, count := app.QueryList("", 1, 10)
	fmt.Println(len(*table))
	fmt.Println(count)
}
