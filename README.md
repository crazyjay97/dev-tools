>[中文](./README_ZH.md)

#### Supported Databases
* mysql

#### Environment (_Ignore If Startup Without Problems_)
* go v1.11
* node v10.16.0
* npm v6.9.0


#### Run
* go run main.go
* [View](http://127.0.0.1:9999) `Default Port 9999`


#### Before Build (_Static File To Binary_)
```bash
cd web/
npm i #cnpm i
npm run build
cd ..
go-bindata -o=./asset/asset.go -pkg=asset ./asset/... ./configs/...
```
###### If **go-bindata** Command Not Found，You Need Run ```go get -u github.com/jteeuwen/go-bindata/...```  



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
