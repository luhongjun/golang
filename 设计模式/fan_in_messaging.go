package main

// （Fan-In Messaging Patterns）扇形传入消息的设计模式

// 将复杂且多样的数据整理到一个有序的结构中
// 场景：用于在工作人员（客户端：源，服务器：目标）之间创建工作漏斗。

import "sync"

type WriteOut chan int
type WriteIn chan<- int

// Merge different channels in one channel
func Merge(ins ...WriteIn) WriteOut {
	var wg sync.WaitGroup
	wg.Add(len(ins))

	// 定义输出漏斗
	out := make(chan int)

	// Start an send goroutine for each input channel in cs. send
	// copies values from c to out until c is closed, then calls wg.Done.
	// 写入漏斗
	for _, in := range ins {
		go func(c WriteIn) {
			for n := range c {
				out <- n
			}
			wg.Done()
		}(in)
	}

	// Start a goroutine to close out once all the send goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main()  {
	// channel A
	channelA := make(WriteIn, 3)
	channelA <- 1
	channelA <- 2
	channelA <- 3

	// channel B
	channelB := make(WriteIn, 1)
	channelB <- 4

	_ = Merge(channelA, channelB)
}