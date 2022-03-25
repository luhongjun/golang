# reflect - 反射

## 关于静态类型

Go是静态类型语言，比如"int"、"float32"、"[]byte"等等。每个变量都有一个静态类型，且在编译时就确定了。 那么考虑一下如下一种类型声明:
``` 
type Myint int

var i int
var j Myint
```

Q: i 和j 类型相同吗？ 

A：i 和j类型是不同的。 二者拥有不同的静态类型，没有类型转换的话是不可以互相赋值的，尽管二者底层类型是一样的。

## 特殊的静态类型interface

interface类型是一种特殊的类型，它代表方法集合。 它可以存放任何实现了其方法的值。

经常被拿来举例的是io包里的这两个接口类型：
``` 
// Reader is the interface that wraps the basic Read method.
type Reader interface {
    Read(p []byte) (n int, err error)
}

// Writer is the interface that wraps the basic Write method.
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

任何类型，比如某struct，只要实现了其中的Read()方法就被认为是实现了Reader接口，只要实现了Write()方法，就被认为是实现了Writer接口，不过方法参数和返回值要跟接口声明的一致。

接口类型的变量可以存储任何实现该接口的值。

## 特殊的interface类型

最特殊的interface类型为空interface类型，即interface {}，前面说了，interface用来表示一组方法集合，所有实现该方法集合的类型都被认为是实现了该接口。那么空interface类型的方法集合为空，也就是说所有类型都可以认为是实现了该接口。

一个类型实现空interface并不重要，重要的是一个空interface类型变量可以存放所有值，记住是所有值，这才是最最重要的。 这也是有些人认为Go是动态类型的原因，这是个错觉。

## 反射三定律

1. 反射可以将interface类型变量转换成反射对象

```gotemplate
var x float64 = 3.4
t := reflect.TypeOf(x)  //t is reflext.Type
fmt.Println(t) //输出：float64

v := reflect.ValueOf(x)     //v is reflext.Value
fmt.Println(v)    //输出：3.4
```

2. 反射可以将反射对象还原成interface对象

```gotemplate
var x float64 = 3.4
v := reflect.ValueOf(x) //v is reflext.Value

var y float64 = v.Interface().(float64)
fmt.Println(y)    //输出： 3.4
```

3. 反射对象可修改，value值必须是可设置的

```gotemplate
var x float64 = 3.4
v := reflect.ValueOf(&x)
v.SetFloat(7.1) // Error: will panic.
```
错误原因即是v是不可修改的。

反射对象是否可修改取决于其所存储的值，回想一下函数传参时是传值还是传址就不难理解上例中为何失败了。

上例中，传入reflect.ValueOf()函数的其实是x的值，而非x本身。即通过v修改其值是无法影响x的，也即是无效的修改，所以golang会报错。

-----------------------------

想到此处，即可明白，如果构建v时使用x的地址就可实现修改了，但此时v代表的是指针地址，我们要设置的是指针所指向的内容，也即我们想要修改的是*v。 那怎么通过v修改x的值呢？

reflect.Value提供了`Elem()`方法，可以获得指针向指向的value。看如下代码：
```gotemplate
var x float64 = 3.4
v := reflect.ValueOf(&x)
v.Elem().SetFloat(7.1)
fmt.Println(x) //输出：7.1
```
