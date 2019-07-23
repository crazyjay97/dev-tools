#!/usr/bin/env bash

#app generate path
appPath="./app/"

#create folder
echo "create folder"
if [[ ! -d "$appPath" ]]; then
   mkdir app
fi

#build for mac
echo "build for mac"
go build -o ${appPath}/main main.go

#build for windows
echo "build for windows"
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${appPath}/main.exe main.go

#build for linux
echo "build for linux"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${appPath}/main.out main.go



#copy dependency
echo "copy dependency"
cp -r dist tpl types.ini config.json ${appPath}

#zip
echo "zip files"
#zip -r code-generator-app.zip ${appPath}
tar czvf code-generator-app.tar.gz ${appPath}

echo "package success"

