# io函数

## Copy 和 CopyN 函数

```gotemplate
// Copy 将 src 复制到 dst，直到在 src 上到达 EOF 或发生错误。它返回复制的字节数，如果有错误的话，还会返回在复制时遇到的第一个错误。
// 成功的 Copy 返回 err == nil，而非 err == EOF。由于 Copy 被定义为从 src 读取直到 EOF 为止，因此它不会将来自 Read 的 EOF 当做错误来报告。
// 若 dst 实现了 ReaderFrom 接口，其复制操作可通过调用 dst.ReadFrom(src) 实现。此外，若 src 实现了 WriterTo 接口，其复制操作可通过调用 src.WriteTo(dst) 实现。

func Copy(dst Writer, src Reader) (written int64, err error)



// CopyN 将 n 个字节(或到一个error)从 src 复制到 dst。 它返回复制的字节数以及在复制时遇到的最早的错误。当且仅当err == nil时,written == n 。
// 若 dst 实现了 ReaderFrom 接口，复制操作也就会使用它来实现。

func CopyN(dst Writer, src Reader, n int64) (written int64, err error)
```

## ReadAtLeast 和 ReadFull 函数

ReadAtLeast 函数签名：
```gotemplate
// ReadAtLeast 将 r 读取到 buf 中，直到读了最少 min 个字节为止。它返回复制的字节数，如果读取的字节较少，还会返回一个错误。若没有读取到字节，错误就只是 EOF。如果一个 EOF 发生在读取了少于 min 个字节之后，ReadAtLeast 就会返回 ErrUnexpectedEOF。若 min 大于 buf 的长度，ReadAtLeast 就会返回 ErrShortBuffer。对于返回值，当且仅当 err == nil 时，才有 n >= min。
func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error)
```

ReadFull 函数签名：
```gotemplate
// ReadFull 精确地从 r 中将 len(buf) 个字节读取到 buf 中。它返回复制的字节数，如果读取的字节较少，还会返回一个错误。若没有读取到字节，错误就只是 EOF。如果一个 EOF 发生在读取了一些但不是所有的字节后，ReadFull 就会返回 ErrUnexpectedEOF。对于返回值，当且仅当 err == nil 时，才有 n == len(buf)。
func ReadFull(r Reader, buf []byte) (n int, err error)
```

## WriteString 函数

```gotemplate
// WriteString 将s的内容写入w中，当 w 实现了 WriteString 方法时，会直接调用该方法，否则执行 w.Write([]byte(s))。
func WriteString(w Writer, s string) (n int, err error)
```

## MultiReader 和 MultiWriter 函数

合并操作

```gotemplate
// 它们接收多个 Reader 或 Writer，返回一个 Reader 或 Writer。我们可以猜想到这两个函数就是操作多个 Reader 或 Writer 就像操作一个。
func MultiReader(readers ...Reader) Reader
func MultiWriter(writers ...Writer) Writer
```

示例：
```gotemplate
file, err := os.Create("tmp.txt")
if err != nil {
    panic(err)
}
defer file.Close()
writers := []io.Writer{
    file,
    os.Stdout,
}
writer := io.MultiWriter(writers...)
writer.Write([]byte("Go语言中文网"))
```


## TeeReader函数

```gotemplate
// TeeReader 返回一个 Reader，它将从 r 中读到的数据写入 w 中。所有经由它处理的从 r 的读取都匹配于对应的对 w 的写入。它没有内部缓存，即写入必须在读取完成前完成。任何在写入时遇到的错误都将作为读取错误返回。
func TeeReader(r Reader, w Writer) Reader
```

也就是说，我们通过 Reader 读取内容后，会自动写入到 Writer 中去。例子代码如下：
```gotemplate
reader := io.TeeReader(strings.NewReader("Go语言中文网"), os.Stdout)
reader.Read(make([]byte, 20))
```