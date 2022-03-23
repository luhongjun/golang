# Golang汇编 - 变量

为了简单，我们先用Go语言定义并赋值一个整数变量，然后查看生成的汇编代码。

首先创建一个pkg.go文件，内容如下：
```go
package pkg

var Id = 9527
```

代码中只定义了一个int类型的包级变量，并进行了初始化。然后用以下命令查看的Go语言程序对应的伪汇编代码：
```
# go tool compile -S example.go
go.cuinfo.packagename. SDWARFCUINFO dupok size=0
        0x0000 68                                               h
"".Name SNOPTRDATA size=8
        0x0000 37 25 00 00 00 00 00 00                          7%......
```

其中`go tool compile`命令用于调用Go语言提供的底层命令工具，其中`-S`参数表示输出汇编格式。
输出的汇编比较简单，其中`"".Name`对应Name变量符号，变量的内存大小为8个字节。变量的初始化内容为`37 25 00 00 00 00 00 00`，对应十六进制格式的`0x2537`，对应十进制为`9527`。`SNOPTRDATA`是相关的标志，其中`NOPTR`表示数据中不包含指针数据。

## 定义整数变量


- 变量定义

Go汇编语言提供了DATA命令用于初始化包变量，DATA命令的语法如下：
``` 
DATA symbol+offset(SB)/width, value
```

其中:
- `symbol`为变量在汇编语言中对应的标识符
- `offset`是符号开始地址的偏移量
- `width`是要初始化内存的宽度大小 
- `value`是要初始化的值。
  
其中当前包中Go语言定义的符号symbol，在汇编代码中对应·symbol，其中“·”中点符号为一个特殊的unicode符号。

我们采用以下命令可以给Name变量初始化为十六进制的0x2537，对应十进制的9527（常量需要以美元符号$开头表示）:
``` 
DATA ·Name+0(SB)/1,$0x37
DATA ·Name+1(SB)/1,$0x25
```

------------------------

变量定义好之后需要导出以供其它代码引用。Go汇编语言提供了GLOBL命令用于将符号导出：
``` 
GLOBL symbol(SB), width
```
其中symbol对应汇编中符号的名字，width为符号对应内存的大小。用以下命令将汇编中的·Id变量导出：
``` 
GLOBL ·Name, $8
```

我们将完整的汇编代码放到pkg_amd64.s文件中：
``` 
GLOBL ·Name(SB),$8

DATA ·Name+0(SB)/1,$0x37
DATA ·Name+1(SB)/1,$0x25
DATA ·Name+2(SB)/1,$0x00
DATA ·Name+3(SB)/1,$0x00
DATA ·Name+4(SB)/1,$0x00
DATA ·Name+5(SB)/1,$0x00
DATA ·Name+6(SB)/1,$0x00
DATA ·Name+7(SB)/1,$0x00
```

文件名pkg_amd64.s的后缀名表示AMD64环境下的汇编代码文件。

虽然pkg包是用汇编实现，但是用法和之前的Go语言版本完全一样：
``` 
package main

import pkg "pkg包的路径"

func main() {
    println(pkg.Name)
}
```

## 定义字符串变量

还是按照上面方式，创建文件：
```go
package pkg

var Name = "gopher"
```

然后用以下命令查看的Go语言程序对应的伪汇编代码：
``` 
go.cuinfo.packagename. SDWARFCUINFO dupok size=0
        0x0000 68                                               h
go.string."gopher" SRODATA dupok size=6
        0x0000 67 6f 70 68 65 72                                gopher
"".Name SDATA size=16
        0x0000 00 00 00 00 00 00 00 00 06 00 00 00 00 00 00 00  ................
        rel 0+8 t=1 go.string."gopher"+0
```
输出中出现了一个新的符号`go.string."gopher"`，根据其长度和内容分析可以猜测是对应底层的"gopher"字符串数据。

因为Go语言的字符串并不是值类型，Go字符串其实是一种只读的引用类型。如果多个代码中出现了相同的"gopher"只读字符串时，程序链接后可以引用的同一个符号go.string."gopher"。因此，该符号有一个`SRODATA`标志表示这个数据在只读内存段，`dupok`表示出现多个相同标识符的数据时只保留一个就可以了。

