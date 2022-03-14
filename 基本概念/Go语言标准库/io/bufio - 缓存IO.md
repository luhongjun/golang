# bufio - 缓存IO

bufio 包实现了缓存IO。它包装了 io.Reader 和 io.Writer 对象，创建了另外的Reader和Writer对象，它们也实现了 io.Reader 和 io.Writer 接口，不过它们是有缓存的。该包同时为文本I/O提供了一些便利操作。

