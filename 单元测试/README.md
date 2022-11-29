# Golang 单元测试

Golang有自带的 [testing库包](https://pkg.go.dev/testing) 来提供单元测试的基本方法，编写完测试用例后再通过`go test`来运行用例。

Golang 语言原生支持了单元测试，使用上非常简单，测试代码只需要放到以 `_test.go` 结尾的文件中即可。

**推荐阅读**
- [Golang 官方包"testing"](https://pkg.go.dev/testing)
- [Go语言标准库讲解 - 测试](http://books.studygolang.com/The-Golang-Standard-Library-by-Example/chapter09/09.0.html)

## testing 官方标准库

Go编写测试可以分为三类：
- 单元测试
- 性能测试
- 功能测试
- 模糊测试(Go1.18以后)

## go test 命令

[go test指令详解](go%20test%20指令详解.md)

## 流行的测试框架

| 框架 | 文档 | 概述 |
| --- | --- | --- |
| goconvey | https://github.com/smartystreets/goconvey/wiki/Documentation | - |
| go-sqlmock | https://github.com/DATA-DOG/go-sqlmock | sqlmock是一个实现sql/driver的模拟库。它有一个也是唯一的目的 - 模拟测试中的任何sql驱动程序行为，而不需要真正的数据库连接 |
| monkey | https://github.com/bouk/monkey | 可以通过自定义的mock来替换函数或者方法 <br> 其工作原理见此文章：https ://bou.ke/blog/monkey-patching-in-go/ |
| gomock | https://github.com/golang/mock |
