# Golang 的设计模式

**推荐阅读**

- [Golang设计模式](http://books.studygolang.com/go-patterns)

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