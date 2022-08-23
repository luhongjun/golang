# 标准库：net.http包

| 数据结构（内部） | 方法 |
| --- | --- |
| http.Client | Get(url string) (resp *Response, err error) <br> Post(url string, contentType string, body io.Reader) (resp *Response, err error) <br> PostForm(url string, data url.Values) (resp *Response, err error) <br> Head(url string) (resp *Response, err error) <br> Do(req *Request) (*Response, error) <br> CloseIdleConnections() <br> |
| http.Cookie | String() string | 
| http.Request | Context() context.Context <br> WithContext(ctx context.Context) *Request <br> Clone(ctx context.Context) *Request  <br> ProtoAtLeast(major int, minor int) bool  <br> UserAgent() string  <br> Cookies() []*Cookie  <br> Cookie(name string) (*Cookie, error)  <br> AddCookie(c *Cookie)  <br> Referer() string  <br> MultipartReader() (*multipart.Reader, error)  <br> Write(w io.Writer) error  <br> WriteProxy(w io.Writer) error  <br> BasicAuth() (username string, password string, ok bool)  <br> SetBasicAuth(username string, password string)  <br> ParseForm() error <br> ParseMultipartForm(maxMemory int64) error  <br> FormValue(key string) string <br> PostFormValue(key string) string <br> FormFile(key string) (multipart.File, *multipart.FileHeader, error) <br> WithT(t *testing.T) *Request <br> ExportIsReplayable() bool |
| http.Response | Cookies() []*Cookie <br> Location() (*url.URL, error) <br> ProtoAtLeast(major int, minor int) bool <br> Write(w io.Writer) error |
| http.Server | Close() error <br> Shutdown(ctx context.Context) error <br> RegisterOnShutdown(f func()) <br> ListenAndServe() error <br> Serve(l net.Listener) error <br> ServeTLS(l net.Listener, certFile string, keyFile string) error <br> SetKeepAlivesEnabled(v bool) <br> ListenAndServeTLS(certFile string, keyFile string) error <br> ExportAllConnsIdle() bool | 
| http.Transport | 略 |

```gotemplate
//客户端
type Client struct {
    Transport     RoundTripper
    CheckRedirect func(req *Request, via []*Request) error
    Jar           CookieJar
    Timeout       time.Duration
}

//Cookie
type Cookie struct {
    Name       string
    Value      string
    Path       string
    Domain     string
    Expires    time.Time
    RawExpires string
    MaxAge     int
    Secure     bool
    HttpOnly   bool
    SameSite   SameSite
    Raw        string
    Unparsed   []string
}

//Request请求
type Request struct {
    Method           string
    URL              *url.URL
    Proto            string
    ProtoMajor       int
    ProtoMinor       int
    Header           Header
    Body             io.ReadCloser
    GetBody          func() (io.ReadCloser, error)
    ContentLength    int64
    TransferEncoding []string
    Close            bool
    Host             string
    Form             url.Values
    PostForm         url.Values
    MultipartForm    *multipart.Form
    Trailer          Header
    RemoteAddr       string
    RequestURI       string
    TLS              *tls.ConnectionState
    Cancel           <-chan struct{}
    Response         *Response
    ctx              context.Context
}

//响应Response
type Response struct {
    Status           string
    StatusCode       int
    Proto            string
    ProtoMajor       int
    ProtoMinor       int
    Header           Header
    Body             io.ReadCloser
    ContentLength    int64
    TransferEncoding []string
    Close            bool
    Uncompressed     bool
    Trailer          Header
    Request          *Request
    TLS              *tls.ConnectionState
}

//服务器Server
type Server struct {
    Addr              string
    Handler           Handler
    TLSConfig         *tls.Config
    ReadTimeout       time.Duration
    ReadHeaderTimeout time.Duration
    WriteTimeout      time.Duration
    IdleTimeout       time.Duration
    MaxHeaderBytes    int
    TLSNextProto      map[string]func(*Server, *tls.Conn, Handler)
    ConnState         func(net.Conn, ConnState)
    ErrorLog          *log.Logger
    BaseContext       func(net.Listener) context.Context
    ConnContext       func(ctx context.Context, c net.Conn) context.Context
    inShutdown        atomicBool
    disableKeepAlives int32
    nextProtoOnce     sync.Once
    nextProtoErr      error
    mu                sync.Mutex
    listeners         map[*net.Listener]struct{}
    activeConn        map[*conn]struct{}
    doneChan          chan struct{}
    onShutdown        []func()
}
```