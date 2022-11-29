# Go1.17新特性



Go1.17 主要新增或者变更如下：

- 编译优化：go1.17将使用栈传递参数和返回值替换为使用寄存器。实现性能提升5%，最终生成的二进制包大小减少2%（目前支持Linux、macOS、Windows的64位X86架构。官方表示后续会支持更多架构）
- 跨平台支持：支持windows系统64位ARM架构
- go module新增`pruned module graphs`功能：当go.mod文件中指定了go 1.17或者更高版本，且依赖的包同样是go 1.17或者更高版本，go.mod中只保留直接依赖
- 新增语言特性
  - 新增 unsafe.Add 方法，定义为`func Add(ptr Pointer, len IntegerType) Pointer`，方便指针运算，是`Pointer(uintptr(ptr) + uintptr(len)`的简化；
  - 新增 unsafe.Slice 方法，定义为`func Slice(ptr *ArbitraryType, len IntegerType) []ArbitraryType`，方便将指针转换为slice，是`(*[len]ArbitraryType)(unsafe.Pointer(ptr))[:]`的简化
  - 支持从slice到array转换：slice转为array指针将生成一个指向slice底层数组的指针。如果slice的长度小于array的长度，会panic
- 性能提升和bug修复
  - 提升crypto/x509性能
  - 修复URL Query解析bug：在1.17版本前，分号（;）和&一样可以作为参数的分隔符号。1.17版本去掉了分号作为分隔符
- 泛型前瞻：go预计会在1.18版本正式发布泛型，在这之前可以通过一些其他方法来体验，例如通过编译go master分支源码，或者使用官方提供的[Play ground](https://gotipplay.golang.org/)

### 使用unsafe.Slice方法
```go
package main

import (
    "fmt"
    "unsafe"
    "reflect"
)

func main() {
    s := []int{1,2,3}
    fmt.Println(s)
    printSliceHeader(&s)
    s1 := unsafe.Slice(&s[0], 3)
    fmt.Println(s1)
    printSliceHeader(&s1)
}

//输出：
// [1 2 3]
// &{824634818560 3 3}
// [1 2 3]
// &{824634818560 3 3}

func printSliceHeader(s *[]int) {
    header := (*reflect.SliceHeader)(unsafe.Pointer(s))
    fmt.Println(header)
}
```

### 实现slice转换为数组

```go
package main

import (
    "fmt"
    "reflect"
    "unsafe"
)

func main() {
    s := []int{1, 2, 3}
    fmt.Println("slice中内容：", s)
    printSliceHeader(&s)

    s0 := (*[3]int)(s)
    fmt.Println("array中内容：", s0)
    fmt.Printf("array地址：%p \n", s0)
    fmt.Println("")

    s0[0] = 0
    fmt.Println("1. 改变slice中的元素，array中元素同样改变")
    fmt.Println("slice中内容：", s)
    fmt.Println("array中内容：", s0)
    fmt.Println("")

    fmt.Println("2. 对slice进行扩容")
    s = append(s, 4)
    fmt.Println("slice中内容：", s)
    printSliceHeader(&s)
    fmt.Println("array中内容：", s0)
    fmt.Printf("array地址：%p \n", s0)
    fmt.Println("")

    fmt.Println("3. array传参会对数组中元素进行复制")
    arrayParam(*s0)
    fmt.Println("array中内容：", s0)
}

func printSliceHeader(s *[]int) {
    header := (*reflect.SliceHeader)(unsafe.Pointer(s))
    //fmt.Println(header)
    fmt.Printf("slice底层数组地址：0x%x \n", header.Data)
}

func arrayParam(s0 [3]int) {
    s0[0] = 10
}

/**
输出如下：
slice中内容： [1 2 3]
slice底层数组地址：0xc000016018 
array中内容： &[1 2 3]
array地址：0xc000016018 

1. 改变slice中的元素，array中元素同样改变
slice中内容： [0 2 3]
array中内容： &[0 2 3]

2. 对slice进行扩容
slice中内容： [0 2 3 4]
slice底层数组地址：0xc000078060 
array中内容： &[0 2 3]
array地址：0xc000016018 

3. array传参会对数组中元素进行复制
array中内容： &[0 2 3]
*/
```
