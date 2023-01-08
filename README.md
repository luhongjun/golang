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
    - go env：获取/修改golang的环境变量获取/修改golang的环境变量
    - ...
  - [Go语言标准库](基本概念/Go语言标准库/README.md)
  - 包和工具
    - [GOROOT路径下的目录结构](基本概念/包和工具/GOROOT路径下的目录结构说明.md)
  - [历史版本的发行说明](基本概念/历史版本的发行说明/README.md) （官网 release History：https://go.dev/doc/devel/release）
    - [>> Go1.19 Release Notes](https://go.dev/doc/go1.19)
    - [>> Go1.18 Release Notes](https://go.dev/doc/go1.18)
    - [>> Go1.17 Release Notes](https://go.dev/doc/go1.17)
    - [>> Go1.16 Release Notes](https://go.dev/doc/go1.16)
- Go底层原理
  - CGO编程：Go语言通过自带的一个叫CGO的工具来支持C语言函数调用，同时我们可以用Go语言导出C动态库接口给其它语言使用
    - [CGO编程](Go底层原理/CGO编程/CGO编程.docx)
  - Golang汇编
    - [Go汇编语言.docx](Go底层原理/Golang汇编/Go汇编语言.docx)
  - Go内存管理
    - [Go语言底层机制.docx](Go底层原理/Go内存管理/Go语言底层机制.docx)
    - [Go内存管理.docx](Go底层原理/Go内存管理/Go内存管理.docx)
    - 标准库：
      - [reflect反射](Go底层原理/reflect-反射.docx)
      - [unsafe非类型安全操作](Go底层原理/unsafe-非类型安全操作.docx)
- [Go并发编程](Go并发编程/README.md)
  - 标准库及其数据结构
    - 互斥锁 sync.Mutex
    - 读写锁 sync.RWMutex
    - 条件变量 sync.Cond
    - 全局等待 sync.WaitGroup
    - 只执行一次 sync.Once
    - 临时对象池 sync.Pool
    - 并发安全字典 sync.Map
    - 原子 sync/atomic
  - [Go并发编程实战（第2版）.pdf](Go并发编程/Go并发编程实战（第2版）.pdf)
  - [Goroutine调度.docx](Go并发编程/Goroutine调度.docx)
- Go网络编程 //@todo
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
- 数据结构与算法：
  - 数据结构：
    - [heap堆](数据结构与算法/数据结构/heap堆.md)
    - [list链表](数据结构与算法/数据结构/list链表.md)
    - [ring环](数据结构与算法/数据结构/ring环.md)
  - 算法：
    - [>> LeetCode](https://books.halfrost.com/leetcode/)
- [Go设计模式](设计模式/README.md)
  - 其他设计模式参考
    - [链式方式](设计模式/链式方式.md)
    - [wire依赖注入](设计模式/Wire依赖项注入编译组件.md)
  - 主要的设计模式

| 设计模式  | 示例                                                           |
|-------|--------------------------------------------------------------|
| 创建型模式 | [单例模式](设计模式/单例模式.md)、[对象池](设计模式/对象池.md)、[工厂方法](设计模式/工厂方法.md) |
| 结构型模式 | [装饰器模式](设计模式/装饰器模式.md)、[代理模式](设计模式/代理模式.md)                  |
| 行为型模式 | [观察者模式](设计模式/事件订阅与通知.md)、[策略模式](设计模式/策略模式.md)                |
| 同步型模式 | [信号量](设计模式/信号量.md)                                           |
| 稳定型模式 | [断路器模式](设计模式/断路器模式.md)                                       |
| 函数    | [函数式选项模式](设计模式/函数式编程.md)                                     |

- Golang的开发工具，提高效能
  - goup






-------------------
数据库：

- gorm：https://gorm.io/zh_CN/docs/index.html

- 
-------------------
数据结构与算法：

- 






