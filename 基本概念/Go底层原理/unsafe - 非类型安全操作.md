# unsafe 非类型安全操作

unsafe库徘徊在“类型安全”边缘，由于它们绕过了 Golang 的内存安全原则，一般被认为使用该库是不安全的。但是，在许多情况下，unsafe库的作用又是不可替代的，灵活地使用它们可以实现对内存的直接读写操作。在reflect库、syscall库以及其他许多需要操作内存的开源项目中都有对它的引用。

unsafe库源码极少，只有两个类型的定义和三个方法的声明。

## 常见函数

| 函数签名 | 说明 |
| --- | --- |
| func Sizeof(v ArbitraryType) uintptr | 返回变量 v 占用的内存空间的字节数，该字节数不是按照变量 v 实际占用的内存计算，而是按照 v 的“ top level ”内存计算。比如，在 64 位系统中，如果变量 v 是 int 类型，会返回 16，因为 v 的“ top level ”内存就是它的值使用的内存；如果变量 v 是 string 类型，会返回 16，因为 v 的“ top level ”内存不是存放着实际的字符串，而是该字符串的地址；如果变量 v 是 slice 类型，会返回 24，这是因为 slice 的描述符就占了 24 个字节 |
| func Offsetof(v ArbitraryType) uintptr | 返回由 v 所指示的某结构体中的字段在该结构体中的位置偏移字节数，注意，v 的表达方式必须是“ struct.filed ”形式 |
| func Alignof(v ArbitraryType) uintptr | 相当于`reflect.TypeOf(x).Align()` |


## Arbitrary 类型与 Pointer 类型 

```gotemplate
type Pointer *ArbitraryType
```

- 官方导出类型`ArbitraryType`只是出于完善文档的考虑，在其他的库和任何项目中都没有使用价值，除非程序员故意使用它。

- 类型`Pointer`比较重要，它是实现定位欲读写的内存的基础。官方文档对该类型有四个重要描述：
    1. 任何类型的指针都可以被转化为 Pointer
    2. Pointer 可以被转化为任何类型的指针
    3. uintptr 可以被转化为 Pointer
    4. Pointer 可以被转化为 uintptr
    
针对第一二点，我们可以实现类型转换：
```go
package main

import (
	"fmt"
	"unsafe"
)

func main()  {
	var f float64 = 5.5
	ptr := (*int64)(unsafe.Pointer(&f))
	fmt.Println(*ptr)       //输出：4617878467915022336
    fmt.Println(int64(f))   //输出：5
}
```
可见，由于通过不同的数据类型在内存中的数据布局不一样，并不能直接转换，否则容易达到非预期的结果

