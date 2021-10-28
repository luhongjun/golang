package main

import (
	"fmt"
	"time"
)

// （Fan-Out Messaging Pattern）扇形传出消息模式

// 就是将一个消息发散到多个发送者中；
// 但是当某个发送者A处理了这条消息，其他发送者则无法处理这条消息了，因此由于Golang channel的特性，一条消息仅能一个发送者处理

// 定义消息体
type Message struct {
	data 	string
	code	int
}

// 定义接受者类型
type Receiver chan<- Message
// 定义发送者类型
type Sender <-chan Message
type Senders []Sender

// 生成接收者
func NewReceiver() chan Message {
	// 定义信息传递的信道
	var pipeline = make(chan Message, 4)

	// 定义一个接受者，缓存为 4
	var receiver Receiver = pipeline

	receiver <- Message{
		data: "hello",
		code: 4,
	}
	receiver <- Message{
		data: "hi",
		code: 5,
	}

	return pipeline
}

// 生成发送者（基于与接受者相同的通道）
func Split(pipeline chan Message, length int) Senders {
	var sender Sender = pipeline
	senders := make(Senders, length)

	// 初始化
	for key := range senders {
		senders[key] = sender
	}

	return senders
}

func main()  {
	pipeline := NewReceiver()

	senders := Split(pipeline, 4)

	for _, sender := range senders{
		go func(sender Sender) {
			value := <-sender

			fmt.Println(value)
		}(sender)
	}

	time.Sleep(time.Second * time.Duration(300))
}
