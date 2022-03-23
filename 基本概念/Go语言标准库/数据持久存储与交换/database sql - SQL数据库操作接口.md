# database/sql - SQL/SQL-Like 数据库操作接口

这是 Go 提供的操作 SQL/SQL-Like 数据库的通用接口，但 Go 标准库并没有提供具体数据库的实现，需要结合第三方的驱动来使用该接口。本书使用的是 mysql 的驱动：github.com/go-sql-driver/mysql。

**注：**
该包有一个子包：driver，它定义了一些接口供数据库驱动实现，一般业务代码中使用 database/sql 包即可，尽量避免使用 driver 这个子包。

## database/sql 是什么

很明显，database/sql 首先是 Go 标准库提供的一个包，用于和 SQL/SQL-Like 数据库 ( 关系或类似关系数据库）通讯。它提供了和 ODBC、Perl 的 DBI、Java 的 JDBC 和 PHP 的 PDO 类似的功能。然而，它的设计却不太一样，掌握了它有利于构建健壮、高性能的基于 database 的应用。

另一方面，database/sql 提供的是抽象概念，和具体数据库无关，具体的数据库实现，有驱动来做，这样可以很方便的更换数据库。

该包提供了一些类型（概括性的），每个类型可能包括一个或多个概念。

### DB

`sql.DB`类型代表了一个数据库。这点和很多其他语言不同，它并不代表一个到数据库的具体连接，而是一个能操作的数据库对象，具体的连接在内部通过连接池来管理，对外不暴露。这点是很多人容易误解的：每一次数据库操作，都产生一个`sql.DB`实例，操作完`Close`。

DB是一个数据库句柄，代表一个具有零到多个底层连接的连接池，它可以安全的被多个 goroutine 同时使用。

**底层处理逻辑：**
sql 包会自动创建和释放连接；它也会维护一个闲置连接的连接池。如果数据库具有单连接状态的概念，该状态只有在事务中被观察时才可信。一旦调用了 BD.Begin，返回的 Tx 会绑定到单个连接。当调用事务 Tx 的 Commit 或 Rollback 后，该事务使用的连接会归还到 DB 的闲置连接池中。连接池的大小可以用 SetMaxIdleConns 方法控制。

----------------------

由于 DB 并非一个实际的到数据库的连接，而且可以被多个 goroutine 并发使用，因此，程序中只需要拥有一个全局的实例即可。所以，经常见到的示例代码：
```gotemplate
db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/test?charset=utf8")
if err != nil {
    panic(err)
}
defer db.Close()
```

- 关于`db.Close`说明

实际中，`defer db.Close()`可以不调用，官方文档关于`DB.Close`的说明也提到了：Close 用于关闭数据库，释放任何打开的资源。一般不会关闭 DB，因为 DB 句柄通常被多个 goroutine 共享，并长期活跃。当然，如果你确定 DB 只会被使用一次，之后不会使用了，应该调用 Close。

所以，实际的 Go 程序，应该在一个 go 文件中的 init 函数中调用 sql.Open 初始化全局的 sql.DB 对象，供程序中所有需要进行数据库操作的地方使用。

- 关于`sql.Open`说明

`sql.DB`并不是实际的数据库连接，因此，sql.Open 函数并没有进行数据库连接，只有在驱动未注册时才会返回 err != nil。

例如：db, err := sql.Open("mysql", "root:@tcp23(localhost233:3306)/test?charset=utf8")。虽然这里的 dsn 是错误的，但依然 err == nil，只有在实际操作数据库（查询、更新等）或调用 Ping 时才会报错。

关于 Open 函数的参数，第一个是驱动名，为了避免混淆，一般和驱动包名一致，在驱动实现中，会有类似这样的代码：
```gotemplate
func init() {
    sql.Register("mysql", &MySQLDriver{})
}
```
其中`mysql`即是注册的驱动名。由于注册驱动是在`init`函数中进行的，这也就是为什么采用`_ "github.com/go-sql-driver/mysql" `这种方式引入驱动包。第二个参数是DSN（数据源名称），这个是和具体驱动相关的，database/sql 包并没有规定，具体书写方式参见驱动文档。

### Results

定义了三种结果类型：sql.Rows、sql.Row 和 sql.Result，分别用于获取多个多行结果、一行结果和修改数据库影响的行数（或其返回 last insert id）。

### Statements

`sql.Stmt`代表一个语句，如：DDL、DML 等。

### Transactions

`sql.Tx`代表带有特定属性的一个事务。

