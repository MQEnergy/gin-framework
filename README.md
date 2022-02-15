## 以gin框架为基础，封装一套适用于面向api编程的快速开发框架
## 一、目录结构
```
├── app                         # 模块存放目录
│   ├── controller              # 控制器
│   └── service                 # 服务层
├── bootstrap                   # 初始化程序加载服务
├── config                      # 解析配置文件
├── entities                    # 存放表对应的实体
├── global                      # 一些全局变量和全局方法
├── main.go                     # 主进程启动文件
├── middleware                  # 中间件
├── migrate.go                  # 数据迁移文件
├── migrations                  # 数据迁移的sql文件目录
├── models                      # 对应数据库的模型
├── pkg                         # 自定义的常用服务，JWT,助手函数等
│   ├── auth                    # jwt
│   ├── lib                     # 日志服务，数据库服务，redis服务
│   ├── response                # http请求返回的状态和格式化
│   ├── util                    # 助手函数
│   └── validator               # 验证器
├── rbac_model.conf             # rbac配置文件
├── router                      # 路由配置
├── runtime                     # 运行时文件 如日志
```

## 二、启动服务
### 1、安装依赖
```shell script
go mod tidy 
```
### 2、执行主进程
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
    code: 0,
    message: "请求成功",
    data: "Pong!"
}
```
## 工具

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

## 参考 
### 初始化一个接口项目需要安装的依赖包
### 初始化go.mod
```shell script
go mod init lyky/gin-framework-template
go mod tidy
```

### 安装gin框架
```shell script
go get -u github.com/gin-gonic/gin
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

# 第一种方式执行迁移
# 执行迁移操作：
migrate -database 'mysql://root:123456@tcp(127.0.0.1:3306)/gin_framework' -path ./migrations up
# 执行回滚操作：
migrate -database 'mysql://root:123456@tcp(127.0.0.1:3306)/gin_framework' -path ./migrations down

# 第二种方式执行迁移
# env为 dev, test, prod与config.*.yaml文件保持一致
# n为：执行的迁移文件数量（回滚的文件数量）
# 执行迁移操作：
go run main.go migrate {n} {env}

# 执行回滚操作：
go run main.go migrate -{n} {env}
```
