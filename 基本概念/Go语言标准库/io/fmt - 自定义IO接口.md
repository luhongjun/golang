# Print 序列函数

## Stringer 接口

```gotemplate
type Stringer interface {
    String() string
}
```

示例：
```gotemplate
type Person struct {
    Name string
    Age  int
    Sex  int
}

func (this *Person) String() string {
    return "This is polaris, He is 28 years old"
}
```

## Formatter 接口

```gotemplate
// Formatter 接口由带有定制的格式化器的值所实现。 Format 的实现可调用 Sprintf 或 Fprintf(f) 等函数来生成其输出。
type Formatter interface {
    Format(f State, c rune)
}
```

**解释说明：**

1. fmt.State 是一个接口。由于 Format 方法是被 fmt 包调用的，它内部会实例化好一个 fmt.State 接口的实例，我们不需要关心该接口；
2. 可以实现自定义占位符，但是同时 fmt 包中和类型相对应的预定义占位符会无效。
3. 实现了 Formatter 接口，相应的 Stringer 接口不起作用。但实现了 Formatter 接口的类型应该实现 Stringer 接口，这样方便在 Format 方法中调用 String() 方法；
4. Format 方法的第二个参数是占位符中%后的字母（有精度和宽度会被忽略，只保留字母）；


示例：
```gotemplate
type Person struct {
    Name string
    Age  int
    Sex  int
}

func (this *Person) Format(f fmt.State, c rune) {
    if c == 'L' {
        f.Write([]byte(this.String()))
        f.Write([]byte(" Person has three fields."))
    } else {
        // 没有此句，会导致 fmt.Printf("%s", p) 啥也不输出
        f.Write([]byte(fmt.Sprintln(this.String())))
    }
}

// 运行
p := &Person{"polaris", 28, 0}
fmt.Printf("%L", p) //This is polaris, He is 28 years old. Person has three fields.
```

## GoStringer 接口

```gotemplate
// 该接口定义了类型的Go语法格式。用于打印(Printf)格式化占位符为 %#v 的值
type GoStringer interface {
    GoString() string
}
```

示例：
```gotemplate
 func (this *Person) GoString() string {
    return "&Person{Name is "+this.Name+", Age is "+strconv.Itoa(this.Age)+", Sex is "+strconv.Itoa(this.Sex)+"}"
}

//运行
p := &Person{"polaris", 28, 0}
fmt.Printf("%#v", p) //输出： &Person{Name is polaris, Age is 28, Sex is 0}
```

# Scan 序列函数

该序列函数和 Print 序列函数相对应，包括：Fscan/Fscanf/Fscanln/Sscan/Sscanf/Sscanln/Scan/Scanf/Scanln。

一般的，我们将Fscan/Fscanf/Fscanln归为一类；Sscan/Sscanf/Sscanln归为一类；Scan/Scanf/Scanln归为另一类。其中，Scan/Scanf/Scanln会调用相应的F开头一类函数。如：
```gotemplate
func Scan(a ...interface{}) (n int, err error) {
    return Fscan(os.Stdin, a...)
}
```