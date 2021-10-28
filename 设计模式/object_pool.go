package main

import "fmt"

// 定义对象池元素的结构
type Object struct {
	name 	string
	id 		string
}

func (*Object) Do()  {
	fmt.Println("do something")
}

// 定义一个对象池（如果通过"通道"的方式，获取和添加都可能存在阻塞；如果用切片的话会没有这个顾虑）
type Pool chan *Object

func New(total int) *Pool {
	p := make(Pool, total)

	for i := 0; i < total; i++ {
		p <- new(Object)
	}

	return &p
}

func main()  {
	p := New(2)

	select {
		case obj := <-*p:
			obj.Do( /*...*/ )
		default:
			// No more objects left — retry later or fail
			return
	}
}