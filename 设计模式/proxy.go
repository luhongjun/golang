package main

import "fmt"

// To use proxy and to object they must implement same methods
type IObject interface {
	ObjDo(action string)
}

// 定义结构体
type Object struct {
	action string
}

// Object 结构体继承接口 IObject，并实现 ObjDo 方法
func (obj *Object) ObjDo(action string) {
	// Action behavior
	fmt.Printf("I can, %s", action)
}

// 继承 Object 结构体
type ProxyObject struct {
	object *Object
}

// ObjDo are implemented IObject and intercept action before send in real Object
func (p *ProxyObject) ObjDo(action string) {
	if p.object == nil {
		p.object = new(Object)
	}
	if action == "Run" {
	    //间接通过 ProxyObject 来实现 object 的 ObjDo接口
		p.object.ObjDo(action) // Prints: I can, Run
	}
}


