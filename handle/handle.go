package handle

import (
	"code-generator/db"
	"code-generator/gen"
	"code-generator/load"
	"encoding/json"
	"github.com/goinggo/mapstructure"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func init() {
	var port int
	if 0 != load.Config.Port {
		port = load.Config.Port
	} else {
		port = 8888
	}
	http.HandleFunc("/generator/list", listHandle)
	http.HandleFunc("/generator/query/all", queryAllHandle)
	http.HandleFunc("/generator/query/columns", queryColumnsHandle)
	http.HandleFunc("/generator/gen", genHandle)
	InitHttpProxy()

	if err := http.ListenAndServe("127.0.0.1:"+strconv.Itoa(port), nil); nil != err {
		panic("Port Already Used")
		panic(err)
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
	json, _ := json.Marshal(&response)
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
