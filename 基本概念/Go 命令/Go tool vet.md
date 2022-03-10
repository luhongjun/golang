# go tool vet

## 参考资料

- [go vet与go tool vet](https://www.bookstack.cn/read/go_command_tutorial/0.11.md)

## 概要

go vet是一个用于检查Go语言源码中静态错误的简单工具。

go vet命令是go tool vet命令的简单封装。它会首先载入和分析指定的代码包，并把指定代码包中的所有Go语言源码文件和以“.s”结尾的文件的相对路径作为参数传递给go tool vet命令。其中，以“.s”结尾的文件是汇编语言的源码文件。如果go vet命令的参数是Go语言源码文件的路径，则会直接将这些参数传递给go tool vet命令。

如果我们直接使用go tool vet命令，则其参数可以传递任意目录的路径，或者任何Go语言源码文件和汇编语言源码文件的路径。路径可以是绝对的也可以是相对的。

## 指令说明

| 指令参数 | 概要 |
| --- | --- |
| -all | 进行全部检查。如果有其他检查标记被设置，则命令程序会将此值变为false。默认值为true |
| -asmdecl | 对汇编语言的源码文件进行检查。默认值为false |
| -assign | 检查赋值语句。默认值为false | 
| -atomic | 检查代码中对代码包sync/atomic的使用是否正确。默认值为false |
| -buildtags | 检查编译标签的有效性。默认值为false |
| -composites | 检查复合结构实例的初始化代码。默认值为false |
| -compositeWhiteList | 是否使用复合结构检查的白名单。仅供测试使用。默认值为true |
| -methods | 检查那些拥有标准命名的方法的签名。默认值为false |
| -printf | 检查代码中对打印函数的使用是否正确。默认值为false |
| -printfuncs | 需要检查的代码中使用的打印函数的名称的列表，多个函数名称之间用英文半角逗号分隔。默认值为空字符串 |
| -rangeloops | 检查代码中对在range语句块中迭代赋值的变量的使用是否正确。默认值为false |
| -structtags | 检查结构体类型的字段的标签的格式是否标准。默认值为false |
| -unreachable | 查找并报告不可到达的代码。默认值为false |