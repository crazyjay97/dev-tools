#### 支持的数据库
* mysql

#### 环境 (_启动无错误时可以忽略_)
* go v1.11
* node v10.16.0
* npm v6.9.0


#### 运行
* go run cmd/main.go
* [页面](http://127.0.0.1:9999) `默认端口 9999`


#### 构建之前 (_静态文件转二进制_)
```bash
cd web/
npm i #cnpm i
npm run build
cd ..
go-bindata -o=./asset/asset.go -pkg=asset ./asset/... ./configs/...
```
###### 如果 **go-bindata** 命令不存在，请运行```go get -u github.com/jteeuwen/go-bindata/...```  



#### 在MacOs上构建
* Windows-64bit版本
```
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o main.exe main.go
```
* Linux-64bit版本
```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go
```

#### 在Windows上构建
* MacOs版本
```
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build -o main main.go
```
* Linux-64bit版本
```
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o main main.go
```

#### 在Linux上构建
* MacOs版本
```
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o main main.go
```
* Windows-64bit版本
```
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o main.exe main.go
```


#### 配置

* configs/config.json 
    * port 端口号
    * db 数据库相关配置
        * dialect 数据库类型/方言
        * url 数据库链接
        * maxIdleCons 数据库最大空闲连接数
        * maxOpenCon 最大数据库链接
        * printLog 是否开启数据库打印
    * tpl 模版相关配置
        * name 输出的文件名
        * root 输出的文件夹名称
        * filename 模版名称
        * needModule 是否需要根据页面划分模块
        * appendFileName 是否需要在输出文件名前追加页面名
        * appendClassName 是否需要在输出文件名前追类名
        * customModule 二级模块名称
        
#### 目录
   
* asset 静态文件
    * dist 编译后的前端内容
    * tpl 模版文件
    * asset.go 二进制静态内容
* bin 可执行文件
* build 编译相关脚本
* cmd 启动文件
* configs 配置文件
    * configs.json 项目配置文件
    * types.ini 类型配置文件
* internal 程序相关代码
    * app 程序逻辑代码
    * common 通用代码
* test 测试相关
* web 前端页面
    
        
