package handle

import (
	"code-generator/internal/app/base"
	"code-generator/internal/common/utils"
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
		bytes, err := utils.GetFileInProject(path)
		if nil != err {
			writer.WriteHeader(http.StatusNotFound)
			writer.Write([]byte(`
				<body style="text-align:center">
					<h1>内容失踪了</h1>
                </body>
			`))
			return
		}
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
