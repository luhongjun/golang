## list链表

链表就是一个有 prev 和 next 指针的数组了。它维护两个 type，( 注意，这里不是 interface)

```gotemplate
type Element struct {
    next, prev *Element  // 上一个元素和下一个元素
    list *List  // 元素所在链表
    Value interface{}  // 元素
}

type List struct {
    root Element  // 链表的根元素
    len  int      // 链表的长度
}
```
基本使用是先创建 list，然后往 list 中插入值，list 就内部创建一个 Element，并内部设置好 Element 的 next,prev 等。具体可以看下例子：
```go
// This example demonstrates an integer heap built using the heap interface.
package main

import (
    "container/list"
    "fmt"
)

func main() {
    list := list.New()
    list.PushBack(1)
    list.PushBack(2)

    fmt.Printf("len: %v\n", list.Len())
    fmt.Printf("first: %#v\n", list.Front())
    fmt.Printf("second: %#v\n", list.Front().Next())
}
```


output:
```shell
len: 2
first: &list.Element{next:(*list.Element)(0x2081be1b0), prev:(*list.Element)(0x2081be150), list:(*list.List)(0x2081be150), Value:1}
second: &list.Element{next:(*list.Element)(0x2081be150), prev:(*list.Element)(0x2081be180), list:(*list.List)(0x2081be150), Value:2}
```

list 对应的方法有：
```gotemplate
type Element
func (e *Element) Next() *Element
func (e *Element) Prev() *Element
type List
func New() *List
func (l *List) Back() *Element   // 最后一个元素
func (l *List) Front() *Element  // 第一个元素
func (l *List) Init() *List  // 链表初始化
func (l *List) InsertAfter(v interface{}, mark *Element) *Element // 在某个元素后插入
func (l *List) InsertBefore(v interface{}, mark *Element) *Element  // 在某个元素前插入
func (l *List) Len() int // 在链表长度
func (l *List) MoveAfter(e, mark *Element)  // 把 e 元素移动到 mark 之后
func (l *List) MoveBefore(e, mark *Element)  // 把 e 元素移动到 mark 之前
func (l *List) MoveToBack(e *Element) // 把 e 元素移动到队列最后
func (l *List) MoveToFront(e *Element) // 把 e 元素移动到队列最头部
func (l *List) PushBack(v interface{}) *Element  // 在队列最后插入元素
func (l *List) PushBackList(other *List)  // 在队列最后插入接上新队列
func (l *List) PushFront(v interface{}) *Element  // 在队列头部插入元素
func (l *List) PushFrontList(other *List) // 在队列头部插入接上新队列
func (l *List) Remove(e *Element) interface{} // 删除某个元素
```
