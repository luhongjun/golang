# iota 类型

iota常用于const表达式中，我们还知道其值是从零开始，const声明块中每增加一行iota值自增1

## 示例

以下示例便于我们理解 iota 用法：
```gotemplate
//代码取日志模块，定义了一组代表日志级别的常量，常量类型为Priority，实际为int类型
type Priority int
const (
    LOG_EMERG Priority = iota  //0
    LOG_ALERT   //1
    LOG_CRIT    //2
    LOG_ERR     //3
    LOG_WARNING //4
    LOG_NOTICE  //5
    LOG_INFO    //6
    LOG_DEBUG   //7
)
```

```gotemplate
//代码取自Go互斥锁Mutex的实现，用于指示各种状态位的地址偏移
const (
    mutexLocked = 1 << iota // mutex is locked
    mutexWoken  //1 << 1 = 2
    mutexStarving   //1 << 2 = 4
    mutexWaiterShift = iota //3
    starvationThresholdNs = 1e6 //1e+06
)
```

```gotemplate
const (
    bit0, mask0 = 1 << iota, 1<<iota - 1  //1, 0
    bit1, mask1 //2, 1
    _, _
    bit3, mask3 //8, 7
)
```

## 编译原理

iota代表了const声明块的行索引（下标从0开始）

const块中每一行在GO中使用spec数据结构描述，spec声明如下：
```gotemplate
// A ValueSpec node represents a constant or variable declaration
// (ConstSpec or VarSpec production).
//
ValueSpec struct {
    Doc     *CommentGroup // associated documentation; or nil
    Names   []*Ident      // value names (len(Names) > 0)
    Type    Expr          // value type; or nil
    Values  []Expr        // initial values; or nil
    Comment *CommentGroup // line comments; or nil
}
```
这里我们只关注ValueSpec.Names， 这个切片中保存了一行中定义的常量，如果一行定义N个常量，那么ValueSpec.Names切片长度即为N。

const块实际上是spec类型的切片，用于表示const中的多行。


所以编译期间构造常量时的伪算法如下：
``` 
for iota, spec := range ValueSpecs {
    for i, name := range spec.Names {
        obj := NewConst(name, iota...) //此处将iota传入，用于构造常量
        ...
    }
}
```
从上面可以更清晰的看出iota实际上是遍历const块的索引，每行中即便多次使用iota，其值也不会递增。