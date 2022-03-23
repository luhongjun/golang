# Golang汇编 - 特殊符号

## 特殊字符

Go语言函数或方法符号在编译为目标文件后，目标文件中的每个符号均包含对应包的绝对导入路径。因此目标文件的符号可能非常复杂，比如“path/to/pkg.(*SomeType).SomeMethod”或“go.string."abc"”等名字。目标文件的符号名中不仅仅包含普通的字母，还可能包含点号、星号、小括弧和双引号等诸多特殊字符。而Go语言的汇编器是从plan9移植过来的二把刀，并不能处理这些特殊的字符，导致了用Go汇编语言手工实现Go诸多特性时遇到种种限制。

Go汇编语言同样遵循Go语言少即是多的哲学，它只保留了最基本的特性：定义变量和全局函数。其中在变量和全局函数等名字中引入特殊等分隔符号支持Go语言等包体系。

为了简化Go汇编器的词法扫描程序的实现，特别引入了Unicode中的`中点·`和大写的`除法/`，对应的Unicode码点为U+00B7和U+2215。汇编器编译后，`中点·`会被替换为`ASCII中的点“.”`，`大写的除法`会被替换为`ASCII码中的除法“/”`，比如math/rand·Int会被替换为math/rand.Int。这样可以将中点和浮点数中的小数点、大写的除法和表达式中的除法符号分开，可以简化汇编程序词法分析部分的实现。

即使暂时抛开Go汇编语言设计取舍的问题，在不同的操作系统不同等输入法中如何输入中点·和除法/两个字符就是一个挑战。这两个字符在 https://golang.org/doc/asm 文档中均有描述，因此直接从该页面复制是最简单可靠的方式。

如果是macOS系统，则有以下几种方法输入`中点·`：在不开输入法时，可直接用 option+shift+9 输入；如果是自带的简体拼音输入法，输入左上角~键对应·，如果是自带的Unicode输入法，则可以输入对应的Unicode码点。其中Unicode输入法可能是最安全可靠等输入方式。


## 无需分号

Go汇编语言中分号可以用于分隔同一行内的多个语句。下面是用分号混乱排版的汇编代码：
``` 
TEXT ·main(SB), $16-0; MOVQ ·helloworld+0(SB), AX; MOVQ ·helloworld+8(SB), BX;
MOVQ AX, 0(SP);MOVQ BX, 8(SP);CALL runtime·printstring(SB);
CALL runtime·printnl(SB);
RET;
```

和Go语言一样，也可以省略行尾的分号。当遇到末尾时，汇编器会自动插入分号。下面是省略分号后的代码：
``` 
TEXT ·main(SB), $16-0
    MOVQ ·helloworld+0(SB), AX; MOVQ AX, 0(SP)
    MOVQ ·helloworld+8(SB), BX; MOVQ BX, 8(SP)
    CALL runtime·printstring(SB)
    CALL runtime·printnl(SB)
    RET
```

和Go语言一样，语句之间多个连续的空白字符和一个空格是等价的。