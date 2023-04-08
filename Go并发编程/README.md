# Go并发编程

## 相关概念

| 名词 | 解释说明 |
| --- | --- |
| 竞态条件（race condition） | 一旦数据被多个线程共享，那么就很有可能会产生争用和冲突的情况，这种情况被称为竞态条件，这往往会破坏数据的一致性。举个例子，同时有多个线程连续向同一个缓冲区写入数据块，在线程A还没写完一个数据块的时候，线程B就开始写入另外一个数据块。 |
| 共享资源 | 比如存储资源、计算资源、I/O资源、网络资源等 |
| 同步（sync） | 指的是控制多个线程对共享资源的访问；同步可以避免多个线程在同一时刻操作同一个共享资源，也可以避免它们在同一时刻执行同一个共享资源。 |
| 原子操作 | 原子操作在进行过程中是不允许中断的；在底层，由CPU提供芯片级别的支持，所以绝对有效。即便在拥有多CPU核心，或者多CPU的计算机系统中，原子操作的保证也是不可撼动的。 |


## 标准库及其数据结构

- （互斥锁）sync.Mutex
  - Lock：阻塞性加锁
  - Unlock：解锁
  - TryLock：非阻塞性加锁
- （读写锁）sync.RWMutex
  - Lock：阻塞性写加锁
  - Unlock：写解锁，会试图唤醒所有因欲进行读锁定而被阻塞的goroutine
  - TryLock
  - RLock：阻塞性读加锁
  - RUnlock：读解锁
- （条件变量）sync.Cond
  - sync.NewCond(l locker)：创建Cond
  - Wait：等待通知。会自动地对与该条件变量关联的那个锁进行解锁，并且使它所在的 goroutine阻塞。一旦接收到通知，该方法所在的goroutine就会被唤醒，并且该方法会立即尝试锁定该锁
  - Signal：单发通知
  - Broadcast：广播
- （只执行一次）sync.Once
  - Do
- sync.WaitGroup
  - Add(delta int)
  - Done()
  - Wait()
- （临时对象池）sync.Pool
  - Put（在当前池存放对象）
  - Get（在当前池获取对象）：一般会先尝试从与本地P对应的那个本地私有池和本地共享池中获取一个对象值。如果获取失败，它就会试图从其他P的本地共享池中偷一个对象值并直接返回给调用方。如果依然未果，它就只能把希望寄托于当前临时对象池的对象值生成函数了
- （并发安全字典）sync.Map
  - LoadAndDelete
  - Delete
  - Store
- （原子操作）atomic
  - 原子操作的运算： 
    - 加法（Add）
    - 比较并交换（compare and swap， CAS）
    - 加载（load）
    - 存储（store）
    - 交换（swap）
  - 适用的数据类型：
    - int32
    - int64
    - uint32
    - uint64
    - uintptr
    - unsafe.Pointer