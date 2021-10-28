package main

// 策略行为设计模式允许在运行时选择算法的行为。
//【注意注意】是运行时选择算法！！！ 策略模式允许您更改对象的内部
//
// 它定义算法，封装算法，并互换使用算法。

import "fmt"

type Operator interface {
	Apply(int, int) int
}

type Operation struct {
	Operator Operator
}

func (o *Operation) Operate(leftValue, rightValue int) int {
	return o.Operator.Apply(leftValue, rightValue)
}

type Addition struct{}

func (Addition) Apply(lval, rval int) int {
	return lval + rval
}

type Multiplication struct{}

func (Multiplication) Apply(lval, rval int) int {
	return lval * rval
}

// 使用示例
func main()  {
	add := Operation{Addition{}}
	fmt.Println(add.Operate(3, 5)) // 8

	mult := Operation{Multiplication{}}
	fmt.Println(mult.Operate(3, 5)) // 15
}