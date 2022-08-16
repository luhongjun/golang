# testing.TB 接口

TB接口，顾名思义，是testing.T(单元测试)和testing.B(性能测试)共用的接口。

TB接口通过在接口中定义一个名为private(）的私有方法，保证了即使用户实现了类似的接口，也不会跟testing.TB接口冲突。

其实，这些接口在testing.T和testing.B公共成员testing.common中已经实现。

其数据结构如下：
```go
package testing

// TB is the interface common to T and B.
type TB interface {
    Error(args ...interface{})
    Errorf(format string, args ...interface{})
    Fail()
    FailNow()
    Failed() bool
    Fatal(args ...interface{})
    Fatalf(format string, args ...interface{})
    Log(args ...interface{})
    Logf(format string, args ...interface{})
    Name() string
    Skip(args ...interface{})
    SkipNow()
    Skipf(format string, args ...interface{})
    Skipped() bool
    Helper()

    // A private method to prevent users implementing the
    // interface and so future additions to it will not
    // violate Go 1 compatibility.
    // 其中私有接口private()用于控制该接口的唯一性，即便用户代码中某个类型实现了这些方法，由于无法实现这个私有接口，也不能被认为是实现了TB接口，所以不会跟用户代码产生冲突
    private()
}
```