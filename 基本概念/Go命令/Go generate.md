# Go Generate 

## 参考资料

- [简书go generate介绍](https://www.jianshu.com/p/a866147021da)
- [go blog-generate](https://blog.golang.org/generate)
- [go generate 文档](https://docs.google.com/document/d/1V03LUfjSADDooDMhe-_K59EgpTEm3V8uvQRuNMAEnjg/edit#heading=h.j6dsjy94dn2q)
- [深入理解Go之generate](https://studygolang.com/articles/22984?fr=sidebar)

## 概要

go generate命令是go 1.4版本里面新添加的一个命令，当运行go generate时，它将扫描与当前包相关的源代码文件，找出所有包含"//go:generate"的特殊注释，提取并执行该特殊注释后面的命令，命令为可执行程序，形同shell下面执行。

**注意**
- 特殊注释必须在.go文件内
- 每个源码文件可以包含多个generate特殊注释
- 运行go generate命令时，才会执行特殊注释后面的命令
- go generate是串行执行的，如果出错，就终止后面的执行；
-注释必须以"//go:generate"开头，双斜线后面没有空格

## 命令说明

| 命令参数 | 概要 |
| --- | --- |
| -x | 打印出执行的具体命令，并执行 |
| -v | 打印执行的包和文件 |
| -n | 打印出执行的具体命令，但不执行！ |


在执行`go generate`时也可以使用如下环境变量：

| 变量名 | 含义 |
| --- | --- |
| $GOARCH | 体系架构（arm、amd64 等） |
| $GOOS | 当前的 OS 环境（linux、windows 等） |
| $GOFILE | 当前处理中的文件名 |
| $GOLINE | 当前命令在文件中的行号 |
| $GOPACKAGE | 当前处理文件的包名 |


## 应用场景

比如我们可以在我们项目的.go文件这样写：
```gotemplate
//go:generate go mod edit -require=github.com/gin-gonic/gin@v1.7.7

package main

// ...
```
我们在当前目录执行`go generate`后，就是执行`go mod edit`的操作

- yacc：从 .y 文件生成 .go 文件
  
```gotemplate
//go:generate -command yacc go tool yacc
//go:generate yacc -o foo.go foo.y
```

- protobufs：从 protocol buffer 定义文件（.proto）生成 .pb.go 文件
  
```gotemplate
//go:generate protoc -go_out=. file.proto
//go:generate protoc -go_out=. file1.proto file2.proto
```
- Unicode：从 UnicodeData.txt 生成 Unicode 表。
- HTML：将 HTML 文件嵌入到 go 源码 。
- bindata：将形如 JPEG 这样的文件转成 go 代码中的字节数组

```gotemplate
//go:generate bindata -o jpegs.go pic1.jpg pic2.jpg pic3.jpg
```

## 自动化

我们可以在构建项目的脚本中执行这个命令，例如在 makefile 中：
```
go generate && go build .
 ```