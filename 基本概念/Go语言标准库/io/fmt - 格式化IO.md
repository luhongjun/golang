# fmt - 格式化IO

fmt 包实现了格式化I/O函数，类似于C的 printf 和 scanf. 格式“占位符”衍生自C，但比C更简单。

## Printing 输出

```gotemplate
type user struct {
    name string
}

func main() {
    u := user{"tang"}
    //Printf 格式化输出
    fmt.Printf("% + v\n", u)     //格式化输出结构
    fmt.Printf("%#v\n", u)       //输出值的 Go 语言表示方法
    fmt.Printf("%T\n", u)        //输出值的类型的 Go 语言表示
    fmt.Printf("%t\n", true)     //输出值的 true 或 false
    fmt.Printf("%b\n", 1024)     //二进制表示
    fmt.Printf("%c\n", 11111111) //数值对应的 Unicode 编码字符
    fmt.Printf("%d\n", 10)       //十进制表示
    fmt.Printf("%o\n", 8)        //八进制表示
    fmt.Printf("%q\n", 22)       //转化为十六进制并附上单引号
    fmt.Printf("%x\n", 1223)     //十六进制表示，用a-f表示
    fmt.Printf("%X\n", 1223)     //十六进制表示，用A-F表示
    fmt.Printf("%U\n", 1233)     //Unicode表示
    fmt.Printf("%b\n", 12.34)    //无小数部分，两位指数的科学计数法6946802425218990p-49
    fmt.Printf("%e\n", 12.345)   //科学计数法，e表示
    fmt.Printf("%E\n", 12.34455) //科学计数法，E表示
    fmt.Printf("%f\n", 12.3456)  //有小数部分，无指数部分
    fmt.Printf("%g\n", 12.3456)  //根据实际情况采用%e或%f输出
    fmt.Printf("%G\n", 12.3456)  //根据实际情况采用%E或%f输出
    fmt.Printf("%s\n", "wqdew")  //直接输出字符串或者[]byte
    fmt.Printf("%q\n", "dedede") //双引号括起来的字符串
    fmt.Printf("%x\n", "abczxc") //每个字节用两字节十六进制表示，a-f表示
    fmt.Printf("%X\n", "asdzxc") //每个字节用两字节十六进制表示，A-F表示
    fmt.Printf("%p\n", 0x123)    //0x开头的十六进制数表示
}
```


### 自定义格式化

若一个操作数实现了 Formatter 接口，该接口就能更好地用于控制格式化。

若其格式（它对于 Println 等函数是隐式的 %v）对于字符串是有效的 （%s %q %v %x %X），以下两条规则也适用：

1. 若一个操作数实现了 error 接口，Error 方法就能将该对象转换为字符串，随后会根据占位符的需要进行格式化。

```gotemplate
// 来自 fmt 包
type error interface {
	Error() string
}
```   

2. 若一个操作数实现了 String() string 方法，该方法能将该对象转换为字符串，随后会根据占位符的需要进行格式化。

```gotemplate
// 来自 fmt 包
type Stringer interface {
	String() string
}
```

### 格式化发生错误

如果给占位符提供了无效的实参（例如将一个字符串提供给 %d），所生成的字符串会包含该问题的描述，如下例所示：

```gotemplate
类型错误或占位符未知：%!verb(type=value)
        Printf("%d", hi):          %!d(string=hi)
实参太多：%!(EXTRA type=value)
    Printf("hi", "guys"):      hi%!(EXTRA string=guys)
实参太少： %!verb(MISSING)
    Printf("hi%d"):            hi %!d(MISSING)
宽度或精度不是int类型: %!(BADWIDTH) 或 %!(BADPREC)
    Printf("%*s", 4.5, "hi"):  %!(BADWIDTH)hi
    Printf("%.*s", 4.5, "hi"): %!(BADPREC)hi
```

**注意：**
所有错误都始于“%!”，有时紧跟着单个字符（占位符），并以小括号括住的描述结尾。


## Scanning

一组类似的函数通过扫描已格式化的文本来产生值。 

Scan、Scanf 和 Scanln 从 os.Stdin 中读取； 

Fscan、Fscanf 和 Fscanln 从指定的 io.Reader 中读取； 

Sscan、Sscanf 和 Sscanln 从实参字符串中读取。 

Scanln、Fscanln 和 Sscanln 在换行符处停止扫描，且需要条目紧随换行符之后； 

Scanf、Fscanf 和 Sscanf 需要输入换行符来匹配格式中的换行符；其它函数则将换行符视为空格。

### 自定义扫描

在所有的扫描参数中，若一个操作数实现了 Scan 方法（即它实现了 Scanner 接口）， 该操作数将使用该方法扫描其文本。此外，若已扫描的实参数少于所提供的实参数，就会返回一个错误。

```gotemplate
type Scanner interface {
	Scan(state ScanState, verb rune) error
}
```

**注意：**
所有需要被扫描的实参都必须是基本类型或 Scanner 接口的实现