# golang
归纳了Golang语言的整套知识体系，包括Golang基本概念、经典代码案列以及其他相关资料等



## 知识点导航

- 技术书籍
  - [Go Rpc开发指南.pdf](技术书籍/Go%20Rpc开发指南.pdf)
  - [Go1.5源码剖析.pdf](技术书籍/Go1.5源码剖析.pdf)
  - [Go专家编程.pdf](技术书籍/Go专家编程.pdf)
  - [Go语言圣经（中文版）.pdf](技术书籍/Go语言圣经（中文版）.pdf)
  - [Go语言学习笔记(详细书签).pdf](技术书籍/Go语言学习笔记（详细书签）.pdf)
  - [Go语言标准库.pdf](技术书籍/Go语言标准库.pdf)
  - [Go高级编程.pdf](技术书籍/Go高级编程.pdf)
- 基本概念
  - [Go命令](基本概念/Go命令/README.md) (Command go详见官网： http://docscn.studygolang.com/cmd/go/)
  - Go语言标准库
  - 包和工具
  - [历史版本的发行说明](基本概念/历史版本的发行说明/README.md) （官网 release History：https://go.dev/doc/devel/release）
    - [>> Go1.19 Release Notes](https://go.dev/doc/go1.19)
    - [>> Go1.18 Release Notes](https://go.dev/doc/go1.18)
    - [>> Go1.17 Release Notes](https://go.dev/doc/go1.17)
    - [>> Go1.16 Release Notes](https://go.dev/doc/go1.16)
- Go底层原理
  - CGO编程
  - Golang汇编
  - Go内存管理
- [Go并发编程](Go并发编程/README.md)
  - 标准库及其数据结构
    - 互斥锁 sync.Mutex
    - 读写锁 sync.RWMutex
    - 条件变量 sync.Cond
    - 全局等待 sync.WaitGroup
    - 只执行一次 sync.Once
    - 临时对象池 sync.Pool
    - 并发安全字典 sync.Map
  - [Go并发编程实战（第2版）.pdf](Go并发编程/Go并发编程实战（第2版）.pdf)
  - [Goroutine调度.docx](Go并发编程/Goroutine调度.docx)
- Go网络编程
- [单元测试](./单元测试/README.md)
  - [go test 指令详解](./单元测试/go%20test%20指令详解.md)
  - [go test的实现原理.docx](./单元测试/go%20test实现原理.docx)
  - [性能分析](./单元测试/性能分析/README.md)：
    - [基于http协议的网络服务的性能分析：net/http/pprof](./单元测试/性能分析/基于http协议的网络服务的性能分析.md)
    - [程序性能分析基础（上）.pdf](./单元测试/性能分析/程序性能分析基础（上）.pdf)
    - [程序性能分析基础（下）.pdf](./单元测试/性能分析/程序性能分析基础（下）.pdf)
  - 流行的单元测试框架：
    - [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock)：sqlmock是一个实现sql/driver的模拟库。它有一个也是唯一的目的 - 模拟测试中的任何sql驱动程序行为，而不需要真正的数据库连接
    - [monkey](https://github.com/bouk/monkey)：可以通过自定义的mock来替换函数或者方法
      其工作原理见此文章：`https://bou.ke/blog/monkey-patching-in-go/`
    - [goconvey](https://github.com/smartystreets/goconvey/wiki/Documentation)：支持golang的单元测试框架，能够自动监控文件修改并启动测试，并可以将测试结果实时输出到web界面，goconvey提供了丰富的断言简化测试用例的编写
    - [gomock](https://github.com/golang/mock)：可使用`mockgen`指令为某些模拟类生成源代码
- [Go设计模式](设计模式/README.md)






-------------------
数据库：

- gorm：https://gorm.io/zh_CN/docs/index.html

- 
-------------------
数据结构与算法：

- LeetCode：https://books.halfrost.com/leetcode/

-------------------
测试：

- Go自带testing：https://pkg.go.dev/testing





