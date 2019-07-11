#### Supported Databases
* mysql

#### Required File And Folder
* mime.ini
* config.json
* dist
* tpl

#### Install Dependence

1. dep init
2. dep ensure

#### Run
* go run main.go
* [View](http://127.0.0.1:8888) `Default Port 8888`

#### Build On MacOs
* For Windows-64bit  
```
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o main.exe main.go
```
* For Linux-64bit
```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go
```
#### Build On Windows
* For MacOs
```
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build -o main main.go
```
* For Linux-64bit
```
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o main main.go
```
#### Build On Linux
* For MacOs
```
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o main main.go
```
* For Windows-64bit
```
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o main.exe main.go
```
