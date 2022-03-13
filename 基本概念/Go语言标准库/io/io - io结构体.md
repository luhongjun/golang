## io结构体

## SectionReader 类型

```gotemplate
// 读取数据流中部分数据
type SectionReader struct {
    r     ReaderAt    // 该类型最终的 Read/ReadAt 最终都是通过 r 的 ReadAt 实现
    base  int64        // NewSectionReader 会将 base 设置为 off
    off   int64        // 从 r 中的 off 偏移处开始读取数据
    limit int64        // limit - off = SectionReader 流的长度
}
```

## LimitedReader 类型

```gotemplate
//从 R 读取但将返回的数据量限制为 N 字节。每调用一次 Read 都将更新 N 来反应新的剩余数量。
type LimitedReader struct {
    R Reader // underlying reader，最终的读取操作通过 R.Read 完成
    N int64  // max bytes remaining
}
```

## PipeReader 和 PipeWriter 类型

io.Pipe() 用于创建一个同步的内存管道 (synchronous in-memory pipe)，函数签名：
```gotemplate
func Pipe() (*PipeReader, *PipeWriter) {
	p := &pipe{
		wrCh: make(chan []byte),
		rdCh: make(chan int),
		done: make(chan struct{}),
	}
	return &PipeReader{p}, &PipeWriter{p}
}
```

**说明**
它将 io.Reader 连接到 io.Writer。一端的读取匹配另一端的写入，直接在这两端之间复制数据；它没有内部缓存。它对于并行调用 Read 和 Write 以及其它函数或 Close 来说都是安全的。一旦等待的 I/O 结束，Close 就会完成。并行调用 Read 或并行调用 Write 也同样安全：同种类的调用将按顺序进行控制。

正因为是同步的，因此不能在一个 goroutine 中进行读和写。

--------------

PipeReader（一个没有任何导出字段的 struct）是管道的读取端。它实现了 io.Reader 和 io.Closer 接口。结构定义如下：

```gotemplate
type PipeReader struct {
    p *pipe
}

// 从管道中读取数据。该方法会堵塞，直到管道写入端开始写入数据或写入端被关闭。如果写入端关闭时带有 error（即调用 CloseWithError 关闭），该Read返回的 err 就是写入端传递的error；否则 err 为 EOF。
func (r *PipeReader) Read(data []byte) (n int, err error) {
    return r.p.Read(data)
}
```

PipeWriter（一个没有任何导出字段的 struct）是管道的写入端。它实现了 io.Writer 和 io.Closer 接口。结构定义如下：
```gotemplate
type PipeWriter struct {
    p *pipe
}

// 写数据到管道中。该方法会堵塞，直到管道读取端读完所有数据或读取端被关闭。如果读取端关闭时带有 error（即调用 CloseWithError 关闭），该Write返回的 err 就是读取端传递的error；否则 err 为 ErrClosedPipe。
func (w *PipeWriter) Write(data []byte) (n int, err error) {
    return w.p.Write(data)
}
```