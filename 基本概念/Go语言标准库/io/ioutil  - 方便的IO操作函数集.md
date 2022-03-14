# ioutil — 方便的IO操作函数集

| 函数 | 签名 | 功能 |
| --- | --- | --- |
| NopCloser 函数 | func NopCloser(r io.Reader) io.ReadCloser | - |
| ReadAll 函数 | func ReadAll(r io.Reader) ([]byte, error) | 从io.Reader 中一次读取所有数据 |
| ReadDir 函数 | func ReadDir(dirname string) ([]os.FileInfo, error) | 读取目录并返回排好序的文件和子目录名 |
| ReadFile 函数 | func ReadFile(filename string) ([]byte, error) | 读取文件内容 |
| WriteFile 函数 | func WriteFile(filename string, data []byte, perm os.FileMode) error | 写入内容到文件中 |
| TempDir 函数 | func TempDir() string | 创建一个自定义的临时目录 |
| TempFile 函数 | func TempFile(dir, pattern string) (f *os.File, err error) | 创建临时文件 |
|
