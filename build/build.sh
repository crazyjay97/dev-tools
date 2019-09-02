#!/usr/bin/env bash

projectPath=$(cd `dirname $0`/..; pwd)

#app generate path
appPath="${projectPath}/bin/"

#create folder
    echo "create folder"
if [[ ! -d "$appPath" ]]; then
   mkdir bin
fi

echo "asset process"
go-bindata -o=${projectPath}/asset/asset.go -pkg=asset ${projectPath}/asset/dist/... ${projectPath}/asset/tpl/... ${projectPath}/configs/types.ini ${projectPath}/configs/config.json

#build for mac
echo "build for mac"
go build -o ${appPath}/code-generator.app ${projectPath}/cmd/main.go

#build for windows
echo "build for windows"
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${appPath}/code-generator.exe ${projectPath}/cmd/main.go

#build for linux
echo "build for linux"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${appPath}/code-generator.out ${projectPath}/cmd/main.go



echo "package success"

