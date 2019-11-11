package handle

import (
	"code-generator/internal/app/base"
	"code-generator/internal/common/utils"
	"net/http"
	"os"
	"strings"
)

func InitHttpProxy() {
	loadPage("asset/dist", "")
}

func loadPage(path string, urlPath string) {
	var files []os.FileInfo
	if path == "" {
		path = "./"
	}
	//获取当前文件夹下所有文件包括文件夹
	files, _ = utils.GetDirInProject(path)
	for _, f := range files {
		if f.IsDir() { //如果当前是一个文件夹
			loadPage(appendPath(path, f.Name()), appendUrlPath(urlPath, f.Name()))
		} else {
			//绑定文件名
			if path == "./" {
				loadHandle(path+f.Name(), urlPath+f.Name())
			} else {
				if f.Name() == "index.html" {
					loadHandle(path+"/"+f.Name(), urlPath+"/")
				}
				loadHandle(path+"/"+f.Name(), urlPath+"/"+f.Name())
			}
		}
	}
}

//文件名和文件绑定
func loadHandle(path string, urlPath string) {
	bytes, _ := utils.GetFileInProject(path)
	http.HandleFunc(urlPath, func(writer http.ResponseWriter, request *http.Request) {
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

//拼接路径
func appendPath(prefixPath string, suffixPath string) string {
	if prefixPath == "./" {
		return prefixPath + suffixPath
	}
	return prefixPath + "/" + suffixPath
}

//拼接url->端口号后部分
func appendUrlPath(prefixUrlPath string, suffixUrlPath string) string {
	if prefixUrlPath == "/" {
		return prefixUrlPath + suffixUrlPath
	}
	return prefixUrlPath + "/" + suffixUrlPath
}
