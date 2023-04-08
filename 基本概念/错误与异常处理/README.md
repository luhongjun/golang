# Golang错误与异常处理

## 错误

处理error的库有：
- `github.com/pkg/errors`
    - `func As(err error, target interface{}) bool`：判断err调用链中是否能匹配target
    - `func Cause(err error) error`：找出err调用链中最早的err（需继承cause接口）
    - `func Errorf(format string, args ...interface{}) error`
    - `func Is(err, target error) bool`
    - `func New(message string) error`
    - `func Unwrap(err error) error`
    - `func WithMessage(err error, message string) error`
    - `func WithMessagef(err error, format string, args ...interface{}) error`
    - `func WithStack(err error) error`
    - `func Wrap(err error, message string) error`
    - `func Wrapf(err error, format string, args ...interface{}) error`