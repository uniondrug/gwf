# Logger

> 日志集成了`Tracing`应用, 生成的每条日志`保障`在同一个`请求链`下, 如何保障?

### 在Logic

> 在`Logic`等`iris`框架规则中, 使用 Debugfc()、Infofc()、Warnfc()、Errorfc() 方法, 并将 `iris.Context` 为作为第1个参数传入即可.

```text

func (o *ExampleLogic) Run(ctx iris.Context) interface{} {
    var err error

    xlog.Debugfc(ctx, "logic begin.")
    defer func(){
        if err != nil {
            xlog.Warnfc(ctx, "logic error - %v.", err)
        } else {
            xlog.Debugfc(ctx, "end logic.")
        }
    }
    
    ...
    
}

```

### DB应用

```text

func (o *ExampleLogic) Run(ctx iris.Context) interface{} {
    sess := xdb.SlaveContext(ctx)
    _, _ = NewExampleService(sess).GetById(1)
}

```

### 独立应用

```text

func (o *ExampleLogic) Run(ctx iris.Context) interface{} {
    tracing := xdb.NewTracing().UseIris(ctx)
    xlog.Infofc(tracing, "tracing")
}

```


##### logger

```text
[16:36:45.268][DEBUG][s=f2ae8c7a65fa11eb836ea683e76f2900][v=0.0] middleware:tracing
[16:36:45.268][DEBUG][s=f2ae8c7a65fa11eb836ea683e76f2900][v=0.1] middleware:recover
[16:36:45.268][ INFO][s=f2ae8c7a65fa11eb836ea683e76f2900][v=0.2] 查看结算单操作日志开始.
[16:36:45.295][ INFO][s=f2ae8c7a65fa11eb836ea683e76f2900][v=0.3] [SQL][d=0.026169] SELECT `LogId`, `StatementNo` FROM `v3_statement_log` WHERE (statementNo = ?) LIMIT 30 - [DS20210201115892].
[16:36:45.295][ INFO][s=f2ae8c7a65fa11eb836ea683e76f2900][v=0.4] 查看结算单操作日志完成.
```
