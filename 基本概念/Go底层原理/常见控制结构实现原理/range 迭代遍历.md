# range 迭代遍历

```gotemplate
func RangeSlice(slice []int) {
    for index, value := range slice {
        _, _ = index, value
    }
}
```
遍历过程中每次迭代会对index和value进行赋值，如果数据量大或者value类型为string时，对value的赋值操作可能是多余的，可以在for-range中忽略value值，使用slice[index]引用value值。
```gotemplate
func RangeSlice(slice []int) {
    for index, _ := range slice {
        value := slice[index]
    }
}
```

--------------------

```gotemplate
func RangeMap(myMap map[int]string) {
    for key, _ := range myMap {
        _, _ = key, myMap[key]
    }
}
```
函数中for-range语句中只获取key值，然后跟据key值获取value值，虽然看似减少了一次赋值，但通过key值查找value值的性能消耗可能高于赋值消耗。能否优化取决于map所存储数据结构特征、结合实际情况进行。

--------------------

请问如下程序是否能正常结束？
```gotemplate
func main() {
    v := []int{1, 2, 3}
    for i:= range v {
        v = append(v, i)
    }
}
```
能够正常结束。循环内改变切片的长度，不影响循环次数，循环次效在循环开始前就已经确定了。

## range 实现原理

编译器源码gofrontend/go/statements.cc/For_range_statement::do_lower()方法中有如下注释：
```gotemplate
// Arrange to do a loop appropriate for the type.  We will produce
//   for INIT ; COND ; POST {
//           ITER_INIT
//           INDEX = INDEX_TEMP
//           VALUE = VALUE_TEMP // If there is a value
//           original statements
//   }
```
可见range实际上是一个C风格的循环结构。range支持数组、数组指针、切片、map和channel类型，对于不同类型有些细节上的差异。

- range for slice: 如果遍历过程中新增元素是遍历不到的
- range for map: 遍历过程中新增元素不一定会遍历到
- range for channel: 遍历过程中新增元素是会遍历到的，但是如果没有元素，会一直阻塞着

### range for slice

下面的注释解释了遍历slice的过程：
``` 
// The loop we generate:
//   for_temp := range
//   len_temp := len(for_temp)
//   for index_temp = 0; index_temp < len_temp; index_temp++ {
//           value_temp = for_temp[index_temp]
//           index = index_temp
//           value = value_temp
//           original body
//   }
```
遍历slice前会先获以slice的长度len_temp作为循环次数，循环体中，每次循环会先获取元素值，如果for-range中接收index和value的话，则会对index和value进行一次赋值。

由于循环开始前循环次数就已经确定了，所以循环过程中新添加的元素是没办法遍历到的。

### range for map

下面的注释解释了遍历map的过程：
``` 
// The loop we generate:
//   var hiter map_iteration_struct
//   for mapiterinit(type, range, &hiter); hiter.key != nil; mapiternext(&hiter) {
//           index_temp = *hiter.key
//           value_temp = *hiter.val
//           index = index_temp
//           value = value_temp
//           original body
//   }
```
遍历map时没有指定循环次数，循环体与遍历slice类似。由于map底层实现与slice不同，map底层使用hash表实现，插入数据位置是随机的，所以遍历过程中新插入的数据不能保证遍历到。


### range for channel

遍历channel是最特殊的，这是由channel的实现机制决定的：
``` 
// The loop we generate:
//   for {
//           index_temp, ok_temp = <-range
//           if !ok_temp {
//                   break
//           }
//           index = index_temp
//           original body
//   }
```
channel遍历是依次从channel中读取数据,读取前是不知道里面有多少个元素的。如果channel中没有元素，则会阻塞等待，如果channel已被关闭，则会解除阻塞并退出循环。

**注：**
- 上述注释中index_temp实际上描述是有误的，应该为value_temp，因为index对于channel是没有意义的。
- 使用for-range遍历channel时只能获取一个返回值。

## 编程Tip

- 遍历过程中可以适情况放弃接收index或value，可以一定程度上提升性能
- 遍历channel时，如果channel中没有数据，可能会阻塞
- 尽量避免遍历过程中修改原数据
- 使用index,value接收range返回值会发生一次数据拷贝