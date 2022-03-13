io 包为 I/O 原语提供了基本的接口。它主要包装了这些原语的已有实现。

由于这些被接口包装的I/O原语是由不同的低级操作实现，因此，在另有声明之前不该假定它们的并发执行是安全的。

## Reader 和 Writer 接口

在 io 包中最重要的是两个接口：Reader 和 Writer 接口。

```gotemplate
// Read 将 len(p) 个字节读取到 p 中。它返回读取的字节数 n（0 <= n <= len(p)） 以及任何遇到的错误。即使 Read 返回的 n < len(p)，它也会在调用过程中占用 len(p) 个字节作为暂存空间。若可读取的数据不到 len(p) 个字节，Read 会返回可用数据，而不是等待更多数据。

// 当 Read 在成功读取 n > 0 个字节后遇到一个错误或 EOF (end-of-file)，它会返回读取的字节数。它可能会同时在本次的调用中返回一个non-nil错误,或在下一次的调用中返回这个错误（且 n 为 0）。 一般情况下, Reader会返回一个非0字节数n, 若 n = len(p) 个字节从输入源的结尾处由 Read 返回，Read可能返回 err == EOF 或者 err == nil。并且之后的 Read() 都应该返回 (n:0, err:EOF)。

调用者在考虑错误之前应当首先处理返回的数据。这样做可以正确地处理在读取一些字节后产生的 I/O 错误，同时允许EOF的出现。
type Reader interface {
	Read(p []byte) (n int, err error)
}

// Write 将 len(p) 个字节从 p 中写入到基本数据流中。它返回从 p 中被写入的字节数 n（0 <= n <= len(p)）以及任何遇到的引起写入提前停止的错误。若 Write 返回的 n < len(p)，它就必须返回一个 非nil 的错误。
type Writer interface {
    Write(p []byte) (n int, err error)
}

// Reader 接口和 Writer 接口的简单组合（内嵌）
// 有些时候同时需要某两个接口的所有功能，即必须同时实现了某两个接口的类型才能够被传入使用。可见，io 包中有大量的“小接口”，这样方便组合为“大接口”
type ReadWriter interface {
    Reader
    Writer
}
```

目前已实现了`io.Reader`或`io.Writer`接口有：

| 包 |  实现`io.Reader` | 实现`io.Writer` |
| --- | --- | --- |
| os.File | √ | √ |
| strings.Reader | √ | - |
| bufio.Reader/Writer | √ | √ |
| bytes.Buffer | √ | √ |
| bytes.Reader | √ | - |
| compress/gzip.Reader/Writer | √ | √ |
| crypto/cipher.StreamReader/StreamWriter | √ | √ |
| crypto/tls.Conn | √ | √ |
| encoding/csv.Reader/Writer | √ | √ |
| mime/multipart.Part | √ | - |
| net/conn | √ | √ |

除此之外，io 包本身也有这两个接口的实现类型。如：
``` 
实现了 Reader 的类型：LimitedReader、PipeReader、SectionReader
实现了 Writer 的类型：PipeWriter
```

## ReaderAt 和 WriterAt 接口

```gotemplate
// ReadAt 从基本输入源的偏移量 off 处开始，将 len(p) 个字节读取到 p 中。它返回读取的字节数 n（0 <= n <= len(p)）以及任何遇到的错误。
//当 ReadAt 返回的 n < len(p) 时，它就会返回一个 非nil 的错误来解释 为什么没有返回更多的字节。在这一点上，ReadAt 比 Read 更严格。
//即使 ReadAt 返回的 n < len(p)，它也会在调用过程中使用 p 的全部作为暂存空间。若可读取的数据不到 len(p) 字节，ReadAt 就会阻塞,直到所有数据都可用或一个错误发生。 在这一点上 ReadAt 不同于 Read。
//若 n = len(p) 个字节从输入源的结尾处由 ReadAt 返回，Read可能返回 err == EOF 或者 err == nil
//若 ReadAt 携带一个偏移量从输入源读取，ReadAt 应当既不影响偏移量也不被它所影响。
//可对相同的输入源并行执行 ReadAt 调用。
type ReaderAt interface {
    ReadAt(p []byte, off int64) (n int, err error)
}

//WriteAt 从 p 中将 len(p) 个字节写入到偏移量 off 处的基本数据流中。它返回从 p 中被写入的字节数 n（0 <= n <= len(p)）以及任何遇到的引起写入提前停止的错误。若 WriteAt 返回的 n < len(p)，它就必须返回一个 非nil 的错误。
//若 WriteAt 携带一个偏移量写入到目标中，WriteAt 应当既不影响偏移量也不被它所影响。
//若被写区域没有重叠，可对相同的目标并行执行 WriteAt 调用。
type WriterAt interface {
    WriteAt(p []byte, off int64) (n int, err error)
}
```

## ReaderFrom 和 WriterTo 接口

```gotemplate
//ReadFrom 从 r 中读取数据，直到 EOF 或发生错误。其返回值 n 为读取的字节数。除 io.EOF 之外，在读取过程中遇到的任何错误也将被返回。
//如果 ReaderFrom 可用，Copy 函数就会使用它。
type ReaderFrom interface {
    ReadFrom(r Reader) (n int64, err error)
}

//WriteTo 将数据写入 w 中，直到没有数据可写或发生错误。其返回值 n 为写入的字节数。 在写入过程中遇到的任何错误也将被返回。
//如果 WriterTo 可用，Copy 函数就会使用它。
type WriterTo interface {
    WriteTo(w Writer) (n int64, err error)
}
```

## Seeker 接口

```gotemplate
//Seek 设置下一次 Read 或 Write 的偏移量为 offset，它的解释取决于 whence： 0 表示相对于文件的起始处，1 表示相对于当前的偏移，而 2 表示相对于其结尾处。 Seek 返回新的偏移量和一个错误，如果有的话。
type Seeker interface {
    Seek(offset int64, whence int) (ret int64, err error)
}
```

## Closer接口

```gotemplate
// 用于关闭数据流
type Closer interface {
    Close() error
}
```
