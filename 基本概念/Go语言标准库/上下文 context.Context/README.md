# Context 初识

Go1.7加入了一个新的标准库context，它定义了Context类型，专门用来简化 对于处理单个请求的多个 goroutine 之间与请求域的数据、取消信号、截止时间等相关操作，这些操作可能涉及多个 API 调用。

对服务器传入的请求应该创建上下文，而对服务器的传出调用应该接受上下文。它们之间的函数调用链必须传递上下文，或者可以使用WithCancel、WithDeadline、WithTimeout或WithValue创建的派生上下文。当一个上下文被取消时，它派生的所有上下文也被取消。


## Context 接口

在 Go 里并没有直接为我们提供一个统一的 context 对象，而是设计了一个接口类型的 Context。然后在这些接口上来实现了几种具体类型的 context。

这样的好处就是我们只要根据开放出来的接口定义，也能够实现属于自己的 context，进而跟官方的 context 一起配合使用。

在分析官方的几种 context 之前，我们先来看看 context 要求实现的几个接口：
``` 
//如果有截止时间的话，得返回对应 deadline 时间；如果没有，则 ok 的值为 false。
Deadline() (deadline time.Time, ok bool)

//关于 channel 的数据通信，而且它的数据类型是 struct{}，一个空结构体，因此在 Go 里都是直接通过 close channel 来进行通知的，不会涉及具体数据传输
Done() <-chan struct{}

//返回的是一个错误 error，如果上面的 Done() 的 channel 没被 close，则 error 为 nil；如果 channel 已被 close，则 error 将会返回 close 的原因，比如超时或手动取消
Err() error

//用来存储具体数据的方法
Value(key interface{}) interface{}
```

## Context 类型

官方的 context 类型,主要有四种，分别是 emptyCtx，cancelCtx，timerCtx，valueCtx：

- emptyCtx：空的 context，实现了上面的 4 个接口，但都是直接 return 默认值，没有具体功能代码
```gotemplate
context.Background()
context.TODO()
```

- cancelCtx：取消通知用的 context

```gotemplate
//从 context.Context 中衍生出一个新的子上下文并返回用于取消该上下文的函数。一旦我们执行返回的取消函数，当前上下文以及它的子上下文都会被取消，所有的 Goroutine 都会同步收到这一取消信号
func context.WithCancel(parent context.Context) (ctx context.Context, cancel CancelFunc)
```

- timerCtx：超时通知用的 context

```gotemplate
func context.WithDeadline(parent context.Context, d time.Time) (context.Context, CancelFunc)
```

- valueCtx：用来传值的 context

```gotemplate
func context.WithValue(parent context.Context, key, val interface{}) context.Context
```

## 总结

context 的设计类似于链式的设计（层层相套）