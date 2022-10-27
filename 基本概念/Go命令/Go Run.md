## go run

## 参考资料

- [go run命令——编译并运行](http://c.biancheng.net/view/121.html)

## 概要

Python 或者 Lua 语言可以在不输出二进制的情况下，将代码使用虚拟机直接执行。Go语言虽然不使用虚拟机，但可使用go run指令达到同样的效果

go run命令会编译源码，并且直接执行源码的 main() 函数，不会在当前目录留下可执行文件。

go run不会在运行目录下生成任何文件，可执行文件被放在临时文件中被执行，工作目录被设置为当前目录。

在go run的后部可以添加参数，这部分参数会作为代码可以接受的命令行输入提供给程序。
```gotemplate
# go run . arg1 arg2 arg3
```
这样的话，是可以在程序中接收到的：
```gotemplate
import "os"
func main() {
    fmt.Println("args:", os.Args) // "os.Args" of Type : string[]
}
```

