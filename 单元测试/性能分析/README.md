# Golang 的程序性能分析


## 参考资料

- Go 程序性能分析：https://zhuanlan.zhihu.com/p/235292978

----------------------------

在Go语言中，用于分析程序性能的概要文件有三种：

| 概要文件 | 说明 |
| --- | --- |
| CPU概要文件（CPU Profile） | 记录CPU上正在执行的Go代码 |
| 内存概要文件（Men Profile） | 记录某个采样时刻，正在执行的Go代码以及堆内存的使用情况，这里包括已分配和已释放的字节数量和对象数量 |
| 阻塞概要文件（Block Profile） | 记录着Go程序中的goroutine阻塞事件 |

这些概要文件中包含的都是在某一段时间内，对Go的相关指标进行多次采样后得到的概要信息。


## 采集

### 采集CPU概要信息

| 函数 | 说明 |
| --- | --- |
| runtime/pprof.StartCPUProfile(w io.Writer) error | 当前进程启用CPU配置文件。分析过程中配置文件将被缓冲并写入w。如果已启用分析，StartCPUProfile 函数将返回错误。<br> 在类Unix系统上，默认情况下，对于使用`-buildmode=c-archive`或`-buildmodel=c-shared`构建的Go代码，StartCPUProfile 函数不起作用。<br> StartCPUProfile 依赖于SIGPROF信号，但该信号将传递给主程序的SIGPROF信号处理程序（如果有），而不是Go使用的信号处理程序。要使其工作，请使用系统调用 os/signal.Notify，但是请注意，这样做可能会破坏主程序正在进行的任何分析。 |
| runtime/pprof.StopCPUProfile() | 停止当前CPU配置文件 |
| runtime.SetCPUProfileRate(hz int) | 函数设置CPU采样频率。如果 hz <= 0则会停止CPU采样 <br> 经过大量的实验，Go预研团队发现 100Hz 是一个比较合适的设定，因为操作系统本身对高频采样的处理能力是优先的，一般情况下超过500Hz就很可能得不到及时响应 |

### 采集内存概要信息

针对内存概要信息的采样会按照一定比例收集Go程序在运行期间的堆内存使用情况。设定内存概要信息采样频率的方法很简单，只要为`runtime.MemProfileRate`变量赋值即可。

这个变量的含义是，平均每分配多少个字节，就对堆内存的使用情况进行一次采样

| 函数 | 说明 |
| --- | --- |
| runtime/pprof.WriteHeapProfile(w io.Writer) error | 将内存概要写到缓存区 w 中 |
| runtime.ReadMemStats(m *MemStats) |  使用内存分配器统计信息填充m。此方法可以获取内存的实时信息 |


### 采集阻塞概要信息

| 函数 | 说明 |
| --- | --- |
| runtime.SetBlockProfileRate(rate int) | 控制阻塞配置文件中报告的goroutine阻塞事件的比例。以每纳秒的速率平均采样一个阻塞事件 |
| runtime/pprof.Lookup(name string) *Profile | 提供与给定的名称相对应的概要信息，这个概要信息会由 Profile 结构体来代表。如果没有则返回nil，表示不存在与给定名称对应的概要信息；这个给定的名称可以是：`goroutine、heap、allocs、threadcreate、block、mutex` |

**备注**
```gotemplate
type Profile struct {
	name  string
	mu    sync.Mutex
	m     map[interface{}][]uintptr
	count func() int
	write func(io.Writer, int) error
}

var goroutineProfile = &Profile{
name:  "goroutine",
count: countGoroutine,
write: writeGoroutine,
}

var threadcreateProfile = &Profile{
name:  "threadcreate",
count: countThreadCreate,
write: writeThreadCreate,
}

var heapProfile = &Profile{
name:  "heap",
count: countHeap,
write: writeHeap,
}

var allocsProfile = &Profile{
name:  "allocs",
count: countHeap, // identical to heap profile
write: writeAlloc,
}

var blockProfile = &Profile{
name:  "block",
count: countBlock,
write: writeBlock,
}

var mutexProfile = &Profile{
name:  "mutex",
count: countMutex,
write: writeMutex,
}
```

