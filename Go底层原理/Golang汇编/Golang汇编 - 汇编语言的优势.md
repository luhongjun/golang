# Golang汇编 - 汇编语言的优势

汇编语言的真正威力来自两个维度：

1. 突破框架限制，实现看似不可能的任务；可以通过Go汇编语言直接访问系统调用，和直接调用C语言函数
2. 突破指令限制，通过高级指令挖掘极致的性能

## 系统调用

系统调用是操作系统为外提供的公共接口。因为操作系统彻底接管了各种底层硬件设备，因此操作系统提供的系统调用成了实现某些操作的唯一方法。从另一个角度看，系统调用更像是一个RPC远程过程调用，不过信道是寄存器和内存。在系统调用时，我们向操作系统发送调用的编号和对应的参数，然后阻塞等待系统调用地返回。因为涉及到阻塞等待，因此系统调用期间的CPU利用率一般是可以忽略的。另一个和RPC地远程调用类似的地方是，操作系统内核处理系统调用时不会依赖用户的栈空间，一般不会导致爆栈发生。因此系统调用是最简单安全的一种调用了。

系统调用虽然简单，但是它是操作系统对外的接口，因此不同的操作系统调用规范可能有很大地差异。

我们先看看Linux在AMD64架构上的系统调用规范，在`syscall/asm_linux_amd64.s`文件中有注释说明：
``` 
//
// System calls for AMD64, Linux
//

// func Syscall(trap int64, a1, a2, a3 uintptr) (r1, r2, err uintptr);
// Trap # in AX, args in DI SI DX R10 R8 R9, return in AX DX
// Note that this differs from "standard" ABI convention, which
// would pass 4th arg in CX, not R10.
```
这是syscall.Syscall函数的内部注释，简要说明了Linux系统调用的规范。系统调用的前6个参数直接由DI、SI、DX、R10、R8和R9寄存器传输，结果由AX和DX寄存器返回。macOS等类UINX系统调用的参数传输大多数都采用类似的规则。

macOS的系统调用编号在`/usr/include/sys/syscall.h`头文件，Linux的系统调用号在`/usr/include/asm/unistd.h`头文件。虽然在UNIX家族中是系统调用的参数和返回值的传输规则类似，但是不同操作系统提供的系统调用却不是完全相同的，因此系统调用编号也有很大的差异。以UNIX系统中著名的write系统调用为例，在macOS的系统调用编号为4，而在Linux的系统调用编号却是1。

--------------------------

我们将基于write系统调用包装一个字符串输出函数。下面的代码是macOS版本：
``` 
// func SyscallWrite_Darwin(fd int, msg string) int
TEXT ·SyscallWrite_Darwin(SB), NOSPLIT, $0
    MOVQ $(0x2000000+4), AX // #define SYS_write 4
    MOVQ fd+0(FP),       DI
    MOVQ msg_data+8(FP), SI
    MOVQ msg_len+16(FP), DX
    SYSCALL
    MOVQ AX, ret+0(FP)
    RET
```
其中第一个参数是输出文件的文件描述符编号，第二个参数是字符串的头部。字符串头部是由reflect.StringHeader结构定义，第一成员是8字节的数据指针，第二个成员是8字节的数据长度。在macOS系统中，执行系统调用时还需要将系统调用的编号加上0x2000000后再行传入AX。然后再将fd、数据地址和长度作为write系统调用的三个参数输入，分别对应DI、SI和DX三个寄存器。最后通过SYSCALL指令执行系统调用，系统调用返回后从AX获取返回值。

这样我们就基于系统调用包装了一个定制的输出函数。在UNIX系统中，标准输入stdout的文件描述符编号是1，因此我们可以用1作为参数实现字符串的输出：
```gotemplate
func SyscallWrite_Darwin(fd int, msg string) int

func main() {
    if runtime.GOOS == "darwin" {
        SyscallWrite_Darwin(1, "hello syscall!\n")
    }
}
```

## 直接调用C函数

