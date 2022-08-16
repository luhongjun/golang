# Golang 的设计模式

**推荐阅读**

- [Golang设计模式](http://books.studygolang.com/go-patterns)
- [Github - google/wire, Go的编译时依赖注入的工具](https://github.com/google/wire)

| 类型 | 设计模式 | 代码示例 | 
| --- | --- | --- | 
| 创建型模式 | 单例模式 | [代码](./singleton.go) |
| 创建型模式 | 对象池 | [代码](./object_pool.go)，也可以用数组来实现的 |
| 创建型模式 | 工厂方法 | [代码](./factory_method.go) |
| 创建型模式 | 构建器 | [代码](./builder.go) |
| 结构型模式 | 装饰器模式 | [代码](./decorator.go)，支持运行时动态改变现有对象（对象指的是函数）功能 | 
| 结构型模式 | 代理模式 | [代码](./proxy.go)，不支持动态修改对象功能，只能通过新定义类去覆盖父类的方法去改变 |
| 行为型模式 | （Observer）观察者模式 | [代码](./observer.go) |
| 行为型模式 | （Strategy）策略模式 | [代码](./strategy.go) |
| 同步型模式 | （Semaphore Pattern）信号量模式 | [代码](./semaphore.go) |
| 并发型模式 | （Generator Pattern）生成器模式 | [代码](./generator.go) |
| 并发型模式 | （Parallelism Pattern）并行模式 | 代码忽略。并行模式允许多个作业或者任务同时异步地进行，本质是通过golang的协程来实现的 |
| 并发型模式 | （Bounded Parallelism Pattern）有界限的并行模式 | [代码](./bounded_parallelism.go) |
| 消息传递 | （Fan-In Messaging Patterns）扇形传入消息模式 | [代码](./fan_in_messaging.go) |
| 消息传递 | （Fan-Out Messaging Pattern）扇形传出消息模式 | [代码](./fan_out_messaging.go) |
| 消息传递 | （Publish & Subscribe Messaging Pattern） | [代码](./publish_and_subscribe_messaging.go) ，与观察者设计模式不同，实现方式是通过golang的信号 |
| 稳定型模式 | （Circuit Breaker Pattern）断路器模式 | [代码](./circuit_breaker.go) |
| 函数 | （Functional Options）函数式选项模式 | [代码](./function.go) | 

## 设计模式：实现总耗时的统计
```gotemplate
func Duration(invocation time.Time, name string) {
    elapsed := time.Since(invocation)

    log.Printf("%s lasted %s", name, elapsed)
}

// 延期执行，最后可以算到运行的总时间
defer Duration(time.Now(), "IntFactorial")
```

## 设计模式：链式方法
```gotemplate
type Student struct {
	name	string
	number	int32
}

func (student *Student) SetName(name string) *Student {
    student.name = name
    return student
}

func (student *Student) SetNumber(number int32) *Student {
    student.number = number
    return student
}

student := new(Student)
student.SetName("小明").SetNumber("078")
```

## [Wire —— 依赖项注入自动连接组件]()

---------------------------

- 基于 sync.Pool 临时对象池的设计，适合用作针对某种数据的缓存，有利于复用已申请过的内存空间
    - fmt.Fprintln / fmt.Sprintln 等，fmt包定义了私有变量 ppFree 来提供 Get/Put 打印器
    - gin.ServeHTTP，每次发起请求时都会复用 Context
    
**解释：临时对象池是怎样利用内部数据结构来存取值的**
```text
临时对象池的Put方法总会先试图把新的临时对象，存储到对应的本地池的private字段中，以便在后面获取临时对象的时候，可以快速地拿到一个可用的值。

只有当这个private字段已经存有某个值时，该方法才会去访问本地池的shared字段。

相应的，临时对象池的Get方法，总会先试图从对应的本地池的private字段处获取一个临时对象。只有当这个private字段的值为nil时，它才会去访问本地池的shared字段。

一个本地池的shared字段原则上可以被任何goroutine中的代码访问到，不论这个goroutine关联的是哪一个P。这也是我把它叫做共享临时对象列表的原因。

相比之下，一个本地池的private字段，只可能被与之对应的那个P所关联的goroutine中的代码访问到，所以可以说，它是P级私有的。

以临时对象池的Put方法为例，它一旦发现对应的本地池的private字段已存有值，就会去
访问这个本地池的shared字段。当然，由于shared字段是共享的，所以此时必须受到互斥
锁的保护。
```