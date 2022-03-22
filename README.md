## :zap::rocket: 以gin框架为基础，封装一套基于go1.18+的适用于面向api编程的快速开发框架

本项目积极拥抱Go 1.18+版本，建议用户升级到v1.18及以上版本

```
注意：因本框架使用了一个基于Go1.18+泛型的Lodash风格的Go库，所以需要go版本升级到v1.18，如果Go版本 < v1.18
可自行修改pagination.go文件的Pagination方法中此段代码，自行比较：
difference, _ := lo.Difference[string](fields, pb.Fields)
如非特殊情况，建议将版本升级到Go v1.18及以上，后续迭代将完全支持Go 1.18+
```

## 一、目录结构

```
├── app                         # 模块存放目录
│   ├── controller              # 控制器
│   └── service                 # 服务层
├── bootstrap                   # 初始化程序加载服务
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
├── rbac_model.conf             # rbac配置文件
├── router                      # 路由配置
├── runtime                     # 运行时文件 如日志
```

## 二、启动服务

```
注意启动前需要将 mysql服务和redis服务开启，并配置config.dev.yaml文件(默认读取dev环境)中的mysql和redis配置
```

### 1、安装依赖

```shell script
go mod tidy 
```

### 2、服务启动

```shell script
go run main.go 
# 查看 main.go的参数
go run main.go --help
```

### 3、访问如下表示成功启动

```shell script
# 请求：
http://127.0.0.1:9527/ping

# 返回：
{
    "status": 200,
    "errcode": 0,
    "requestid": "9ac7f4f2-1271-4f87-8df7-599a478af9cb",
    "message": "Pong!",
    "data": ""
}
```

### 4、安装热更新

```shell script
go install github.com/cosmtrek/air@latest
```

命令行敲入：air 即可执行热更新 代码编辑即更新

### 5、打包上线

```shell script
go build main.go
# 执行可查看命令
./main help 

# 安装在GOPATH的bin目录中
go install

# 执行可查看命令
mqenergy-go help
```

## 三、组件使用

### 1、基于gorm的查询分页构造器

单表查询使用方法如下：

```shell script
import 	"mqenergy-go/pkg/paginator"
var memberList = make([]models.GinAdmin, 0)
paginator, err := paginator.Builder.
    WithDB(global.DB).
    WithModel(models.GinAdmin{}).
    WithFields([]string{"password", "salt", "updated_at", "_omit"}).
    WithCondition("id = ?", 1).
    Pagination(memberList, 1, 10)
return paginator, err
```

连表查询使用方法如下：

```
可查看案例 获取用户列表：
app/controller/backend/user.go
app/service/backend/user.go
router/routes/common.go

访问地址：
http://127.0.0.1:9527/user/index?page=1
可查看效果
```

返回数据格式如下：

```json
{
  "status": 200,
  "errcode": 0,
  "requestid": "9ac7f4f2-1271-4f87-8df7-599a478af9cb",
  "message": "请求成功",
  "data": {
    "list": [],
    "current_page": 2,
    "count": 13,
    "last_page": 2,
    "per_page": 10
  }
}
```

#### 1）`必须在链式操作中` WithDB(db *gorm.DB) *PageBuilder db连接方法

```
传入全局global.DB 
```

#### 2）`必须在链式操作中` WithModel(db *gorm.DB) *PageBuilder model连接方法

```
传入查询主表model  例如：models.GinAdmin 参数不能传结构体取地址
```

#### 3）`非必须在链式操作中` WithFields(fields []string) *PageBuilder 查询或过滤字段方法

```
最后一个参数默认为_select（可不传），如传_omit为过滤前面传输的字段。
注意： _select / _omit 必须在最后
      WithModel 参数不能传结构体取地址 例如：&models.GinAdmin 必须 models.GinAdmin
      不然 _omit 参数失效
如：
WithFields([]string{"created_at", "updated_at", "_omit"}) // 表示过滤前面字段
WithFields([]string{"created_at", "updated_at", "_select"}) // 表示查询前面的字段
```

#### 4）`非必须在链式操作中` WithCondition(query interface{}, args interface{}) *PageBuilder 数据查询条件方法

```
如上所示 传入查询条件 支持gorm中where条件中的一些查询方式（非struct方式） query, args参数参照gorm where条件传入方式
```

