# sleep的内部实现

从`time/sleep.go`源文件可以看出，是调用了`runtime/time.go`文件中的`timeSleep`。
```gotemplate
//文件：$GOPATH/src/time/sleep.go

// Sleep pauses the current goroutine for at least the duration d.
// A negative or zero duration causes Sleep to return immediately.
func Sleep(d Duration)
```