而真正的Go字符串变量Name对应的大小却只有16个字节了。其实Name变量并没有直接对应“gopher”字符串，而是对应16字节大小的reflect.StringHeader结构体：
```gotemplate
type reflect.StringHeader struct {
    Data uintptr
    Len  int
}
```

从汇编角度看，Name变量其实对应的是reflect.StringHeader结构体类型。前8个字节对应底层真实字符串数据的指针，也就是符号go.string."gopher"对应的地址。后8个字节对应底层真实字符串数据的有效长度，这里是6个字节。

-------------------------

现在创建pkg_amd64.s文件，尝试通过汇编代码重新定义并初始化Name字符串：
``` 
GLOBL ·NameData(SB),$8
DATA  ·NameData(SB)/8,$"gopher"

GLOBL ·Name(SB),$16
DATA  ·Name+0(SB)/8,$·NameData(SB)
DATA  ·Name+8(SB)/8,$6
```

因为在Go汇编语言中，go.string."gopher"不是一个合法的符号，因此我们无法通过手工创建（这是给编译器保留的部分特权，因为手工创建类似符号可能打破编译器输出代码的某些规则）。因此我们新创建了一个`·NameData`符号表示底层的字符串数据。然后定义·Name符号内存大小为16字节，其中前8个字节用·NameData符号对应的地址初始化，后8个字节为常量6表示字符串长度。

当用汇编定义好字符串变量并导出之后，还需要在Go语言中声明该字符串变量。然后就可以用Go语言代码测试Name变量了：
```go
package main

import pkg "path/to/pkg"

func main() {
    println(pkg.Name)
}
```

不幸的是这次运行产生了以下错误：
``` 
pkgpath.NameData: missing Go //type information for global symbol: size 8
```
错误提示汇编中定义的NameData符号没有类型信息。其实Go汇编语言中定义的数据并没有所谓的类型，每个符号只不过是对应一块内存而已，因此NameData符号也是没有类型的。

但是Go语言是再带垃圾回收器的语言，而Go汇编语言是工作在自动垃圾回收体系框架内的。当Go语言的垃圾回收器在扫描到NameData变量的时候，无法知晓该变量内部是否包含指针，因此就出现了这种错误。错误的根本原因并不是NameData没有类型，而是NameData变量没有标注是否会含有指针信息。

通过给NameData变量增加一个NOPTR标志，表示其中不会包含指针数据可以修复该错误：
``` 
#include "textflag.h"

GLOBL ·NameData(SB),NOPTR,$8
```

通过给·NameData增加NOPTR标志档方式表示其中不含指针数据。我们也可以通过给·NameData变量在Go语言中增加一个不含指针并且大小为8个字节的类型来修改该错误：
```go
package pkg

var NameData [8]byte
var Name string
```

我们将NameData声明为长度为8的字节数组。编译器可以通过类型分析出该变量不会包含指针，因此汇编代码中可以省略NOPTR标志。现在垃圾回收器在遇到该变量的时候就会停止内部数据的扫描。

-------------------

在这个实现中，Name字符串底层其实引用的是NameData内存对应的“gopher”字符串数据。因此，如果NameData发生变化，Name字符串的数据也会跟着变化。
```gotemplate
func main() {
    println(pkg.Name)

    pkg.NameData[0] = '?'
    println(pkg.Name)
}
```

**注意：**
当然这和字符串的只读定义是冲突的，正常的代码需要避免出现这种情况。最好的方法是不要导出内部的NameData变量，这样可以避免内部数据被无意破坏。

``` 
GLOBL ·Name(SB),$24

DATA ·Name+0(SB)/8,$·Name+16(SB)
DATA ·Name+8(SB)/8,$6
DATA ·Name+16(SB)/8,$"gopher"
```
在新的结构中，Name符号对应的内存从16字节变为24字节，多出的8个字节存放底层的“gopher”字符串。·Name符号前16个字节依然对应reflect.StringHeader结构体：Data部分对应$·Name+16(SB)，表示数据的地址为Name符号往后偏移16个字节的位置；Len部分依然对应6个字节的长度。这是C语言程序员经常使用档技巧。


