### `项目简述`
这是一个简历管理系统后端使用golang \
有兴趣可以看看这个项目和GO项目[react项目](https://github.com/focusdroid/react-ts-resume) \ [Golang项目链接](https://github.com/focusdroid/ResumeManagement)
项目使用前后端分离，前端使用react + ts，后端使用golang开发，其中后端功能如下 \ 
添加简历，标记重点简历，每日待办（添加、编辑），简历编辑，提交简历信息，在线预览简历，删除简历，接口查询，逻辑删除，接口返回，

### `go version`
- golang版本使用1.19.5,目前最新的版本是1.2

### `本地部署须知`
 1. 本地部署需要安装 [go1.19.5版本](https://go.dev/dl/) 选择指定版本进行本地安装，使用go env查看是否正确安装
 2. 安装本地或远程的mysql数据库，并创建名为resume的数据库 charset=utf8
 3. 本地安装redis并启动redis(window电脑有的版本需要手动启动redis-serve)一般是默认启动的
 4. 以上安装成功之后运行 `go run main.go` 建议安装[fresh](https://github.com/gravityblast/fresh)
 5. 使用 `swag init` 本地使用swagger查看接口文档 `http://127.0.0.1:8080/swagger/index.html#/`


### `项目目录`
- 根目录存放项目启动文件(main.go)和文件说明
- controllers目录下存放各个路由对应的具体功能方法(结构体函数、接口查询、数据库操作、接口数据返回)
- docs是配置swagger之后存放的目录，这个是有 [go-swagger](https://github.com/go-swagger/go-swagger)自动生成的文件，如果每次修改之后需要重新使用`swag init`重新初始化获取最新的文档
- file目录是使用gin测试上传的目录，目前其中图片没有使用，但不确定后期是否使用
- helper目录是对各个功能函数进行具体的抽离，是的接口功能更加清晰
- middleware(中间件)目录是针对项目中接口进行token校验或者其他公共处理的
- models初始化数据库表和数据库字段（初始化之前需要手动创建数据库，并使用该数据库地址和数据库名称，数据库密码进行链接）
- routers是对项目中所有的路由进行管理，庞大的项目需要细化各个功能，路由配置就很好管理各个功能模块
- test目录是对项目中需要使用的到功能进行测试，需要使用go中的testing包
- tmp目录是使用[fresh](https://github.com/gravityblast/fresh)之后自动生成的运行目录。ps: go get github.com/pilu/fresh@latest
- go.mod是go项目在升级之后go包版本管理目录，如果切换golang版本之后需要删除之后重新init之后tidy

### `版本分支说明`
```text
featrue20230225 分支添加黑名单功能
```

### `golang项目使用到的包地址`
```text
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get github.com/gin-contrib/sessions
github.com/jordan-wright/email
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
go get -u github.com/go-redis/redis/v9
go get -u gorm.io/driver/mysql
go get -u github.com/golang-jwt/jwt/v4
```
[go get -u github.com/gin-gonic/gin](https://github.com/gin-gonic/gin) gin \
[go get -u github.com/swaggo/swag/cmd/swag](github.com/swaggo/swag/cmd/swag) swag \
[go get -u gorm.io/gorm](https://gorm.io/) gorm \
[go get github.com/gin-contrib/sessions](https://github.com/gin-contrib/sessions#redis) redis \
[go get github.com/jordan-wright/email](github.com/jordan-wright/email) email \
[go get -u github.com/go-redis/redis/v9](github.com/go-redis/redis/v9) redis \
[go get -u gorm.io/driver/mysql](gorm.io/driver/mysql) mysql \
[go get -u github.com/golang-jwt/jwt/v4](github.com/golang-jwt/jwt/v4) mysql \