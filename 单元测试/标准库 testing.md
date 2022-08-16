# 标准库 testing

**推荐阅读**
- [testing包内部实现原理](http://books.studygolang.com/GoExpertProgramming/chapter07/7.3-foreword.html)

Go编写测试可以分为三类：

- 功能测试（test）：用于测试程序的一些逻辑行为是否正确；go test命令会调用这些测试函数并报告测试结果是PASS或FAIL
  
```gotemplate
//必须以“Test”为前缀，唯一参数类型必须是 *testing.T 类型
func Test_something(t *testing.T){}
```
- 基准测试（benchmark，也称为性能测试）：用于衡量一些函数的性能;go test命令会多次运行基准测试函数以计算一个平均的执行时间
  
```gotemplate
//必须以“Benchmark”为前缀，唯一参数类型必须是 *testing.B 类型
func Benchmark_something(b *testing.B){}
```
- 示例测试（example）：提供一个由编译器保证正确性的示例文档
```gotemplate
//必须以“Example”为前缀，但对函数参数列表没有强制规定
// 检测多行输出
func ExampleSayGoodbye() {
gotest.SayGoodbye()
// OutPut:
// Hello,
// goodbye
}
```
**备注**：其用法可通过文章：http://books.studygolang.com/GoExpertProgramming/chapter07/7.1.3-example_test.html 初步了解


## 进阶测试

### 子测试

子测试提供一种在一个测试函数中执行多个测试的能力，比如原来有TestA、TestB和TestC三个测试函数，每个测试函数执行开始都需要做些相同的初始化工作，那么可以利用子测试将这三个测试合并到一个测试中，这样初始化工作只需要做一次。
```gotemplate

// sub1 为子测试，只做加法测试
func sub1(t *testing.T) {
    t.Parallel()
    //do something
}

// sub2 为子测试，只做加法测试
func sub2(t *testing.T) {
    //do something
}

// TestSub 内部调用sub1、sub2和sub3三个子测试
func TestSub(t *testing.T) {
    // setup code

    //【方式一】：
    //name参数为子测试的名字，f为子测试函数
    //Run()会启动新的协程来执行f，并阻塞等待f执行结束才返回，除非f中使用t.Parallel()设置子测试为并发
    //但是如果子测试使用 t.Parallel() 指定并发，那么就没办法共享teardown了，因为执行顺序很可能是setup->子测试1->teardown->子测试2...
    t.Run("A=1", sub1)
    t.Run("A=2", sub2)

    //【方式二】：
    //如果子测试可能并发，则可以把子测试通过Run()再嵌套一层，Run()可以保证其下的所有子测试执行结束后再返回。
    t.Run("group", func(t *testing.T) {
        t.Run("Test1", parallelTest1)
        t.Run("Test2", parallelTest2)
        t.Run("Test3", parallelTest3)
    })

    // tear-down code
}


```


### Main测试（TestMain）

我们知道子测试的一个方便之处在于可以让多个测试共享Setup和Tear-down。但这种程度的共享有时并不满足需求，有时希望在整个测试程序做一些全局的setup和Tear-down，这时就需要Main测试了。

所谓Main测试，即声明一个func TestMain(m *testing.M)，它是名字比较特殊的测试，参数类型为testing.M指针。如果声明了这样一个函数，当前测试程序将不是直接执行各项测试，而是将测试交给TestMain调度：

```gotemplate
func TestMain(m *testing.M)  {
	//setup ...
    //TestMain 运行在主 goroutine 中 , 可以在调用 m.Run 前后做任何设置和拆卸

	code := m.Run()

	//teardown ...
    //在 TestMain 函数的最后，应该使用 m.Run 的返回值作为参数去调用 os.Exit
	os.Exit(code)
}
```