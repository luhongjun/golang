# Golang汇编语言

## 相关资料

- [golang汇编语言基础](https://blog.csdn.net/u010853261/article/details/101060546)

在深入阅读runtime和标准库的源码时候，发现底层有大片代码都会与汇编打交道。


本文涉及到计算机架构体系相关的情况时，请假设我们是运行在 linux/amd64 平台上。

## 伪汇编

Go 编译器会输出一种抽象可移植的汇编代码，这种汇编并不对应某种真实的硬件架构。之后 Go 的汇编器使用这种伪汇编，为目标硬件生成具体的机器指令。

伪汇编这一个额外层可以带来很多好处，最主要的一点是方便将 Go 移植到新的架构上。相关的信息可以参考文后列出的 Rob Pike 的 The Design of the Go Assembler。

以下是Golang汇编的示例：
```gotemplate
package  main
//go:noinline
func add(a, b int32) (int32, bool) {
    return a + b, true
}

func main() {
    add(10, 32)
}
```

