# Go IRIS

> about `go-iris` usage.

### 模块定义

> go.mod

```text
require (
	github.com/uniondrug/gwf v1.0.2
)
```

### 引用模块

> 在Go源码中

```text
import (
    "github.com/uniondrug/gwf"
)
```

## 关于项目

1. 目录结构.
   ```text
   /- data/apps/sketch
   ├─ app
   │   ├── controllers
   │   │   └── index_controller.go
   │   ├── logics
   │   │   └── index
   │   │       └── index_logic.go
   │   ├── middlewares
   │   │   └── example_middleware.go
   │   ├── models
   │   │   └── example.go
   │   └── services
   │       └── example_service.go
   └── config
      ├── app.yaml
      ├── logger.yaml
      └── service.yaml
   ```
1. Config配置.
    1. `app.yaml` - 用于应用级配置, 如项目名、版本号、端口号
        1. `addr` - 服务地址, 默认: `sketch`.
        1. `name` - 项目名称, 默认: `sketch`.
        1. `version` - 版本号, 默认: `0.0.0`.
    1. `logger.yaml` - 用于配置日志处理方式.
        1. `adapter-name` - 适配器名称, 支持 `term`、`file`、`redis`, 默认: `term`.
        1. `level-name` - 日志级别, 支持 `DEBUG`、`INFO`、`WARN`、`ERROR`, 默认: `INFO`.
        1. `time-format` - 日志时间格式, 默认: `2006-01-02 15:04:05.999`.
    1. `service.yaml` - 用于配置数据库连接方式.
        1. `driver` - DB适配器名称.
        1. `dsn` - DB连接源.
        1. `max-idle` - 最大空闲连接数, 默认: `5`.
        1. `max-open` - 最大连接数, 默认: `300`.
        1. `max-lifetime` - 连接空闲多长时间时断开, 默认: `60`.
        1. `mapper` - 表字段映射, 支持 `same`、`snake`
1. 模型/Model
    1. 文件位于 `app/models` 目录下.
    1. 文件无需加额外后缀, 如: `example.go`.
    1. 命名规范
        1. 文件名为表名(不包括前缀后), 如: `Example` 表示有 `example` 数据表.
        1. 字段定义用于映射字段名, 规范是`大驼峰`, 如: `MemberId` 表示数据表有`member_id`字段.
        1. 使用xorm标签映射字段名.
            1. 主键 `xorm:"pk autoincr"`
            1. 映射 `xorm:"member_id"` 表示, 指定使用`member_id`映射到`MemberId`字段上.
    1. 功能说明
        1. 使用常量定义状态. 如: `const ExampleStatusEnabled = 1`
        1. 使用 `TableName` 指定表名, 如: `func (m *Example) TableName() string {...}`
        1. 使用 `Is` 前缀表示状态等属性, 如 `func (m *Example) IsEnabled() bool {...}`
1. 业务/Service
    1. 文件位于 `app/services` 目录下.
    1. 文件需以 `_service.go` 为文件名后缀, 如: `example_service.go`.
    1. 命名规范
        1. 结构体以 `Service` 为后缀, 如: `ExampleService`.
        1. 按规范须以 `New` 前缀加业务名称定义构造函数, 如: `func NewExampleService(sess ... *xorm.Session) *ExampleService {...}`.
        1. 读方法
            1. `GetById(id int64) (*Example, error)`
            1. `ListByStatus(status int, limit, page int) ([]*Example, error)`
        1. 写方法
            1. `Add(req *Example) error`
            1. `SetStatusAsEnabled(model *Example) error`
1. 控制器/Controller.
    1. 文件位于 `app/controllers` 目录下.
    1. 文件需以 `_controller.go` 为文件名后缀, 如: `example_controller.go`.
    1. 命名与注释
        1. 结构体以 `Controller` 为后缀, 如: `ExampleController`.
        1. 注释使用 `@Post()`、`@Post()` 表示路由与路径, 用于导出文档, 如: `// @Post(/example/index)`.
        1. 注释使用 `@Request()` 表示入参, 用于导出文档, 如: `// @Request(AddExampleRequest)`.
        1. 注释使用 `@Response()` 表示出参, 用于导出文档, 如: `// @Response(AddExampleResponse)`.
    1. 方法定义
        1. 参入固定 `ctx iris.Context`, 如: `func (c *ExampleController) PostIndex(ctx iris.Context) interface{} {...}`
        1. 返回固定 `interface{}`, 示例同上.
        1. 前缀使用 `Post`、`Get` 开头, 表示请求方式, 如上例必须以 `POST` 方式请求.
1. 逻辑.
    1. 文件位于 `app/logics` 目录下.
    1. 文件需以 `_logic` 为文件名后缀, 如: `example/add_logic.go`
    1. 命名规范
        1. 入参需要以 `Request` 为后缀, 如: `type AddRequset struct`.
        1. 出参需要以 `Response` 为后缀, 如: `type AddResponse struct`.
        1. 业务逻辑以 `Logic` 为后缀, 如: `type AddLogic struct`.
        1. 按规范须以 `New` 前缀加业务逻辑名称定义构造函数, 如: `func NewAddLogic() *AddLogic {...}`.
        1. 按规范须以 `Run()` 方法定义主逻辑, 如: `func (o *AddLogic) Run(ctx iris.Context) interface{} {...}`
        1. 按规范须以 `interface{}` 类型作为 `Run()` 方法的返回类型.
1. 应用
    1. Logger.
    1. 事务
    1. context.Context
    1. channel
