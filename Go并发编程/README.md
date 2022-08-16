# Go并发编程

## 相关概念

| 名词 | 解释说明 |
| --- | --- |
| 竞态条件（race condition） | 一旦数据被多个线程共享，那么就很有可能会产生争用和冲突的情况，这种情况被称为竞态条件，这往往会破坏数据的一致性。举个例子，同时有多个线程连续向同一个缓冲区写入数据块，在线程A还没写完一个数据块的时候，线程B就开始写入另外一个数据块。 |
| 共享资源 | 比如存储资源、计算资源、I/O资源、网络资源等 |
| 同步（sync） | 指的是控制多个线程对共享资源的访问；同步可以避免多个线程在同一时刻操作同一个共享资源，也可以避免它们在同一时刻执行同一个共享资源。 |
| 原子操作 | 原子操作在进行过程中是不允许中断的；在底层，由CPU提供芯片级别的支持，所以绝对有效。即便在拥有多CPU核心，或者多CPU的计算机系统中，原子操作的保证也是不可撼动的。 |


## 相关的官方包与数据结构

```text
//互斥锁
sync.Mutex： Lock / Unlock
//读写锁
sync.RWMutex：RLock （读锁，不能写） / Lock（写锁，不能读写）

//条件变量+互斥锁
sync.Cond： Wait（等待通知）、单发通知（Signal）和广播通知（Broadcast）
----------------
var mailbox uint8 //定义信箱
var lock sync.RWMutex
cond := sync.Cond(sync.locker)
sendCond 
----------------

//原子性操作
sync/atomic
原子操作的运算： 加法（Add）、比较并交换（compare and swap， CAS）、加载（load）、存储（store）、交换（swap）
适用的数据类型：int32、int64、uint32、uint64、uintptr、unsafe.Pointer
atomic.Value：Store / load

//
sync.WaitGroup：适用于一对多的goroutine协作流程，三个指针方法：Add、Done和Wait
sync.Once：只保证执行一次

//临时对象池，一般当作针对某种数据的缓存来用
sync.Pool：结构体类型，只有两个方法： Put（在当前池存放对象） / Get（在当前池获取对象）

//并发安全字典
sync.Map：


## race


## unsafe
## runtime

## atomic


```