#### 5）`非必须在链式操作中` WithJoins(joinType string, joinFields OnJoins) *PageBuilder 数据查询条件方法

```
joinType：join类型 可传入：left,right,inner
joinFields结构体： LeftTableField：如：主表.ID  RightTableField：如：关联表.主表ID
具体参考 上述连表查询使用案例
```

#### 6）`必须在链式操作中最后一环` Pagination(list interface{}, currentPage, pageSize int) (Page, error) 分页返回方法

```
list 传入数据表的struct model，currentPage 为当前页码，pageSize为每页查询数量
```

#### 7）获取当前页码

```
paginator.CurrentPage
```

#### 8）获取分页列表

```
paginator.List
```

#### 9）获取数据总数

```
paginator.Count
```

#### 10）获取最后一页页码

```
paginator.LastPage
```

#### 10）获取每页数据条数

```
paginator.PerPage
```

### 2、基于gin上传组件

```
UploadFile(path string, r *gin.Context) (*FileHeader, error)
```

默认存储在项目中upload目录，如果没有会自动创建 path：upload目录模块目录 如：user 则目录是：upload/user/{yyyy-mm-dd}/... 上传附件案例参照：

```
目录：
app/controller/backend/attachment.go
router/routes/common.go
```

## 四、工具

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

### 1、执行migrate

```shell script
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
# env： dev, test, prod与config.*.yaml文件保持一致 默认是dev
# n：执行的迁移文件数量（回滚的文件数量）例如：1，2，3...

# 执行迁移操作：
go run main.go migrate {n} {env}

# 执行回滚操作：
go run main.go migrate -{n} {env}
```

### 2、自动生成model

```shell script
# 参数一：all：生成所有 或者写入数据表名生成单个
# env：dev, test, prod与config.*.yaml文件保持一致 默认是dev

# 执行生成所有model
go run main.go model all {env}
#
go run main.go model {数据表名} {env}
# 例如：go run model.go gin_admin
```

### 3、自动生成controller

```shell script
# 参数一：all：生成所有 或者写入数据表名生成单个
# env：dev, test, prod与config.*.yaml文件保持一致 默认是dev

go run main.go controller {controller名称} {module名称}
# module名称是app/controller目录下的模块名称
# 例如：go run main.go controller admin backend
```

### 4、自动生成service

```shell script
# env：dev, test, prod与config.*.yaml文件保持一致 默认是dev

go run main.go service {service名称} {module名称}
# module名称是app/controller目录下的模块名称
# 例如：go run main.go service admin backend
```

### 5、创建后台管理员账号（基于gin_admin表的，可自行修改代码基于其他表）

```
go run main.go account {账号名称} {密码}  
```

## 五、参考

### 初始化一个接口项目需要安装的依赖包（主要）

### 初始化go.mod

```shell script
go mod init mqenergy-go/gin-framework
go mod tidy
```

### 安装gin框架

```shell script
go get -u github.com/gin-gonic/gin
```

### 安装model自动生成包

```shell script
go get -u github.com/MQEnergy/gorm-model
```

### 安装gorm

```shell script
go get -u gorm.io/gorm
// 如果下载不了 十之八九是因为 GOSUMDB的原因 
export GOSUMDB=
// GOSUMDB置空就行
```

### 安装命令行工具

```shell script
go get -u github.com/urfave/cli/v2
```

### 安装log日志

```shell script
go get -u github.com/sirupsen/logrus
go get -u github.com/lestrrat-go/file-rotatelogs
```

### 安装redis

```shell script
go get -u github.com/go-redis/redis/v8
go get -u github.com/go-redsync/redsync/v4
```

### 安装jwt

```shell script
go get -u github.com/dgrijalva/jwt-go
```

### 安装cors跨域

```shell script
go get -u github.com/gin-contrib/cors
```

### 安装casbin

```shell script
go get -u github.com/casbin/casbin/v2
go get -u github.com/casbin/gorm-adapter/v3
```

### 安装snowflake

```shell script
go get -u github.com/bwmarrin/snowflake
```

### 安装golang-migrate迁移组件

```shell script
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

```shell script
go install github.com/cosmtrek/air@latest
```

### 基于Go 1.18+泛型的Lodash风格的Go库

```shell script
go get -u github.com/samber/lo
```
