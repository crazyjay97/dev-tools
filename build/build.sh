#!/usr/bin/env bash

projectPath=$(cd `dirname $0`/..; pwd)

cd ${projectPath}

echo Current Path `pwd`

#app generate path
appPath="${projectPath}/bin/"


#create folder
    echo "create folder"
if [[ ! -d "$appPath" ]]; then
   mkdir bin
fi

echo "asset process"
go-bindata -o=./asset/asset.go -pkg=asset ./asset/... ./configs/...

#build for mac
echo "build for mac"
go build -o ${appPath}/code-generator ${projectPath}/cmd/main.go

#build for windows
echo "build for windows"
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${appPath}/code-generator.exe ${projectPath}/cmd/main.go

#build for linux
echo "build for linux"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${appPath}/code-generator-linux ${projectPath}/cmd/main.go



echo "package success"

