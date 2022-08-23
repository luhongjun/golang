# 标准库：net.smtp包

| 数据结构（内部） | 方法 |
| --- | --- |
| net.smtp.Client | Close() error <br> Hello(localName string) error <br> StartTLS(config *tls.Config) error <br> TLSConnectionState() (state tls.ConnectionState, ok bool) <br> Verify(addr string) error <br> Auth(a Auth) error <br> Mail(from string) error <br> Rcpt(to string) error <br> Data() (io.WriteCloser, error) <br> Extension(ext string) (bool, string) <br> Reset() error <br> Noop() error <br> Quit() error |


```gotemplate
//A Client represents a client connection to an SMTP server.
type Client struct {
    Text       *textproto.Conn
    conn       net.Conn
    tls        bool
    serverName string
    ext        map[string]string
    auth       []string
    localName  string
    didHello   bool
    helloError error
}
```

