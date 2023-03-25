## 调用栈追踪工具

- runtime.Caller：Golang 中用于获取当前 goroutine 调用栈信息的函数。该函数可以获取指定层数的调用者（称为 caller）的文件名、行号和函数名等信息，常用于错误追踪和日志记录等场景
```gotemplate
/**
    skip 参数表示跳过的堆栈帧数，如果传入 0，则返回的是 Caller 函数的调用者信息
    pc 表示程序计数器，即函数指针
    file 表示文件名
    line 表示代码行号
    ok 表示是否成功获取到调用者信息
**/
func Caller(skip int) (pc uintptr, file string, line int, ok bool)

/**
注意：runtime.Callers 只能在使用了 -race 编译标志时获得正确的结果，否则可能会出现错误的调用栈信息
**/
func Callers(skip int, pc []uintptr) int
```

- runtime.FuncForPC：返回与程序计数器（Program Counter，pc）对应的函数对象

在 Golang 中，每个函数都有一个唯一的地址，可以使用 funcName.Pointer() 获取。而 runtime.FuncForPC 函数则可以通过该地址获取到对应的函数对象。
```gotemplate
func FuncForPC(pc uintptr) *Func

//返回函数的指针地址
func (f *Func) Entry() uintptr

//返回函数所在的文件以及行号
func (f *Func) FileLine(pc uintptr) (file string, line int)

//返回函数名称
func (f *Func) Name() string
```

- runtime.CallersFrames：获取 Call Stack 中所有调用栈信息的函数。它可以用于跟踪程序在运行时的调用流程，通常用于调试和性能分析。
```go
/**
堆栈地址转换成可读的Function结构体切片（Frames），其中每个Function结构体包含文件名、函数名、行号等信息，可以通过 Frames.Next() 方法逐个遍历 Function 切片中的元素，直到结束为止
**/
func CallersFrames(callers []uintptr) *Frames
func (ci *Frames) Next() (frame Frame, more bool)

/**
Frame：可读的Function结构体切片（Frames），其中每个Function结构体包含文件名、函数名、行号等信息
**/
type Frame struct {
    // PC is the program counter for the location in this frame.
    // For a frame that calls another frame, this will be the
    // program counter of a call instruction. Because of inlining,
    // multiple frames may have the same PC value, but different
    // symbolic information.
    PC uintptr
    
    // Func is the Func value of this call frame. This may be nil
    // for non-Go code or fully inlined functions.
    Func *Func
    
    // Function is the package path-qualified function name of
    // this call frame. If non-empty, this string uniquely
    // identifies a single function in the program.
    // This may be the empty string if not known.
    // If Func is not nil then Function == Func.Name().
    Function string
    
    // File and Line are the file name and line number of the
    // location in this frame. For non-leaf frames, this will be
    // the location of a call. These may be the empty string and
    // zero, respectively, if not known.
    File string
    Line int
    
    // Entry point program counter for the function; may be zero
    // if not known. If Func is not nil then Entry ==
    // Func.Entry().
    Entry uintptr
    // contains filtered or unexported fields
}
```

- runtime.Stack：获取当前 goroutine 的堆栈信息
```go
/**
buf 是用于存储堆栈信息的字节数组，all 表示是否获取所有 goroutine 的堆栈信息。当 all 为 false 时，只获取当前 goroutine 的堆栈信息。

返回值是实际写入到 buf 中的字节数。如果 buf 太小，无法容纳完整的堆栈信息，函数会返回一个错误
**/
func Stack(buf []byte, all bool) int

//方便地输出所有 goroutine 的堆栈信息
debug.PrintStack()
```