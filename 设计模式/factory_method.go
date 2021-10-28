package main

import "io"

// 定义存储类型，告诉工厂需要创建哪种类型的存储类
type StorageType int
const (
	DiskStorage StorageType = 1 << iota
	TempStorage
	MemoryStorage
)

// 定义一个接口（不同的存储类应该要继承这个打开的方法）
type Store interface {
	Open(string) (io.ReadWriteCloser, error)
}

/** 实现类 - 内存存储类 **/
type newMemoryStorage struct {
	Store
}
func (*newMemoryStorage) Open(fileName string) (io.ReadWriteCloser, error) {
	return nil,nil
}

/** 实现类 - 磁盘存储类 **/
type newDiskStorage struct {
	Store
}
func (*newDiskStorage) Open(fileName string) (io.ReadWriteCloser, error) {
	return nil,nil
}

/** 实现类 - 临时性存储类 **/
type newTempStorage struct {
	Store
}
func (*newTempStorage) Open(fileName string) (io.ReadWriteCloser, error) {
	return nil,nil
}

//
func NewStore(t StorageType) Store {
	switch t {
		case MemoryStorage:
			return &newMemoryStorage{}
		case DiskStorage:
			return &newDiskStorage{}
		case TempStorage:
			return &newTempStorage{}
		default:
			panic("error storage type")
	}
}

// 使用示例
func main()  {
	tempStorage := NewStore(TempStorage)
	_, _ = tempStorage.Open("liu_de_hua.png")

	memoryStorage := NewStore(MemoryStorage)
	_, _ = memoryStorage.Open("file")
}