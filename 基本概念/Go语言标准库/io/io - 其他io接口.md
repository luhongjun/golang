## ByteReader 和 ByteWriter

```gotemplate

type ByteReader interface {
    ReadByte() (c byte, err error)
}


type ByteWriter interface {
    WriteByte(c byte) error
}
```

## ByteScanner、RuneReader 和 RuneScanner

```gotemplate
type ByteScanner interface {
    ByteReader
    // 将上一次 ReadByte 的字节还原，使得再次调用 ReadByte 返回的结果和上一次调用相同，也就是说，UnreadByte 是重置上一次的 ReadByte
    UnreadByte() error
}

type RuneReader interface {
    //读取一个UTF-8编码的字符串
    ReadRune() (r rune, size int, err error)
}

type RuneScanner interface {
    RuneReader
    UnreadRune() error
}
```

## ReadCloser、ReadSeeker、ReadWriteCloser、ReadWriteSeeker、ReadWriter、WriteCloser 和 WriteSeeker 接口

这些接口都是io接口的组合。

这些接口的作用是：有些时候同时需要某两个接口的所有功能，即必须同时实现了某两个接口的类型才能够被传入使用。可见，io 包中有大量的“小接口”，这样方便组合为“大接口”。
