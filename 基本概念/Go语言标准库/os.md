# OS库

## 常量
```gotemplate
// Flags to OpenFile wrapping those of the underlying system. Not all
// flags may be implemented on a given system.
const (
	// Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
	O_RDONLY int = syscall.O_RDONLY // open the file read-only.
	O_WRONLY int = syscall.O_WRONLY // open the file write-only.
	O_RDWR   int = syscall.O_RDWR   // open the file read-write.
	// The remaining values may be or'ed in to control behavior.
	O_APPEND int = syscall.O_APPEND // append data to the file when writing.
	O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
	O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
	O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
	O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.
)
```

## 变量
```gotemplate
// Portable analogs of some common system call errors.
//
// Errors returned from this package may be tested against these errors
// with errors.Is.
var (
	// ErrInvalid indicates an invalid argument.
	// Methods on File will return this error when the receiver is nil.
	ErrInvalid = fs.ErrInvalid // "invalid argument"

	ErrPermission = fs.ErrPermission // "permission denied"
	ErrExist      = fs.ErrExist      // "file already exists"
	ErrNotExist   = fs.ErrNotExist   // "file does not exist"
	ErrClosed     = fs.ErrClosed     // "file already closed"

	ErrNoDeadline       = errNoDeadline()       // "file type does not support deadline"
	ErrDeadlineExceeded = errDeadlineExceeded() // "i/o timeout"
)
```

## 函数

| 函数                                                                  | 说明                                         |
|---------------------------------------------------------------------|--------------------------------------------|
| `func Chdir(dir string) error`                                      | 修改当前目录                                     |
| `func Chmod(name string, mode FileMode) error`                      | 修改文件访问权限                                   |
| `func Chown(name string, uid, gid int) error`                       | 修改文件的用户、用户组                                |
| `func Chtimes(name string, atime time.Time, mtime time.Time) error` | 修改文件的访问时间和改动时间                             |
| `func Clearenv()`                                                   | 清理所有环境变量                                   |
| `func DirFS(dir string) fs.FS`                                      | 返回基于某个目录路径的文件系统                            |
| `func Environ() []string`                                           | 返回格式为"key=value"的环境变量                      |
| `func Executable() (string, error)`                                 | 返回启动当前进程的可执行文件的路径名                         |
| `func Exit(code int)` | 退出当前进程                                     |
| `func Expand(s string, mapping func(string) string) string` | 基于回调函数将 s 字符串含 $var 或者 ${var} 格式的变量替换为指定的值 |
| `func Getegid() int` | 获取调用者的有效的组ID                               |
| `func Getenv(key string) string` | 获取环境变量                                     |
| `func Geteuid() int` | 获取调用者的用户ID                                 |
| `func Getgid() int` | 获取调用者的组ID |
| 