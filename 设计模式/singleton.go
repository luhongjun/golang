package main

import "sync"

// 定义单例模型
type singleton map[string] string

// 定义全局变量
var (
	once sync.Once
	instance singleton
)

func New() singleton {
	once.Do(func() {
		instance = make(singleton)
	})

	return instance
}

s := singleton.New()

s["this"] = "that"

s2 := singleton.New()

fmt.Println("This is ", s2["this"])
// This is that
