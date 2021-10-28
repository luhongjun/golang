// Package main serves as an example application that makes use of the observer pattern.
// Playground: https://play.golang.org/p/cr8jEmDmw0
package main

import (
	"fmt"
	"time"
)

// 定义事件
type (
	// Event defines an indication of a point-in-time occurrence.
	Event struct {
		// Data in this case is a simple int, but the actual
		// implementation would depend on the application.
		Data int64
	}

	// Observer defines a standard interface for instances that wish to list for
	// the occurrence of a specific event.
	Observer interface {
		// OnNotify allows an event to be "published" to interface implementations.
		// In the "real world", error handling would likely be implemented.
		OnNotify(Event)
	}
)

// 定义事件的观察者
type (
	// Notifier is the instance being observed. Publisher is perhaps another decent
	// name, but naming things is hard.
	Notifier interface {
		// Register allows an instance to register itself to listen/observe
		// events.
		Register(Observer)
		// Deregister allows an instance to remove itself from the collection
		// of observers/listeners.
		Deregister(Observer)
		// Notify publishes new events to listeners. The method is not
		// absolutely necessary, as each implementation could define this itself
		// without losing functionality.
		Notify(Event)
	}

	eventObserver struct{
		id int
	}

	eventNotifier struct{
		// Using a map with an empty struct allows us to keep the observers
		// unique while still keeping memory usage relatively low.
		observers map[Observer]struct{}
	}
)

func (o *eventObserver) OnNotify(e Event) {
	fmt.Printf("*** Observer %d received: %d\n", o.id, e.Data)
}

func (o *eventNotifier) Register(l Observer) {
	o.observers[l] = struct{}{}
}

func (p *eventNotifier) Notify(e Event) {
	for o := range p.observers {
		o.OnNotify(e)
	}
}

func main() {
	// Initialize a new Notifier.
	n := eventNotifier{
		observers: map[Observer]struct{}{},
	}

	// Register a couple of observers.
	n.Register(&eventObserver{id: 1})
	n.Register(&eventObserver{id: 2})

	// A simple loop publishing the current Unix timestamp to observers.
	// @see https://studygolang.com/articles/18404?fr=sidebar  关于 time 包的 Timer 定时器
	stop := time.NewTimer(10 * time.Second).C
	tick := time.NewTicker(time.Second).C
	for {
		select {
		case <- stop://	10秒后自动结束
			return
		case t := <-tick:// 每1秒进行一次通知
			n.Notify(Event{Data: t.UnixNano()})
		}
	}
}