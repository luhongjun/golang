package main

// 发布-订阅是一种消息传递模式，用于在不同组件之间传递消息，而这些组件不知道彼此的身份

import (
	"fmt"
	"time"
)

// 定义消息体
type Message struct {
	title	string
	content	string
}

type Receiver 	chan<- Message
type Sender		<-chan Message

// 定义订阅
type Subscription struct {
	receiver	Receiver
	messageBox	Sender
}

// 定义主题
type Topic struct {
	theme			string
	Receivers		[]*Receiver
	Subscribers    	[]User
	MessageHistory 	[]Message
}

// 注册用户（订阅主题）
func (topic *Topic) Subscribe(user User) Subscription {
	// 将用户添加至订阅名单中
	topic.Subscribers = append(topic.Subscribers, user)

	// 为当前订阅的用户创建 Message 通知的通道
	userPipeline := make(chan Message, 100)
	var receiver Receiver = userPipeline
	var messageBox Sender = userPipeline

	// 添加接收列表
	topic.Receivers = append(topic.Receivers, &receiver)

	// Create a subscription
	return Subscription{
		receiver:   receiver,
		messageBox: messageBox,
	}
}

func (topic *Topic) Publish(msg Message) {
	// 记录发布的消息
	topic.MessageHistory = append(topic.MessageHistory, msg)

	// 向全体订阅者发送消息
	for _, receiver := range topic.Receivers {
		go func(receiver *Receiver) {
			*receiver <- msg
		}(receiver)
	}
}

// 定义用户 User
type User struct {
	id		uint64
	name 	string
}

// 监听等待订阅消息
func (user *User) waitMessage(s Subscription)  {
	for  {
		msg := <- s.messageBox
		fmt.Println("通知用户：" + user.name + ";" + "信息：" + msg.content)
	}
}

func main()  {
	// 实例化主题
	topic := Topic{
		theme:          "财经类消息",
		Subscribers:    nil,
		MessageHistory: nil,
	}

	// 实例化用户
	userA := User{
		id:   100,
		name: "Lu",
	}

	userB := User{
		id:   101,
		name: "Hong",
	}

	// 用户订阅
	subscriptionA := topic.Subscribe(userA)
	subscriptionB := topic.Subscribe(userB)
	// 发布消息
	topic.Publish(Message{
		title:   "今日股市情况",
		content: "上涨0.43%",
	})

	go userA.waitMessage(subscriptionA)
	go userB.waitMessage(subscriptionB)

	time.Sleep(time.Second * time.Duration(300))
}
