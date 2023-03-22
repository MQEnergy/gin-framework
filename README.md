# :zap::rocket: Based on the gin framework, we have developed a user-friendly, simple, and fast development framework for API programming using go1.18+.

```
 ██████╗ ██╗███╗   ██╗      ███████╗██████╗  █████╗ ███╗   ███╗███████╗██╗    ██╗ ██████╗ ██████╗ ██╗  ██╗
██╔════╝ ██║████╗  ██║      ██╔════╝██╔══██╗██╔══██╗████╗ ████║██╔════╝██║    ██║██╔═══██╗██╔══██╗██║ ██╔╝
██║  ███╗██║██╔██╗ ██║█████╗█████╗  ██████╔╝███████║██╔████╔██║█████╗  ██║ █╗ ██║██║   ██║██████╔╝█████╔╝
██║   ██║██║██║╚██╗██║╚════╝██╔══╝  ██╔══██╗██╔══██║██║╚██╔╝██║██╔══╝  ██║███╗██║██║   ██║██╔══██╗██╔═██╗
╚██████╔╝██║██║ ╚████║      ██║     ██║  ██║██║  ██║██║ ╚═╝ ██║███████╗╚███╔███╔╝╚██████╔╝██║  ██║██║  ██╗
 ╚═════╝ ╚═╝╚═╝  ╚═══╝      ╚═╝     ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝     ╚═╝╚══════╝ ╚══╝╚══╝  ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝
```
[![GoDoc](https://godoc.org/github.com/MQEnergy/gin-framework/?status.svg)](https://pkg.go.dev/github.com/MQEnergy/gin-framework)
[![Go Report Card](https://goreportcard.com/badge/github.com/MQEnergy/gin-framework)](https://goreportcard.com/report/github.com/MQEnergy/gin-framework)
[![codebeat badge](https://codebeat.co/badges/1bf7dd49-1283-4ec9-b56e-a755e1e9c8dd)](https://codebeat.co/projects/github-com-mqenergy-gin-framework-main)
[![GitHub license](https://img.shields.io/github/license/MQEnergy/gin-framework)](https://github.com/MQEnergy/gin-framework/blob/main/LICENSE)
[![GitHub stars](https://img.shields.io/github/stars/MQEnergy/gin-framework)](https://github.com/MQEnergy/gin-framework/stargazers)

[中文文档](README.zh_CN.md)

# I、Directory Structure

```
├── Dockerfile
├── LICENSE
├── Makefile                        # makefile
├── README.md
├── app                             # Directory holding modules
│   ├── amqp                        # Message queue
│   ├── controller                  # Controller
│   └── service                     # Service layer
├── bootstrap                       # Initialization program loading service
├── cmd                             # Command commands
│   ├── admin.go                    # Generate admin backend account
│   ├── controller.go               # Generate controller
│   ├── migrate.go                  # Generate migrate database migration
│   ├── model.go                    # Generate model data model
│   ├── service.go                  # Generate service layer
├── config
│   ├── config.go                   # Map yaml configuration file to structure
│   ├── white_list.go               # Whitelist
│   └── yaml                        # yaml configuration file directory
├── global                          # Global variables and global methods
├── go.mod
├── go.sum
├── main.go
├── middleware                      # Middleware
├── migrations                      # Migration files
├── models                          # Models
├── pkg                             # Custom common services, JWT, helper functions, etc.
│   ├── auth                        # jwt
│   ├── lib                         # Log service, database service, redis service, etc.
│   ├── paginator                   # Paginator
│   ├── response                    # Http request returns status and formatting
│   ├── util                        # Helper function
│   └── validator                   # Parameter validator
├── router                          # Route configuration
├── runtime                         # Files produced at runtime, such as logs, etc.
├── types                           # All custom structures
```

#### Currently integrated and implemented:
- [x] Support for [jwt](https://github.com/dgrijalva/jwt-go) Authorization token validation component
- [x] Support for [cors](https://github.com/gin-contrib/cors) interface cross-domain component
- [x] Support for [gorm](https://gorm.io) database operation component
- [x] Support for [gorm-model](https://github.com/MQEnergy/gorm-model) self-implemented model structure based on gorm-generated mapping data tables
- [x] Support for [logrus](https://github.com/sirupsen/logrus) log collection component
- [x] Support for [go-redis](https://github.com/go-redis/redis) redis connection component
- [x] Support for [migrate](https://github.com/golang-migrate/migrate) database migration component
- [x] Support for [controller, service](https://github.com/MQEnergy/gin-framework/tree/main/command) command-line code generation tool
- [x] Support for [go-websocket](https://github.com/MQEnergy/go-websocket) real-time communication component based on gorilla/websocket (single client, multiple clients, groups, broadcast, etc.)
- [x] Support for [go-rabbitmq](https://github.com/MQEnergy/go-rabbitmq) message queue component implemented based on rabbitmq's official [amqp](https://github.com/streadway/amqp) consumer and producer encapsulation
- [x] Support for [casbin](https://github.com/casbin/casbin) rbac permissions integrated in middleware [casbin_auth.go](https://github.com/MQEnergy/gin-framework/blob/main/middleware/casbin_auth.go)
- [x] Support for [requestId](https://github.com/gin-contrib/requestid) middleware that implements convenient tracking log middleware [requestid_auth.go](https://github.com/MQEnergy/gin-framework/blob/main/middleware/requestid_auth.go)
- [x] Support for [viper](https://github.com/spf13/viper) configuration file parsing component for yaml, json, toml, etc.
- [x] Support for [validator](https://github.com/go-playground/validator) data field validator component, supporting Chinese language
- [x] Support for [snowflake](https://github.com/bwmarrin/snowflake) generate globally unique IDs with snowflake algorithm
- [x] Implementation of ip whitelist configuration integrated in middleware [ip_auth.go](https://github.com/MQEnergy/gin-framework/blob/main/middleware/ip_auth.go)
- [x] Implementation of [ticker](https://github.com/MQEnergy/gin-framework/blob/main/pkg/util/ticker.go) timer component
- [x] Implementation of [pagination](https://github.com/MQEnergy/gin-framework/blob/main/pkg/paginator/pagination.go) builder component based on gorm
- [x] Implementation of [code](https://github.com/MQEnergy/gin-framework/tree/main/pkg/response/code.go) unified defined return code, [exception](https://github.com/MQEnergy/gin-framework/tree/main/pkg/response/exception.go) unified error return handling component
  
#### Next step plans:
- [ ] Support cron for schedule tasks
- [ ] Support performance analysis component [pprof](https://github.com/gin-contrib/pprof)
- [ ] Support internal link tracking for trace project
- [ ] Support interface flow control component [rate](https://pkg.go.dev/golang.org/x/time/rate)
- [ ] Support [grpc](https://github.com/grpc/grpc-go) rpc component

# II. Start the service.

```
Note that before starting, you need to enable the MySQL and Redis services, and configure the MySQL and Redis settings in the config.dev.yaml file (which defaults to the dev environment).
```

## 1、Install dependencies and initialize.

```bash
go mod tidy 
```

## 2、Service started.

```bash
go run main.go 

# View parameters of main.go.
go run main.go --help
```

## 3、The visit indicates that successful startup.
Request：http://127.0.0.1:9527/ping
```json
{
    "status": 200,
    "errcode": 0,
    "requestid": "9ac7f4f2-1271-4f87-8df7-599a478af9cb",
    "message": "Pong!",
    "data": ""
}
```

## 4、Install hot updates.

```bash
go install github.com/cosmtrek/air@latest
```

Type in the command line: "air" to execute hot update. The code edits will update automatically.

## 5、Deploy Casbin authorization (Important! (Follow the steps below))
```
This step is intended for accessing permission for backend interfaces.
```
### 1）Perform migrate.
```bash
go run main.go migrate -s=all

# Please check the specific parameters in the "help" section.
go run main.go migrate -help
```
### 2）Request the `/routes` interface.
```
This interface will create a super administrator role based on casbin.
```

## 6、Packaging and Launching Online

```bash
# Viewing the make command line
make help

# Basic packaging to generate executable file (based on the current system)
make build

# Packaging for Windows
make windows

# Packaging for Darwin
make darwin

# Packaging for Linux
make linux
```
View packaged files in the releases section.

# III、Component Usage

## 1、Query pagination constructor based on GORM.

```go
import "github.com/MQEnergy/gin-framework/pkg/paginator"
```

### I、Basic Usage
#### 1）Basic usage of pagination for a single table:

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

#### 2）Usage of join queries in databases:
Define receiving struct.
```go
type BaseUser models.GinUser
type GinUserInfo models.GinUserInfo

// UserList get user lists.
type UserList struct {
	BaseUser
	GinUserInfo `gorm:"foreignKey:user_id" json:"user_info"`
}
```

Usage 1:
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

Usage two:
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

#### 3）Preloading query usage (strongly recommended usage):

```
Note:
There is a slight difference between the struct defined by the preload method and the struct defined by the joins query. The name of the struct defined by the preload method must be consistent with the name of the struct for the current table of the model. Also, the name of the struct for the associated table cannot be the same as the name of the struct for the model. For example, the definition of "UserInfo" is as follows:
```    

Definition of receiving a struct.
```go
type BaseUser models.GinUser
type GinUserInfo models.GinUserInfo

type GinUser struct {
	BaseUser
	UserInfo GinUserInfo `gorm:"foreignKey:user_id" json:"user_info"`
}
```

Usage instructions as follows:
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
This method does not recommend using WithFields or WithField to query fields. Instead, it is recommended to directly define the query fields specified by the receiving struct.
```

Visit address: http://127.0.0.1:9527/user/index?page=1 The returned data format is as follows:

```json
{
  "status": 200,
  "errcode": 0,
  "requestid": "9ac7f4f2-1271-4f87-8df7-599a478af9cb",
  "message": "Request Success",
  "data": {
    "list": [],
    "current_page": 1,
    "total": 2,
    "last_page": 1,
    "per_page": 10
  }
}
```
#### 4）View case:
1) Usage: Retrieve User List:
```
entities/user/gin_user.go
app/controller/backend/user.go
app/service/backend/user.go
router/routes/common.go
```
### II. Specific Methods
<details>
<summary>View usage.</summary>

#### 1）`Must be used in a chained operation.` DB Connection Method
```go
WithDB(db *gorm.DB) *PageBuilder
```

#### 2）`Must be used in a chained operation.` Model connection method.
```go
WithModel(model interface{}) *PageBuilder
```

Pass the main table model as a parameter for the query, for example: models.GinAdmin. You cannot pass a structure address as a parameter, like &models.GinAdmin.

#### 3）`Not necessary in chained operations.` Single table query or filtering field method.
```go
WithField(fields []string) *PageBuilder 
```

The last parameter of "fields" defaults to "_select" (can be omitted), but if "_omit" is passed, it filters out the fields transmitted earlier.

Note:
- _select / _omit must be the last parameters.
- WithModel parameter cannot be passed as a pointer to a struct. For example: &models.GinAdmin must be models.GinAdmin, otherwise the _omit parameter will be ineffective.
- This note applies to the WithFields method and WithMultiFields method.

Usage as follows:
```go
// Indicates filtering of front fields.
WithField([]string{"created_at", "updated_at", "_omit"})

// Indicates the fields to be queried.
WithField([]string{"created_at", "updated_at", "_select"})
WithField([]string{"created_at", "updated_at"})
```

#### 4）`Not necessary in chain operations.` Multi-table queries or filtering field methods (when using the preload mode, there may be issues with associated table queries, and it is not recommended to use the preload association query method).
```go
WithFields(model interface{}, table string, fields []string) *PageBuilder
```
The last argument for fields is default to "_select" (can be omitted), if "_omit" is passed, it will filter the fields transmitted earlier.

Usage as follows：
```go
// Indicates filtering the preceding field.
WithFields(models.GinUser{}, models.GinUserTbName, []string{"password", "salt", "_omit"})

// "Indicates the field before the search."
WithFields(models.GinUserInfo{}, models.GinUserInfoTbName, []string{"id", "user_id", "role_ids", "_select"})
WithFields(models.GinUserInfo{}, models.GinUserInfoTbName, []string{"id", "user_id", "role_ids"})
```

#### 5）`Not necessary in chain operations.` Multiple table and field queries (can replace WithFields method)
```go
WithMultiFields(fields []SelectTableField) *PageBuilder
```

Usage as follows：
```go
WithMultiFields([]paginator.SelectTableField{
    {Model: models.GinUser{}, Table: models.GinUserTbName, Field: []string{"password", "salt", "_omit"}},
    {Model: models.GinUserInfo{}, Table: models.GinUserInfoTbName, Field: []string{"id", "user_id", "role_ids"}},
})
```

#### 6）`Not necessary in chain operations.` Multiple table association query with proactive preloading (conditions not currently supported).
    
```go
 WithPreloads(querys []string) *PageBuilder 
```

Usage as follows：
```go
WithPreloads([]string{"UserInfo", "UserRecord"})
```

#### 7）`Not necessary in chain operations.` Active preloading of associated queries (can pass conditions, conditions refer to Gorm).

```go
WithPreload(query string, args ...interface{}) *PageBuilder
```

Usage as follows：
```go
WithPreload("UserInfo", "user_id = ?", "1")
```

#### 8）`Not necessary in chain operations.` Method for Data Query Conditions.

```go
WithCondition(query interface{}, args ...interface{}) *PageBuilder
```

Pass in query conditions that support the querying format used in the 'where' conditions in GORM (non-struct format). The 'query' and 'args' parameters should follow the same method of passing in 'where' conditions as in GORM.

#### 9）`Not necessary in chain operations.` Method for Data Query Conditions.

```go
WithJoins(joinType string, joinFields []OnJoins) *PageBuilder
```

joinType: Type of join. It can be left, right or inner. joinFields structure: LeftTableField, such as the main table's ID, and RightTableField, such as the related table's main table ID.

Usage as follows：
```go
WithJoins("left", []paginator.OnJoins{{
    LeftTableField:  paginator.JoinTableField{Table: models.GinUserTbName, Field: "id"},
    RightTableField: paginator.JoinTableField{Table: models.GinUserInfoTbName, Field: "user_id"},
}})
```

#### 10）`The last link must be in the chain operation.` Page return method.

```go
Pagination(dst interface{}, currentPage, pageSize int) (Page, error)
```

"dst" is a "struct" structure that receives incoming data. Note: It must be passed in application mode, such as "&userList". "model" and "currentPage" represent the current page number, and "pageSize" represents the number of queries per page.

#### 11）`Not necessary in chain operations.` Connecting with native query method.

```go
 NewDB() *gorm.DB
```
After using this method, all methods inside the pagination for chained operations are unavailable. You can continue using the native methods of GORM afterwards.

Usage as follows：
```go
NewDB().Where("id = ?", id).First(&userList)
```

#### 12）Get current page number.

```go
paginator.CurrentPage
```

#### 13）Get paginated list.

```go
paginator.List
```

#### 14）Get total number of data.

```go
paginator.Total
```

#### 15）Get the last page number.

```go
paginator.LastPage
```

#### 16）Get the number of data per page.

```go
paginator.PerPage
```
</details>

## 2、Upload component based on Gin.

```go
UploadFile(path string, r *gin.Context) (*FileHeader, error)
```

By default, the files are stored in the "upload" directory of the project. If it does not exist, it will be created automatically. The path will be "upload" directory module directory, such as "user". The directory structure will be "upload/user/{yyyy-mm-dd}/...".

Usage as follows：
```
app/controller/backend/attachment.go
pkg/util/upload.go
router/routes/common.go
```
## 3、Using RabbitMQ Component
Configure the amqp parameters in the yaml configuration file.
### 1）Start the consumer.
Test case
```shell
go run command/test/consumer.go
```
### 2）Start the producer.
Test case
```shell
go run command/test/producer.go
```

# IV、Tools

Running "go run main.go --help" will display the following command set.

```
COMMANDS:
  migrate     Create a migration command
  account     Create a new admin account
  model       Create a new model class
  controller  Create a new controller class
  service     Create a new service class
  help, h     Shows a list of commands or help for one command
```

## 1、Execute migrate.
<details>
<summary>View Usage</summary>

```bash
# Install migrate cli tool.
curl -L https://github.com/golang-migrate/migrate/releases/download/$version/migrate.$platform-amd64.tar.gz | tar xvz

# MacOS Installation
brew install golang-migrate

# Install Window with scoop https://scoop.sh/
scoop install migrate

# "Creating migration file syntax, for example:"
migrate create -ext sql -dir migrations -seq create_users_table

# The first way to execute migration.
# Carrying out migration operation:
migrate -database 'mysql://root:123456@tcp(127.0.0.1:3306)/gin_framework' -path ./migrations up
# Performing rollback operation:
migrate -database 'mysql://root:123456@tcp(127.0.0.1:3306)/gin_framework' -path ./migrations down

# The second way to perform migration.
# View help command.
go run main.go migrate --help

# The format is as follows：
go run main.go migrate -s {step} -e {env}
# env: Keep consistent with the config.*.yaml file among dev, test, and prod. The default is dev.
# step: The number of migration files to be executed (or rolled back), for example, 1, 2, 3... If you want to execute all, use "all".

# Perform all migration operations:
go run main.go migrate -s all

# Perform partial migration operations：
# eg.：go run main.go migrate -s 1

# Performing rollback operation：
# eg.：go run main.go migrate -s -1
```
</details>

## 2、Automatically generate models.

```bash
# Execute the command to generate all models.
go run main.go model -tb=all {env}

# Please refer to the help for specific parameters.
go run main.go model -help
```

## 3、Auto-generate controller.

```bash
go run main.go controller -c={controller name} -m={module name}
# eg.：go run main.go controller -c=admin -m=backend

# Please refer to the help for specific parameters.
go run main.go controller -help
```

## 4、Automatically generating service.

```bash
go run main.go service -s={service name} -m={module name}
# The module name is the module named "name" located in the app/controller directory.
# eg.：go run main.go service -s=admin -m=backend

# Please refer to the help for specific parameters.
go run main.go service -help
```

## 5、Create a back-end administrator account (based on the gin_admin table, you can modify the code based on other tables as needed).

```bash
go run main.go account -c={account name} -p={password}  

# Check the specific parameters in the help documentation.
go run main.go account -help
```