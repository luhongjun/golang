# container - 容器数据类型：heap、list 和 ring

该包实现了三个复杂的数据结构：堆，链表，环。 这个包就意味着你使用这三个数据结构的时候不需要再费心从头开始写算法了。

- 堆 heap
- 链表 list
- 环 ring

## 堆heap

堆使用的数据结构是最小二叉树，即根节点比左边子树和右边子树的所有值都小。 go 的堆包只是实现了一个接口，我们看下它的定义：
```gotemplate
type Interface interface {
    sort.Interface
    Push(x interface{}) // add x as element Len()
    Pop() interface{}   // remove and return element Len() - 1.
}
```
堆结构继承自 sort.Interface,需实现`Len() int`/`Less(i, j int) bool`/`Swap(i, j int)`方法。

### 内部原理

内部实现，是使用最小堆，索引排序从根节点开始，然后左子树，右子树的顺序方式。索引布局如下：
```
                  0
            1            2
         3    4       5      6
        7 8  9 10   11
假设 (heap[1]== 小明 ) 它的左子树 (heap[3]== 小黑 ) 和右子树 (heap[4]== 大黄 ) 且 小明 > 小黑 > 大黄 ;
```

堆内部实现了 down 和 up 函数 : down 函数用于将索引 i 处存储的值 ( 设 i=1, 即小明 ) 与它的左子树 ( 小黑 ) 和右子树 ( 大黄 ) 相比 , 将三者最小的值大黄与小明的位置交换，交换后小明继续与交换后的子树 (heap[9]和 heap[10]) 相比，重复以上步骤，直到小明位置不变。

假设 heap[11]== 阿花 当从堆中 Pop 一个元素的时候，先把元素和最后一个节点的值 ( 阿花 ) 交换，然后弹出，然后对阿花调用 down，向下保证最小堆。

当往堆中 Push 一个元素的时候，这个元素插入到最后一个节点，本例中为 heap[12]，即作为 heap[5]的右子树，然后调用 up 函数向上比较。

### 实现示例

```gotemplate
// 定义堆
type HeapInt []int


```

## 链表 list

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

list对应的方法有如下：
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

## 环 Ring

环的结构有点特殊，环的尾部就是头部，所以每个元素实际上就可以代表自身的这个环。 它不需要像 list 一样保持 list 和 element 两个结构，只需要保持一个结构就行。

因此，环是一个无序的结构

```gotemplate
type Ring struct {
    next, prev *Ring
    Value      interface{}
}
```

Ring提供的方法有如下：
```gotemplate
type Ring
    func New(n int) *Ring  // 初始化环,n表示设置的环大小；
    func (r *Ring) Do(f func(interface{}))  // 循环环进行操作
    func (r *Ring) Len() int // 环长度
    func (r *Ring) Link(s *Ring) *Ring // 连接两个环
    func (r *Ring) Move(n int) *Ring // 指针从当前元素开始向后移动或者向前（n 可以为负数）
    func (r *Ring) Next() *Ring // 当前元素的下个元素
    func (r *Ring) Prev() *Ring // 当前元素的上个元素
    func (r *Ring) Unlink(n int) *Ring // 从当前元素开始，删除 n 个元素
```