# testing.common 公共类

我们知道单元测试函数需要传递一个testing.T类型的参数，而性能测试函数需要传递一个testing.B类型的参数，该参数可用于控制测试的流程，比如标记测试失败等。

testing.T和testing.B属于testing包中的两个数据类型，该类型提供一系列的方法用于控制函数执行流程，考虑到二者有一定的相似性，所以Go实现时抽象出一个`testing.common`作为一个基础类型，而testing.T和testing.B则属于testing.common的扩展。

testing.common 的数据结构：

```go
package testing

// common holds the elements common between T and B and
// captures common methods such as Errorf.
type common struct {
	mu      sync.RWMutex        //读写锁，仅用于控制本数据内的成员访问
	output  []byte              //存储当前测试产生的日志，每产生一条日志则追加到该切片中，待测试结束后再一并输出
	w       io.Writer           //子测试执行结束需要把产生的日志输送到父测试中的output切片中，传递时需要考虑缩进等格式调整，通过w把日志传递到父测试
	ran     bool                // 仅表示是否已执行过。比如，跟据某个规范筛选测试，如果没有测试被匹配到的话，则common.ran为false，表示没有测试运行过
	failed  bool                //如果当前测试执行失败，则置为true
	skipped bool                //标记当前测试是否已跳过
	done    bool                //表示当前测试及其子测试已结束，此状态下再执行Fail()之类的方法标记测试状态会产生panic
	helpers map[string]struct{} //标记当前为函数为help函数，其中打印的日志，在记录日志时不会显示其文件名及行号

	chatty     bool   // 对应命令行中的-v参数，默认为false，true则打印更多详细日志
	finished   bool   // 如果当前测试结束，则置为true
	hasSub     int32  // 标记当前测试是否包含子测试，当测试使用t.Run()方法启动子测试时，t.hasSub则置为1
	raceErrors int    // 竞态检测错误数
	runner     string // 执行当前测试的函数名

	parent   *common       // 如果当前测试为子测试，则置为父测试的指针
	level    int           // 测试嵌套层数，比如创建子测试时，子测试嵌套层数就会加1
	creator  []uintptr     // 测试函数调用栈
	name     string        // 记录每个测试函数名，比如测试函数TestAdd(t *testing.T), 其中t.name即“TestAdd”。 测试结束，打印测试结果会用到该成员
	start    time.Time     // 记录测试开始的时间
	duration time.Duration // 记录测试所花费的时间
	barrier  chan bool     // 用于控制父测试和子测试执行的channel，如果测试为Parallel，则会阻塞等待父测试结束后再继续
	signal   chan bool     // 通知当前测试结束
	sub      []*T          // 子测试列表
}

// Name 返回common结构体中存储的名称
func (c *common) Name() string {
	return c.name
}

// Fail 标记当前测试为失败，然后继续运行，并不会立即退出当前测试。如果是子测试，则除了标记当前测试结果外还通过c.parent.Fail()来标记父测试失败
func (c *common) Fail() {
	if c.parent != nil {
		c.parent.Fail()
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	// c.done needs to be locked to synchronize checks to c.done in parent tests.
	if c.done {
		panic("Fail in goroutine after " + c.name + " has completed")
	}
	c.failed = true
}

// FailNow 内部会调用Fail()标记测试失败，还会标记测试结束并退出当前测试协程。 可以简单的把一个测试理解为一个协程，FailNow()只会退出当前协程，并不会影响其他测试协程，但要保证在当前测试协程中调用FailNow()才有效，不可以在当前测试创建的协程中调用该方法
func (c *common) FailNow() {
	c.Fail()
	c.finished = true
	runtime.Goexit()
}

// 内部记录日志入口，日志会统一记录到common.output切片中，测试结束时再统一打印出来。 日志记录时会调用common.decorate()进行装饰，即加上文件名和行号，还会做一些其他格式化处理。 调用common.log()的方法，有Log()、Logf()、Error()、Errorf()、Fatal()、Fatalf()、Skip()、Skipf()等
func (c *common) log(s string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.output = append(c.output, c.decorate(s)...)
}

// Helper marks the calling function as a test helper function.
// When printing file and line information, that function will be skipped.
// Helper may be called simultaneously from multiple goroutines.
// 标记当前函数为help函数，所谓help函数，即其中打印的日志，不记录help函数的函数名及行号，而是记录上一层函数的函数名和行号
func (c *common) Helper() {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.helpers == nil {
		c.helpers = make(map[string]struct{})
	}
	c.helpers[callerName(1)] = struct{}{}
}




```