## 查看并分析

在默认情况下，这些概要文件中的信息并不是普通的文本，它们是通过 protocol buffers 生成的二进制数据流或者字节流。

Go 内置获取程序运行数据的工具，包括以下两个标准库：

- net/http/pprof: 采集服务型应用运行时数据进行分析，详细可见：[基于http协议的网络服务的性能分析](基于http协议的网络服务的性能分析.md)

- runtime/pprof: 采集工具型应用运行数据进行分析

运行指令后，会进入控制台交互模式。可以通过输入相应的命令来输出概要文件的信息：
```text
C:\Users\luhj>go tool pprof http://localhost:8082/debug/pprof/profile?seconds=3
Fetching profile over HTTP from http://localhost:8082/debug/pprof/profile?seconds=3
Saved profile in C:\Users\luhj\pprof\pprof.samples.cpu.001.pb.gz
Type: cpu
Time: Aug 18, 2022 at 8:50pm (CST)
Duration: 3.06s, Total samples = 170ms ( 5.56%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) help
  Commands:
    callgrind        Outputs a graph in callgrind format
    comments         Output all profile comments
    disasm           Output assembly listings annotated with samples
    dot              Outputs a graph in DOT format
    eog              Visualize graph through eog
    evince           Visualize graph through evince
    gif              Outputs a graph image in GIF format
    gv               Visualize graph through gv
    kcachegrind      Visualize report in KCachegrind
    list             Output annotated source for functions matching regexp
    pdf              Outputs a graph in PDF format
    peek             Output callers/callees of functions matching regexp
    png              Outputs a graph image in PNG format
    proto            Outputs the profile in compressed protobuf format
    ps               Outputs a graph in PS format
    raw              Outputs a text representation of the raw profile
    svg              Outputs a graph in SVG format
    tags             Outputs all tags in the profile
    text             Outputs top entries in text form
    top              Outputs top entries in text form
    topproto         Outputs top entries in compressed protobuf format
    traces           Outputs all profile samples in text form
    tree             Outputs a text rendering of call graph
    web              Visualize graph through web browser
    weblist          Display annotated source in a web browser
    o/options        List options and their current values
    quit/exit/^D     Exit pprof

  Options:
    call_tree        Create a context-sensitive call tree
    compact_labels   Show minimal headers
    divide_by        Ratio to divide all samples before visualization
    drop_negative    Ignore negative differences
    edgefraction     Hide edges below <f>*total
    focus            Restricts to samples going through a node matching regexp
    hide             Skips nodes matching regexp
    ignore           Skips paths going through any nodes matching regexp
    mean             Average sample value over first value (count)
    nodecount        Max number of nodes to show
    nodefraction     Hide nodes below <f>*total
    noinlines        Ignore inlines.
    normalize        Scales profile based on the base profile.
    output           Output filename for file-based outputs
    prune_from       Drops any functions below the matched frame.
    relative_percentages Show percentages relative to focused subgraph
    sample_index     Sample value to report (0-based index or name)
    show             Only show nodes matching regexp
    show_from        Drops functions above the highest matched frame.
    source_path      Search path for source files
    tagfocus         Restricts to samples with tags in range or matched by regexp
    taghide          Skip tags matching this regexp
    tagignore        Discard samples with tags in range or matched by regexp
    tagshow          Only consider tags matching this regexp
    trim             Honor nodefraction/edgefraction/nodecount defaults
    trim_path        Path to trim from source paths before search
    unit             Measurement units to display

  Option groups (only set one per group):
    cumulative
      cum              Sort entries based on cumulative weight
      flat             Sort entries based on own weight
    granularity
      addresses        Aggregate at the address level.
      filefunctions    Aggregate at the function level.
      files            Aggregate at the file level.
      functions        Aggregate at the function level.
      lines            Aggregate at the source code line level.
  :   Clear focus/ignore/hide/tagfocus/tagignore

  type "help <cmd|option>" for more information
```

