# SQL连接池

## 连接池的工作原理

获取 DB 对象后，连接池是空的，第一个连接在需要的时候才会创建。示例如下：
```gotemplate
db, _ := sql.Open("mysql", "root:@tcp(localhost:3306)/test?charset=utf8")
fmt.Println("please exec show processlist")
time.Sleep(10 * time.Second)
fmt.Println("please exec show processlist again")
db.Ping()
time.Sleep(10 * time.Second)
```
在 Ping 执行之前和之后，show processlist 多了一条记录，即多了一个连接，Command 列是 Sleep。

**连接池的工作方式：**
当调用一个函数，需要访问数据库时，该函数会请求从连接池中获取一个连接，如果连接池中存在一个空闲连接，它会将该空闲连接给该函数；否则，会打开一个新的连接。当该函数结束时，该连接要么返回给连接池，要么传递给某个需要该连接的对象，直到该对象完成时，连接才会返回给连接池。

相关方法的处理说明（假设 sql.DB 的对象是 db），如下：

| 方法 | 处理说明 |
| --- | --- |
| db.Ping() | 将连接立马返回给连接池 |
| db.Exec() | 会将连接立马返回给连接池，但是它返回的 Result 对象会引用该连接，所以，之后可能会再次被使用 |
| db.Query() | 会传递连接给 sql.Rows 对象，直到完全遍历了所有的行或 Rows 的 Close 方法被调用了，连接才会返回给连接池 |
| db.QueryRow() | 会传递连接给 sql.Row 对象，当该对象的 Scan 方法被调用时，连接会返回给连接池 |
| db.Begin() | 会传递连接给 sql.Tx 对象，当该对象的 Commit 或 Rollback 方法被调用时，该链接会返回给连接池 |

从上面的解释可以知道，大部分时候，我们不需要关心连接不释放问题，它们会自动返回给连接池。

**注意：**
如果某个连接有问题（broken connection)，database/sql 内部会进行最多 2 次 的重试，从连接池中获取或新开一个连接来服务，因此，你的代码中不需要重试的逻辑

## 控制连接池

Go1.2.1 之前，没法控制连接池，Go1.2.1 之后，提供了两个方法来控制连接池（Go1.2 提供了控制，不过有 bug）。

- `db.SetMaxOpenConns(n int)`：设置连接池中最多保存打开多少个数据库连接。注意，它包括在使用的和空闲的。如果某个方法调用需要一个连接，但连接池中没有空闲的可用，且打开的连接数达到了该方法设置的最大值，该方法调用将堵塞。默认限制是 0，表示最大打开数没有限制。
- `db.SetMaxIdleConns(n int)`：设置连接池中能够保持的最大空闲连接的数量。默认值是 2

