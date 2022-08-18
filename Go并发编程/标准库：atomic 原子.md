# 标准库： atomic 原子

官方文档：https://pkg.go.dev/sync/atomic

sync/atomic
原子操作的运算： 加法（Add）、比较并交换（compare and swap， CAS）、加载（load）、存储（store）、交换（swap）
适用的数据类型：int32、int64、uint32、uint64、uintptr、unsafe.Pointer
atomic.Value：Store / load