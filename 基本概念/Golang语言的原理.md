# Golang语言的原理

Go（又称Golang）是Google开发的一种静态强类型、编译型、并发型，并具有垃圾回收功能的编程语言

## 资料
[百度百科 - Golang](https://baike.baidu.com/item/go/953521?fr=aladdin)
[Github - golang/go包](https://github.com/golang/go)

## 为什么选择Golang

Golang 也应当属第二象限。它是类型安全的强类型，同时又是具有弱类型声明机制的静态类型。这些特性使 Golang 的编译器、IDE 拥有完善的代码分析和理解能力，编译过程就能暴露出大部分潜在的逻辑性错误，适用于大规模团队协作开发复杂庞大的分布式服务器端应用系统，具有简单高效，风格统一，性能突出的最佳实践组合

## Golang安装目录结构

- `/bin`：包含可执行文件，如：编译器，Go 工具
- `/doc`：包含示例程序，代码工具，本地文档等
- `/lib`：包含文档模版
- `/misc`：包含与支持 Go 编辑器有关的配置文件以及 cgo 的示例
- `/os_arch`：包含标准库的包的对象文件（.a）
- `/src`：包含源代码构建脚本和标准库的包的完整源代码（Go 是一门开源语言）
- `/src/cmd`：包含 Go 和 C 的编译器和命令行脚本

## Golang编译原理

- [Go语言是怎么完成编译的](https://www.cnblogs.com/maomaomaoge/p/14178277.html#%E7%BC%96%E8%AF%91%E5%8E%9F%E7%90%86)

## Golang垃圾回收

Go的GC自打出生的时候就开始被人诟病，但是在引入v1.5的三色标记和v1.8的混合写屏障后，正常的GC已经缩短到10us左右

- [Go GC垃圾回收机制](https://studygolang.com/articles/18281)
- [深入理解 Go垃圾回收](https://segmentfault.com/a/1190000020086769)
