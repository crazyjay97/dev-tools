package test

import (
	"code-generator/internal/app/db"
	"fmt"
	"testing"
)

func Test_queryTotal(t *testing.T) {
	//total := db.QueryTotal("")
	total := db.QueryTotal("sys")
	t.Log(total)
}

func Test_QueryTable(t *testing.T) {
	table := db.QueryTable("sys_user")
	t.Log(table)
}

func Test_QueryColumns(t *testing.T) {
	columns := db.QueryColumns("sys_user")
	t.Log(len(*columns))
}

func Test_QueryList(t *testing.T) {
	table, count := db.QueryList("", 1, 10)
	fmt.Println(len(*table))
	fmt.Println(count)
}
