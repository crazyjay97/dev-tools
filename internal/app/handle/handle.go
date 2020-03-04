package handle

import (
	"bufio"
	"code-generator/internal/app/base"
	"code-generator/internal/app/db"
	"code-generator/internal/app/gen"
	"encoding/json"
	"fmt"
	"github.com/goinggo/mapstructure"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func init() {
	var port int
	if 0 != base.Config.Port {
		port = base.Config.Port
	} else {
		port = 8888
	}
	http.HandleFunc("/generator/list", listHandle)
	http.HandleFunc("/generator/query/all", queryAllHandle)
	http.HandleFunc("/generator/query/columns", queryColumnsHandle)
	http.HandleFunc("/generator/gen", genHandle)
	http.HandleFunc("/generator/query/code", genCodeHandle)
	http.HandleFunc("/generator/query/tplNames", queryALlTplHandle)
	InitHttpProxy()
	fmt.Println("Server Starting")
	addrList, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, address := range addrList {
		if host, ok := address.(*net.IPNet); ok && !host.IP.IsLoopback() {
			if host.IP.To4() != nil {
				fmt.Println("View On:", "http://"+host.IP.String()+":"+strconv.Itoa(port))
			}
		}
	}
	if err := http.ListenAndServe(":"+strconv.Itoa(port), nil); nil != err {
		fmt.Println("Port Already Used，Press 'Ctrl + C' To Exit")
		for bufio.NewScanner(os.Stdin).Scan() {

		}
	}
}

type ListRequest struct {
	Page      string
	Limit     string
	TableName string
}

type ListResponse struct {
	List  *[]*db.Table `json:"list"`
	Total int          `json:"total"`
}

func listHandle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	request := new(ListRequest)
	form2Struct(&r.Form, request)
	page, _ := strconv.Atoi(request.Page)
	limit, _ := strconv.Atoi(request.Limit)
	tables, count := db.QueryList(request.TableName, page, limit)
	response := ListResponse{tables, count}
	json, e := json.Marshal(&response)
	fmt.Println(e)
	w.Write(json)
}

func queryAllHandle(w http.ResponseWriter, r *http.Request) {
	tables, count := db.QueryList("", -1, -1)
	response := ListResponse{tables, count}
	json, _ := json.Marshal(&response)
	w.Write(json)
}

func queryColumnsHandle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	params := form2Map(&r.Form)
	columns := db.QueryColumns((*params)["tableName"])
	json, _ := json.Marshal(&map[string]interface{}{"list": &columns})
	w.Write(json)
}

func genHandle(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var config gen.Config
	decoder.Decode(&config)
	gen.Gen(&config, w)
}

func genCodeHandle(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var config gen.Config
	decoder.Decode(&config)
	gen.CodeGen(&config, w)
}

func queryALlTplHandle(w http.ResponseWriter, r *http.Request) {
	var tplNames []string
	for _, tpl := range base.Config.Tpl {
		tplNames = append(tplNames, (*tpl).Name)
	}
	res, _ := json.Marshal(tplNames)
	w.Write(res)
}

func form2Map(form *url.Values) *map[string]string {
	mapForm := make(map[string]string)
	for k, v := range *form {
		mapForm[k] = strings.Join(v, "")
	}
	return &mapForm
}

func form2Struct(form *url.Values, s interface{}) {
	mapForm := make(map[string]string)
	for k, v := range *form {
		mapForm[k] = strings.Join(v, "")
	}
	mapstructure.Decode(&mapForm, s)
}
