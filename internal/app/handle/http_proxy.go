package handle

import (
	"code-generator/internal/app/base"
	"code-generator/internal/common/utils"
	"log"
	"net/http"
	"strings"
)

const DefaultPage = "index.html"
const AssetRootPath = "asset/dist"

func InitHttpProxy() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		path := request.URL.Path
		if path == "/" {
			path += DefaultPage
		}
		path = AssetRootPath + path
		log.Print(path)
		bytes, _ := utils.GetFileInProject(path)
		suffix := path[strings.LastIndex(path, ".")+1:]
		fileType, err := base.Types.GetValue("mime", suffix)
		if nil == err && suffix != "" {
			writer.Header().Set("Content-Type", fileType)
		} else {
			writer.Header().Set("Content-Type", "text/plain")
		}
		writer.Write(bytes)
	})
}
