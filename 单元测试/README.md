# Golang 单元测试

golang 有自带的["testing"库包](https://pkg.go.dev/testing)来提供单元测试的基本方法；编写完测试用例后再通过`go test`来运行用例；

go 语言原生支持了单元测试，使用上非常简单，测试代码只需要放到以 `_test.go` 结尾的文件中即可。

**推荐阅读**
- [Golang 官方包"testing"](https://pkg.go.dev/testing)

## 单元测试的函数

在*_test.go文件中，有三种类型的函数：测试函数、基准测试（benchmark）函数、示例函数

| 函数 | 特点 | 用法 |
| --- | --- | --- |
| 测试函数 | 以Test为函数名前缀（如：`TestXxx`或`Test_xxx`） | 用于测试程序的一些逻辑行为是否正确；go test命令会调用这些测试函数并报告测试结果是PASS或FAIL |
| 基准测试函数 | 以Benchmark为函数名前缀的函数 | 用于衡量一些函数的性能;go test命令会多次运行基准测试函数以计算一个平均的执行时间 |
| 示例函数 | 以Example为函数名前缀的函数 | 提供一个由编译器保证正确性的示例文档 |

- 通用方法


- 测试函数

```gotemplate
func TestSin(t *testing.T) { /* ... */ }
func TestCos(t *testing.T) { /* ... */ }
func TestLog(t *testing.T) { /* ... */ }
```

- 基准测试（benchmark）函数

基准测试是测量一个程序在固定工作负载下的性能。

在Go语言中，基准测试函数和普通测试函数写法类似，但是以Benchmark为前缀名，并且带有一个`*testing.B`类型的参数；`*testing.B`参数除了提供和`*testing.T`类似的方法，还有额外一些和性能测量相关的方法。

它还提供了一个整数N，用于指定操作执行的循环次数;

- 示例函数

```gotemplate
func ExampleIsPalindrome() {
    fmt.Println(IsPalindrome("A man, a plan, a canal: Panama"))
    fmt.Println(IsPalindrome("palindrome"))
    // Output:
    // true
    // false
}
```

1. 根据示例函数的后缀名部分，godoc这个web文档服务器会将示例函数关联到某个具体函数或包本身，因此ExampleIsPalindrome示例函数将是IsPalindrome函数文档的一部分，Example示例函数将是包文档的一部分。
2. 示例函数的第二个用处是，在go test执行测试的时候也会运行示例函数测试。如果示例函数内含有类似上面例子中的// Output:格式的注释，那么测试工具会执行这个示例函数，然后检查示例函数的标准输出与注释是否匹配。
3. 示例函数的第三个目的提供一个真实的演练场。 http://golang.org 就是由godoc提供的文档服务，它使用了Go Playground让用户可以在浏览器中在线编辑和运行每个示例函数，就像图11.4所示的那样。这通常是学习函数使用或Go语言特性最快捷的方式。

