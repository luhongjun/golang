# go tool cgo

## 参考资料

- [Go语言学习之cgo(golang与C语言相互调用)](https://studygolang.com/articles/10741?fr=sidebar)
- [Go与C语言的互操作](https://tonybai.com/2012/09/26/interoperability-between-go-and-c/)
- [go官方文档-cgo](https://pkg.go.dev/cmd/cgo)

## 概要

几乎所有的编程语言都有C语言的影子，当然golang也不例外。可以看到golang的创始者们与c language有着密切的联系。所有，golang和c语言的相互调用也是理所应当。

那么什么场景需要用到这个工具来将Golang语言转换为C呢？

- 提升局部代码性能时，用C替换一些Go代码；
- Go内存GC性能不足，需要自己手动管理应用内存；
- 实现一些库的Go Wrapper。比如Oracle提供的C版本OCI，但Oracle并未提供Go版本的以及连接DB的协议细节，因此只能通过包装C OCI版本的方式以提供Go开发者使用

但是，虽然Go提供了强大的与C互操作的功能，但目前依旧不完善，比如不支持在Go中直接调用可变个数参数的函数(issue975)，如printf(因此，文档中多用fputs)。


## 详细介绍

下面详细介绍Go调用C的原理。首先，我们编写 example.go 文件，内容如下：
```go
package main

// #include <stdio.h>
// #include <stdlib.h>
/*
void print(char *str) {
    printf("%s\n", str);
}
*/
import "C"

import "unsafe"

func main() {
	s := "Hello Cgo"
	cs := C.CString(s)
	C.print(cs)
	C.free(unsafe.Pointer(cs))
}
```

与正常的Go代码相比，上述代码有几处特殊的地方：

1. 开头的注释中出现了C头文件的include字样
2. 在注释中定义了C函数print
3. import的一个名为C的"包"
4. 在main函数中居然调用了上述的那个C函数-print

这就是在Go源码中调用C代码的步骤，可以看出我们可直接在Go源码文件中编写C代码

首先，Go源码文件中的C代码是需要用注释包裹的，就像上面的include 头文件以及print函数定义；

其次，import "C"这个语句是必须的，而且其与上面的C代码之间不能用空行分隔，必须紧密相连。这里的"C"不是包名，而是一种类似名字空间的概念，或可以理解为伪包，C语言所有语法元素均在该伪包下面；

最后，访问C语法元素时都要在其前面加上伪包前缀，比如C.uint和上面代码中的C.print、C.free等。

------------------

那么我们如何来编译这个go源文件呢？

其实与正常Go源文件没啥区别，依旧可以直接通过`go build`或`go run`来编译和执行。

**原理**
但实际编译过程中，go调用了名为cgo的工具，cgo会识别和读取Go源文件中的C元素，并将其提取后交给C编译器编译，最后与Go源码编译后的目标文件链接成一个可执行程序。

这样我们就不难理解为何Go源文件中的C代码要用注释包裹了，这些特殊的语法都是可以被Cgo识别并使用的。

## 指令说明

cgo工具还支持将Go文件转换为Go和C的源文件

用法：`go tool cgo [cgo options] [-- compiler options] gofiles...`

| 名称 |	默认值 | 说明 |
| --- | --- | --- |
| -cdefs | false | 将改写后的源码内容以C定义模式打印到标准输出，而不生成相关的源码文件 |
| -godefs|false | 将改写后的源码内容以Go定义模式打印到标准输出，而不生成相关的源码文件 |
| -objdir | "" | gcc编译的目标文件所在的路径。若未自定义则为当前目录下的_obj子目录 |
| -dynimport | "" | 如果值不为空字符串，则打印为其值所代表的文件生成的动态导入数据到标准输出 |
| -dynlinker | false | 记录在dynimport模式下的动态链接器信息 |
| -dynout | "" | 将-dynimport的输出（如果有的话）写入到其值所代表的文件中 |
| -gccgo | false | 生成可供gccgo编译器使用的文件 |
| -gccgopkgpath | "" | 对应于gccgo编译器的-fgo-pkgpath选项 |
| -gccgoprefix | "" | 对应于gccgo编译器的-fgo-prefix选项 |
| -debug-define | false | 打印相关的指令符#defines及其后续内容到标准输出 |
| -debug-gcc | false | 打印gcc调用信息到标准输出 |
| -import_runtime_cgo | true | 在生成的代码中加入语句“import runtime/cgo” |
| -import_syscall | true | 在生成的代码中加入语句“import syscall” |
