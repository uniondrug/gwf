# Transaction

> XDB事务中，回调函数按入参先、后顺序执行，执行过程中当有回调出现`error`类型返回时，事务回滚。当`logger`级别开启`INFO`及
> 以上时, SQL语句也将写入到日志中.

```text

func (o *ExampleLogic) Run(ctx iris.Context) interface{} {
    db := xdb.MasterContext(ctx)
    if err := xdb.Transaction(ctx, db, o.f1, o.f2); err != nil {
        return xapp.WithError(ctx, 1, err)
    }
    return xapp.WithData(ctx, true)
}

func (o *ExampleLogic) f1(ctx interface, sess *xorm.Session) error {
    return nil
}

func (o *ExampleLogic) f2(ctx interface, sess *xorm.Session) error {
    return nil
}

```
