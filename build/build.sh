#!/usr/bin/env bash

#app generate path
appPath="./app/"

#create folder
    echo "create folder"
if [[ ! -d "$appPath" ]]; then
   mkdir app
fi

echo "asset process"
go-bindata -o=./asset/asset.go -pkg=asset asset/dist/... asset/tpl/... configs/types.ini configs/config.json

#build for mac
echo "build for mac"
go build -o ${appPath}/code-generator.app main.go

#build for windows
echo "build for windows"
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${appPath}/code-generator.exe main.go

#build for linux
echo "build for linux"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${appPath}/code-generator.out main.go



echo "package success"

