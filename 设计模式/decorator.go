package main

// Decorator结构模式允许动态扩展现有对象的功能，而不改变其内部结构。

// 装饰器提供了一种灵活的方法来扩展对象的功能。

import "log"

type Object func(int) int

// 对 Object 类型的函数体或者结构体进行修饰
func LogDecorate(fn Object) Object {
	return func(n int) int {
		log.Println("Starting the execution with the integer", n)

		result := fn(n)

		log.Println("Execution is completed with the result", result)

		return result
	}
}

func Double(n int) int {
	return n * 2
}

func main()  {
	f := LogDecorate(Double)

	f(5)
	// Starting execution with the integer 5
	// Execution is completed with the result 10
}