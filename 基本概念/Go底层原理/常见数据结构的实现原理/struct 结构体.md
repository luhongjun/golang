# struct 结构体

Go是如何管理struct字段的：
```gotemplate
// A StructField describes a single field in a struct.
type StructField struct {
    // Name is the field name.
    Name string
    ...
    Type      Type      // field type
    Tag       StructTag // field tag string
    ...
}

type StructTag string
```

## Tag 的本质

Go的struct声明允许字段附带Tag来对字段做一些标记。

该Tag不仅仅是一个字符串那么简单，因为其主要用于反射场景，reflect包中提供了操作Tag的方法，所以Tag写法也要遵循一定的规则。

### Tag 规则

Tag本身是一个字符串，但字符串中却是：以空格分隔的 key:value 对。

- key: 必须是非空字符串，字符串不能包含控制字符、空格、引号、冒号。
- value: 以双引号标记的字符串

注意：冒号前后不能有空格

如下代码所示，如此写没有实际意义，仅用于说明Tag规则
```gotemplate
type Server struct {
    ServerName string `key1: "value1" key11:"value11"`
    ServerIP   string `key2: "value2"`
}
```
上述代码ServerName字段的Tag包含两个key-value对。ServerIP字段的Tag只包含一个key-value对。

### 获取 Tag

StructTag提供了Get(key string) string方法来获取Tag，示例如下：
```gotemplate
package main

import (
    "reflect"
    "fmt"
)

type Server struct {
    ServerName string `key1:"value1" key11:"value11"`
    ServerIP   string `key2:"value2"`
}

func main() {
    s := Server{}
    st := reflect.TypeOf(s)

    field1 := st.Field(0)
    fmt.Printf("key1:%v\n", field1.Tag.Get("key1"))
    fmt.Printf("key11:%v\n", field1.Tag.Get("key11"))

    filed2 := st.Field(1)
    fmt.Printf("key2:%v\n", filed2.Tag.Get("key2"))
}
```
输出：
``` 
key1:value1
key11:value11
key2:value2
```

--------------------

### Tag存在的意义

本文示例中tag没有任何实际意义，这是为了阐述tag的定义与操作方法，也为了避免与你之前见过的诸如json:xxx混淆。

使用反射可以动态的给结构体成员赋值，正是因为有tag，在赋值前可以使用tag来决定赋值的动作。 比如，官方的encoding/json包，可以将一个JSON数据Unmarshal进一个结构体，此过程中就使用了Tag. 该包定义一些规则，只要参考该规则设置tag就可以将不同的JSON数据转换成结构体。

总之：正是基于struct的tag特性，才有了诸如json、orm等等的应用。理解这个关系是至关重要的。或许，你可以定义另一种tag规则，来处理你特有的数据。

## Tag常见用法

- JSON数据解析`encoding/json`
``` 
json.Marshal：https://studygolang.com/static/pkgdoc/pkg/encoding_json.htm#Marshal
json.Unmarshal：https://studygolang.com/static/pkgdoc/pkg/encoding_json.htm#Unmarshal
```

- ORM映射
```gotemplate
//详见：https://gorm.io/docs/models.html#%E5%AD%97%E6%AE%B5%E6%A0%87%E7%AD%BE
type ThirdRecommendFlow struct {
	ProjectId                 string `gorm:"column:project_id"`
	RecommendId               string `gorm:"column:recommend_id"`
}
```