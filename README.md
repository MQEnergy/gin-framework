# :zap::rocket: 以gin框架为基础，封装一套基于go1.18+的适用于面向api编程的快速开发框架

本项目积极拥抱Go 1.18+版本，强烈建议用户升级到v1.18及以上版本，全自动化生成Model Service Controller架子，加快业务开发。

[![GitHub license](https://img.shields.io/github/license/MQEnergy/gin-framework)](https://github.com/MQEnergy/gin-framework/blob/main/LICENSE)
[![GitHub stars](https://img.shields.io/github/stars/MQEnergy/gin-framework)](https://github.com/MQEnergy/gin-framework/stargazers)

# 一、目录结构

```
├── app                         # 模块存放目录
│   ├── controller              # 控制器
│   └── service                 # 服务层
├── bootstrap                   # 初始化程序加载服务
│   ├── rbac_model.conf         # rbac配置文件
├── command                     # command命令
├── config                      # 解析配置文件
├── entities                    # 存放表对应的实体 / 请求参数的结构体
├── global                      # 一些全局变量和全局方法
├── main.go                     # 主进程启动文件
├── middleware                  # 中间件
├── migrations                  # 数据迁移的sql文件目录
├── models                      # 对应数据库的模型
├── pkg                         # 自定义的常用服务，JWT,助手函数等
│   ├── auth                    # jwt
│   ├── lib                     # 日志服务，数据库服务，redis服务
│   ├── paginator               # 分页器
│   ├── response                # http请求返回的状态和格式化
│   ├── util                    # 助手函数
│   └── validator               # 验证器
├── config.dev.yaml             # 开发环境配置文件
├── config.test.yaml            # 测试环境配置文件
├── config.prod.yaml            # 正式环境配置文件
├── router                      # 路由配置
├── runtime                     # 运行时文件 如日志
```
#### 目前已集成和实现：
- [x] 支持 [jwt](https://github.com/dgrijalva/jwt-go) Authorization token验证组件
- [x] 支持 [cors](https://github.com/gin-contrib/cors) 接口跨域组件
- [x] 支持 [gorm](https://gorm.io) 数据库操作组件
- [x] 支持 [gorm-model](https://github.com/MQEnergy/gorm-model) 自主实现的基于gorm生成的映射数据表的model结构体
- [x] 支持 [logrus](https://github.com/sirupsen/logrus) 日志收集组件
- [x] 支持 [go-redis](https://github.com/go-redis/redis) redis连接组件
- [x] 支持 [migrate](https://github.com/golang-migrate/migrate) 数据库迁移组件
- [x] 支持 [controller、service](https://github.com/MQEnergy/gin-framework/tree/main/command) 命令行方式生成代码工具
- [x] 支持 [go-websocket](https://github.com/MQEnergy/go-websocket) 基于 gorilla/websocket 实现的即时通讯组件（单个客户端，多个客户端，群组，广播推送等）
- [x] 支持 [go-rabbitmq](https://github.com/MQEnergy/go-rabbitmq) 消息队列组件 基于rabbitmq官方 [amqp](https://github.com/streadway/amqp) 组件封装实现的消费者和生产者
- [x] 支持 [casbin](https://github.com/casbin/casbin) rbac权限 集成于中间件中 [casbin_auth.go](https://github.com/MQEnergy/gin-framework/blob/main/middleware/casbin_auth.go)
- [x] 支持 [requestId](https://github.com/gin-contrib/requestid) 中间件 实现了方便链路追踪日志记录中间件 [requestid_auth.go](https://github.com/MQEnergy/gin-framework/blob/main/middleware/requestid_auth.go)
- [x] 支持 [viper](https://github.com/spf13/viper) yaml、json、toml等配置文件解析组件
- [x] 支持 [validator](https://github.com/go-playground/validator) 数据字段验证器组件，同时支持中文
- [x] 支持 [snowflake](https://github.com/bwmarrin/snowflake) 生成雪花算法全局唯一ID
- [x] 实现 ip白名单配置 集成于中间件中 [ip_auth.go](https://github.com/MQEnergy/gin-framework/blob/main/middleware/ip_auth.go)
- [x] 实现 [ticker](https://github.com/MQEnergy/gin-framework/blob/main/pkg/util/ticker.go) 定时器组件 
- [x] 实现 基于gorm的 [pagination](https://github.com/MQEnergy/gin-framework/blob/main/pkg/paginator/pagination.go) 分页构造器组件
- [x] 实现 [code](https://github.com/MQEnergy/gin-framework/tree/main/pkg/response/code.go) 统一定义的返回码，[exception](https://github.com/MQEnergy/gin-framework/tree/main/pkg/response/exception.go) 统一错误返回处理组件
- [x] 支持 [go-bindata](https://github.com/go-bindata/go-bindata) 将数据文件转换为go代码，便于因编译成二进制找不到yaml文件配置的问题

#### 下一步计划：
- [ ] 支持 定时任务 cron
- [ ] 支持 [pprof](https://github.com/gin-contrib/pprof) 性能剖析组件
- [ ] 支持 trace 项目内部链路追踪
- [ ] 支持 [rate](https://pkg.go.dev/golang.org/x/time/rate) 接口限流组件
- [ ] 支持 [grpc](https://github.com/grpc/grpc-go) rpc组件

# 二、启动服务

```
注意启动前需要将 mysql服务和redis服务开启，并配置config.dev.yaml文件(默认读取dev环境)中的mysql和redis配置
```

## 1、安装依赖和初始化

```bash
go mod tidy 

# 执行make（查看Makefile文件）
make bindata
```
### 注意：
```
当修改yaml配置文件时候，需要执行 `make bindata`，重新将数据文件转换为go代码，本地开发和部署上线才生效。
```

## 2、服务启动

```bash
go run main.go 
# 查看 main.go的参数
go run main.go --help
```

## 3、访问如下表示成功启动
请求：http://127.0.0.1:9527/ping
```json
{
    "status": 200,
    "errcode": 0,
    "requestid": "9ac7f4f2-1271-4f87-8df7-599a478af9cb",
    "message": "Pong!",
    "data": ""
}
```

## 4、安装热更新

```bash
go install github.com/cosmtrek/air@latest
```

命令行敲入：air 即可执行热更新 代码编辑即更新

## 5、部署casbin权限（重要！（按以下步骤执行））
```
此步骤针对于backend接口进行权限访问
```
### 1）执行migrate
```
go run main.go migrate -s all
```
### 2）请求 `/routes` 接口
```
此接口会创建一个基于casbin的超级管理员权限
```

## 6、打包上线

```bash
# 查看make命令行
make help

# 基础打包，生成可执行文件
make build

# 打包windows
make windows

# 打包darwin
make darwin

# 打包linux
make linux
```
在releases中查看打包的文件

# 三、组件使用

## 1、基于gorm的查询分页构造器
引用包

```go
import "mqenergy-go/pkg/paginator"
```

### 一、基础用法
#### 1）单表分页基础用法：

```go
var memberList = make([]models.GinAdmin, 0)
paginator, err := paginator.NewBuilder().
    WithDB(global.DB).
    WithModel(models.GinAdmin{}).
    WithField([]string{"password", "salt", "updated_at", "_omit"}).
    WithCondition("id = ?", 1).
    Pagination(memberList, 1, 10)
return paginator, err
```

#### 2）连表joins查询用法：
定义接收struct
```go
type BaseUser models.GinUser
type GinUserInfo models.GinUserInfo

// UserList 获取关联列表
type UserList struct {
	BaseUser
	GinUserInfo `gorm:"foreignKey:user_id" json:"user_info"`
}
```

用法一：
```go
var userList = make([]user.UserList, 0)
pagination, err := paginator.NewBuilder().
    WithDB(global.DB).
    WithModel(models.GinUser{}).
    WithFields(models.GinUser{}, models.GinUserTbName, []string{"password", "salt", "_omit"}).
    WithFields(models.GinUserInfo{}, models.GinUserInfoTbName, []string{"id", "user_id", "role_ids"}).
    WithJoins("left", []paginator.OnJoins{{
        LeftTableField:  paginator.JoinTableField{Table: models.GinUserTbName, Field: "id"},
        RightTableField: paginator.JoinTableField{Table: models.GinUserInfoTbName, Field: "user_id"},
    }}).
    Pagination(&userList, requestParams.Page, global.Cfg.Server.DefaultPageSize)
return pagination, err
```

用法二：
```go
var userList = make([]user.UserList, 0)
multiFields := []paginator.SelectTableField{
    {Model: models.GinUser{}, Table: models.GinUserTbName, Field: []string{"password", "salt", "_omit"}},
    {Model: models.GinUserInfo{}, Table: models.GinUserInfoTbName, Field: []string{"user_id", "role_ids"}},
}	
pagination, err := paginator.NewBuilder().
    WithDB(global.DB).
    WithModel(models.GinUser{}).
    WithMultiFields(multiFields).
    WithJoins("left", []paginator.OnJoins{{
        LeftTableField:  paginator.JoinTableField{Table: models.GinUserTbName, Field: "id"},
        RightTableField: paginator.JoinTableField{Table: models.GinUserInfoTbName, Field: "user_id"},
    }}).
    Pagination(&userList, requestParams.Page, global.Cfg.Server.DefaultPageSize)
return pagination, err
```

#### 3）预加载preload查询用法（强烈建议用法）：

```
注意：
与joins查询方式定义的struct有些许差别，preload方式定义struct名称必须与model当前表的struct名称一致，
且关联表的struct名称不能跟model对于的struct名称一样 例如：定义的`UserInfo` 写法如下
```    

定义接收struct
```go
type BaseUser models.GinUser
type GinUserInfo models.GinUserInfo

type GinUser struct {
	BaseUser
	UserInfo GinUserInfo `gorm:"foreignKey:user_id" json:"user_info"`
}
```

用法如下：
```go
var userList = make([]user.GinUser, 0)
pagination, err := paginator.NewBuilder().
    WithDB(global.DB).
    WithModel(models.GinUser{}).
    WithPreload("UserInfo").
    Pagination(&userList, requestParams.Page, global.Cfg.Server.DefaultPageSize)
return pagination, err
```
```
此写法不建议使用WithFields、WithField查询字段，建议直接定义接收struct规定的查询字段即可
```

访问地址：http://127.0.0.1:9527/user/index?page=1 返回数据格式如下：

```json
{
  "status": 200,
  "errcode": 0,
  "requestid": "9ac7f4f2-1271-4f87-8df7-599a478af9cb",
  "message": "请求成功",
  "data": {
    "list": [],
    "current_page": 1,
    "total": 2,
    "last_page": 1,
    "per_page": 10
  }
}
```
#### 4）案例查看：
1）用法如下 获取用户列表：
```
entities/user/gin_user.go
app/controller/backend/user.go
app/service/backend/user.go
router/routes/common.go
```

### 二、具体方法
#### 1）`必须在链式操作中` db连接方法
```go
WithDB(db *gorm.DB) *PageBuilder
```
传入全局global.DB

#### 2）`必须在链式操作中` model连接方法
```go
WithModel(model interface{}) *PageBuilder
```
传入查询主表model  例如：models.GinAdmin 参数不能传结构体取地址方式 如：&models.GinAdmin

#### 3）`非必须在链式操作中` 单表查询或过滤字段方法
```go
WithField(fields []string) *PageBuilder 
```

fields 最后一个参数默认为_select（可不传），如传_omit为过滤前面传输的字段。

注意：
- _select / _omit 必须在最后
- WithModel 参数不能传结构体取地址 例如：&models.GinAdmin 必须 models.GinAdmin 不然 _omit 参数失效
- 此注意事项适用于 `WithFields`方法、`WithMultiFields`方法

用法如下：
```go
// 表示过滤前面字段
WithField([]string{"created_at", "updated_at", "_omit"})

// 表示查询前面的字段
WithField([]string{"created_at", "updated_at", "_select"})
WithField([]string{"created_at", "updated_at"})
```

#### 4）`非必须在链式操作中` 多表查询或过滤字段方法（preload模式下 关联表查询有问题，preload关联查询不建议使用此方法）
```go
WithFields(model interface{}, table string, fields []string) *PageBuilder
```
fields 最后一个参数默认为_select（可不传），如传_omit为过滤前面传输的字段。

用法如下：
```go
// 表示过滤前面字段
WithFields(models.GinUser{}, models.GinUserTbName, []string{"password", "salt", "_omit"})

// 表示查询前面的字段
WithFields(models.GinUserInfo{}, models.GinUserInfoTbName, []string{"id", "user_id", "role_ids", "_select"})
WithFields(models.GinUserInfo{}, models.GinUserInfoTbName, []string{"id", "user_id", "role_ids"})
```

#### 5）`非必须在链式操作中` 多表多字段查询（可替代WithFields方法）
```go
WithMultiFields(fields []SelectTableField) *PageBuilder
```

用法如下：
```go
WithMultiFields([]paginator.SelectTableField{
    {Model: models.GinUser{}, Table: models.GinUserTbName, Field: []string{"password", "salt", "_omit"}},
    {Model: models.GinUserInfo{}, Table: models.GinUserInfoTbName, Field: []string{"id", "user_id", "role_ids"}},
})
```

#### 6）`非必须在链式操作中` 多表关联查询主动预加载（暂不支持条件）
    
```go
 WithPreloads(querys []string) *PageBuilder 
```

用法如下：
```go
WithPreloads([]string{"UserInfo", "UserRecord"})
```

#### 7）`非必须在链式操作中` 关联查询主动预加载（可传条件，条件参考gorm）

```go
WithPreload(query string, args ...interface{}) *PageBuilder
```

用法如下：
```go
WithPreload("UserInfo", "user_id = ?", "1")
```

#### 8）`非必须在链式操作中` 数据查询条件方法

```go
WithCondition(query interface{}, args ...interface{}) *PageBuilder
```

传入查询条件 支持gorm中where条件中的查询方式（非struct方式） query, args参数参照gorm的where条件传入方式

#### 9）`非必须在链式操作中` 数据查询条件方法

```go
WithJoins(joinType string, joinFields []OnJoins) *PageBuilder
```

joinType：join类型 可传入：left、right、inner，joinFields结构体： LeftTableField：如：主表.ID  RightTableField：如：关联表.主表ID

用法如下：
```go
WithJoins("left", []paginator.OnJoins{{
    LeftTableField:  paginator.JoinTableField{Table: models.GinUserTbName, Field: "id"},
    RightTableField: paginator.JoinTableField{Table: models.GinUserInfoTbName, Field: "user_id"},
}})
```

#### 10）`必须在链式操作中最后一环` 分页返回方法

```go
Pagination(dst interface{}, currentPage, pageSize int) (Page, error)
```

dst 传入接收数据的struct结构体 注意：必须是应用方式传递 如：&userList，
model，currentPage 为当前页码，pageSize为每页查询数量

#### 11）`非必须在链式操作中` 对接原生查询方式

```go
 NewDB() *gorm.DB
```

用此方法之后的链式操作下pagination里面的方法均不可用，后面跟gorm原生方法即可

用法如下：
```go
NewDB().Where("id = ?", id).First(&userList)
```

#### 12）获取当前页码

```go
paginator.CurrentPage
```

#### 13）获取分页列表

```go
paginator.List
```

#### 14）获取数据总数

```go
paginator.Total
```

#### 15）获取最后一页页码

```go
paginator.LastPage
```

#### 16）获取每页数据条数

```go
paginator.PerPage
```

## 2、基于gin上传组件

```go
UploadFile(path string, r *gin.Context) (*FileHeader, error)
```

默认存储在项目中upload目录，如果没有会自动创建 path：upload目录模块目录 如：user 则目录是：upload/user/{yyyy-mm-dd}/... 

用法如下：
```
app/controller/backend/attachment.go
pkg/util/upload.go
router/routes/common.go
```
## 3、rabbitmq组件使用
配置yaml配置文件中的amqp参数
### 1）启动消费者
测试案例
```shell
go run command/test/consumer.go
```
### 2）启动生产者
测试案例
```shell
go run command/test/producer.go
```

# 四、工具

运行 go run main.go --help 可查看到以下命令集

```
COMMANDS:
  migrate     Create migration command
  account     Create a new admin account
  model       Create a new model class
  controller  Create a new controller class
  service     Create a new service class
  help, h     Shows a list of commands or help for one command
```

## 1、执行migrate

```bash
# 安装migrate cli工具
curl -L https://github.com/golang-migrate/migrate/releases/download/$version/migrate.$platform-amd64.tar.gz | tar xvz

# MacOS安装
brew install golang-migrate

# Window 使用scoop安装 https://scoop.sh/
scoop install migrate

# 创建迁移文件语法例如:
migrate create -ext sql -dir migrations -seq create_users_table

# 第一种方式执行迁移
# 执行迁移操作：
migrate -database 'mysql://root:123456@tcp(127.0.0.1:3306)/gin_framework' -path ./migrations up
# 执行回滚操作：
migrate -database 'mysql://root:123456@tcp(127.0.0.1:3306)/gin_framework' -path ./migrations down

# 第二种方式执行迁移
# 查看help命令
go run main.go migrate --help

# 格式如下：
go run main.go migrate -s {step} -e {env}
# env： dev, test, prod与config.*.yaml文件保持一致 默认是dev
# step：执行的迁移文件数量（回滚的文件数量）例如：1，2，3... 如果执行所有传 all

# 执行所有迁移操作：
go run main.go migrate -s all

# 执行部分迁移操作：
# 如：go run main.go migrate -s 1

# 执行回滚操作：
# 如：go run main.go migrate -s -1
```

## 2、自动生成model

```bash
# 参数一：all：生成所有 或者写入数据表名生成单个
# env：dev, test, prod与config.*.yaml文件保持一致 默认是dev

# 执行生成所有model
go run main.go model all {env}
#
go run main.go model {数据表名} {env}
# 例如：go run main.go model gin_admin
```

## 3、自动生成controller

```bash
go run main.go controller {controller名称} {module名称}
# module名称是app/controller目录下的模块名称
# 例如：go run main.go controller admin backend
```

## 4、自动生成service

```bash
go run main.go service {service名称} {module名称}
# module名称是app/controller目录下的模块名称
# 例如：go run main.go service admin backend
```

## 5、创建后台管理员账号（基于gin_admin表的，可自行修改代码基于其他表）

```bash
go run main.go account {账号名称} {密码}  
```

# 五、参考
## 初始化一个接口项目需要安装的依赖包（主要）
### 初始化go.mod

```bash
go mod init mqenergy-go/gin-framework
go mod tidy
```

### 安装gin框架

```bash
go get -u github.com/gin-gonic/gin
```

### 安装model自动生成包

```bash
go get -u github.com/MQEnergy/gorm-model
```

### 安装gorm

```bash
go get -u gorm.io/gorm
# 如果下载不了 十之八九是因为 GOSUMDB的原因 
export GOSUMDB=
# GOSUMDB置空就行
```

### 安装命令行工具

```bash
go get -u github.com/urfave/cli/v2
```

### 安装log日志

```bash
go get -u github.com/sirupsen/logrus
go get -u github.com/lestrrat-go/file-rotatelogs
```

### 安装redis

```bash
go get -u github.com/go-redis/redis/v8
go get -u github.com/go-redsync/redsync/v4
```

### 安装jwt

```bash
go get -u github.com/dgrijalva/jwt-go
```

### 安装cors跨域

```bash
go get -u github.com/gin-contrib/cors
```

### 安装casbin

```bash
go get -u github.com/casbin/casbin/v2
go get -u github.com/casbin/gorm-adapter/v3
```

### 安装snowflake

```bash
go get -u github.com/bwmarrin/snowflake
```

### 安装golang-migrate迁移组件

```bash
go get -u github.com/golang-migrate/migrate/v4

# 安装migrate cli工具
curl -L https://github.com/golang-migrate/migrate/releases/download/$version/migrate.$platform-amd64.tar.gz | tar xvz

# MacOS安装
brew install golang-migrate

# Window 使用scoop安装 https://scoop.sh/
scoop install migrate

# 创建迁移文件语法例如:
migrate create -ext sql -dir migrations -seq create_users_table

# 执行迁移操作：
migrate -database 'mysql://root:123456@tcp(127.0.0.1:3306)/gin_framework' -path ./migrations up
# 执行回滚操作：
migrate -database 'mysql://root:123456@tcp(127.0.0.1:3306)/gin_framework' -path ./migrations down
```

### 安装热更新

```bash
go install github.com/cosmtrek/air@latest
```

### 基于Go 1.18+泛型的Lodash风格的Go库

```bash
go get -u github.com/samber/lo
```

### 配置文件解析库

```bash
go get -u github.com/spf13/viper
```
