# strings - 字符串操作

## 数据类型：byte、rune与字符串

byte，占用1个节字，就 8 个比特位（2^8 = 256，因此 byte 的表示范围 0->255），所以它和 uint8 类型本质上没有区别，它表示的是 ACSII 表中的一个字符。

rune，占用4个字节，共32位比特位，所以它和 int32 本质上也没有区别。它表示的是一个 Unicode字符（Unicode是一个可以表示世界范围内的绝大部分字符的编码规范）。

**注意：**
在 Go 中单引号与 双引号并不是等价的；单引号用来表示字符，如果你使用双引号，就意味着你要定义一个字符串，赋值时与前面声明的会不一致，这样在编译的时候就会出错。

string，可以使用双引号或者反引号来表示字符串，但是两者是有区别的，反引号不会解析待转义字符`\`的文段：
```gotemplate
var mystr01 string = "\\r\\n"   //输出：\r\n
var mystr02 string = "\r\n"     //会解析为换行符
var mystr03 string = `\r\n`     //输出：\r\n
```
同时反引号可以不写换行符（因为没法写）来表示一个多行的字符串:
```gotemplate
var mystr01 string = `你好呀!
我的公众号是: Go编程时光，欢迎大家关注`

fmt.Println(mystr01)
```

## string 函数

| 函数 | 函数签名 | 说明 |
| --- | --- | --- |
| 比较字符串 | func Compare(a, b string) int | 如果两个字符串相等，返回为 0。如果 a 小于 b ，返回 -1 ，反之返回 1 。不推荐使用这个函数，直接使用 == != > < >= <= 等一系列运算符更加直观 |
| 是否相等 | func EqualFold(s, t string) bool | 计算 s 与 t 忽略字母大小写后是否相等 |
| 是否存在某个字符串子串 | func Contains(s, substr string) bool | 子串 substr 在 s 中，返回 true |
| 是否存在某个字符 | func ContainsAny(s, chars string) bool | chars 中任何一个 Unicode 代码点在 s 中，返回 true |
| 是否存在rune | func ContainsRune(s string, r rune) bool | Unicode 代码点 r 在 s 中，返回 true |
| 查找子串出现次数 | func Count(s, sep string) int | 实现的是 Rabin-Karp 算法（计算子串在字符串中出现的无重叠的次数），当 sep 为空时，Count 的返回值是：utf8.RuneCountInString(s) + 1 |
| 拆分字符串为单个字符串数组 | func Fields(s string) []string | 用一个或多个连续的空格分隔字符串 s，返回子字符串的数组（slice） |
| 使用自定义函数拆分字符串为数组 | func FieldsFunc(s string, f func(rune) bool) []string | 通过实现一个回调函数来指定分隔字符串 s 的字符 |
| 分隔字符串 | func Split(s, sep string) []string | 通过 sep 进行分割字符串 |
| 分隔字符串，保留分隔符 | func SplitAfter(s, sep string) []string | 通过 sep 进行分割字符串，并保留分隔符 |
| 限定分隔字符串 | func SplitN(s, sep string, n int) []string | 带 N 的方法可以通过最后一个参数 n 控制返回的结果中的 slice 中的元素个数，当 n < 0 时，返回所有的子字符串；当 n == 0 时，返回的结果是 nil；当 n > 0 时，表示返回的 slice 中最多只有 n 个元素，其中，最后一个元素不会分割 |
| 限定分隔字符串，保留分隔符 | func SplitAfterN(s, sep string, n int) []string | 同上 |
| 是否以字符串为前缀 | func HasPrefix(s, prefix string) bool | s 中是否以 prefix 开始 |
| 是否以字符串为后缀 | func HasSuffix(s, suffix string) bool | s 中是否以 suffix 结尾 |
| 找出第一次出现字符串的位置 | func Index(s, sep string) int | 在 s 中查找 sep 的第一次出现，返回第一次出现的索引 |
| 找出第一次出现字符的位置 | func IndexByte(s string, c byte) int | 在 s 中查找字节 c 的第一次出现，返回第一次出现的索引 |
| 找出任意一个字符出现的位置 | func IndexAny(s, chars string) int | chars 中任何一个 Unicode 代码点在 s 中首次出现的位置 |
| 找出第一次出现自定义函数判断正确的位置 | func IndexFunc(s string, f func(rune) bool) int | 查找字符 c 在 s 中第一次出现的位置，其中 c 满足 f(c) 返回 true |
| 找出第一次出现rune的位置 | func IndexRune(s string, r rune) int | Unicode 代码点 r 在 s 中第一次出现的位置 |
| - | func LastIndex(s, sep string) int | - |
| - | func LastIndexByte(s string, c byte) int | - |
| - | func LastIndexAny(s, chars string) int | - |
| - | func LastIndexFunc(s string, f func(rune) bool) int | - |
| **数组转字符串** | func Join(a []string, sep string) string | 将字符串数组（或 slice）连接起来 |
| 重复字符串 | func Repeat(s string, count int) string | 将 s 重复 count 次，如果 count 为负数或返回值长度 len(s)*count 超出 string 上限会导致 panic |
| 字符替换 | func Map(mapping func(rune) rune, s string) string | Map 函数，将 s 的每一个字符按照 mapping 的规则做映射替换，如果 mapping 返回值 <0 ，则舍弃该字符 |
| 字符串子串替换，限制替换次数 | func Replace(s, old, new string, n int) string | 用 new 替换 s 中的 old，一共替换 n 个。 |
| 字符串子串全部替换 | func ReplaceAll(s, old, new string) string | - |
| 转小写 | func ToLower(s string) string | - |
| 特殊字符转小写 | func ToLowerSpecial(c unicode.SpecialCase, s string) string | - |
| 转大写 | func ToUpper(s string) string | - |
| 特殊字符转大写 | func ToUpperSpecial(c unicode.SpecialCase, s string) string | - |
| 首字母大写 | func Title(s string) string | 将 s 每个单词的首字母大写，不处理该单词的后续字符 |
| 全部大写 | func ToTitle(s string) string | 将 s 的每个字母大写 |
| 特殊字母大写 | func ToTitleSpecial(c unicode.SpecialCase, s string) string | 将 s 的每个字母大写，并且会将一些特殊字母转换为其对应的特殊大写字母。|
| 修剪字符串左右字符 | func Trim(s string, cutset string) string | 将 s 左侧和右侧中匹配 cutset 中的任一字符的字符去掉 |
| 修剪字符串左侧字符 | func TrimLeft(s string, cutset string) string | 将 s 左侧的匹配 cutset 中的任一字符的字符去掉 |
| 修剪字符串右侧字符 | func TrimRight(s string, cutset string) string | 将 s 右侧的匹配 cutset 中的任一字符的字符去掉 |
| 修剪 | func TrimPrefix(s, prefix string) string | 如果 s 的前缀为 prefix 则返回去掉前缀后的 string , 否则 s 没有变化 |
| 修剪 | func TrimSuffix(s, suffix string) string | 如果 s 的后缀为 suffix 则返回去掉后缀后的 string , 否则 s 没有变化 | 
| *修剪 | func TrimSpace(s string) string | 将 s 左侧和右侧的间隔符去掉。常见间隔符包括：'\t', '\n', '\v', '\f', '\r', ' ', U+0085 (NEL) |
| 自定义修剪 | func TrimFunc(s string, f func(rune) bool) string | 将 s 左侧和右侧的匹配 f 的字符去掉 |
| 自定义修剪左侧字符 | func TrimLeftFunc(s string, f func(rune) bool) string | 将 s 左侧的匹配 f 的字符去掉 |
| 自定义修剪右侧字符 | func TrimRightFunc(s string, f func(rune) bool) string | 将 s 右侧的匹配 f 的字符去掉 |


## Replacer 类型

这是一个结构，没有导出任何字段,实例化通过`func NewReplacer(oldnew ...string) *Replacer`函数进行，其中不定参数 oldnew 是 old-new 对，即进行多个替换。如果 oldnew 长度与奇数，会导致 panic.

示例：
```gotemplate
r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
fmt.Println(r.Replace("This is <b>HTML</b>!")) //输出：This is &lt;b&gt;HTML&lt;/b&gt;!
```

## Reader 类型

Reader类型实现了 io 包中的接口，其结构如下：
```gotemplate
type Reader struct {
  s        string    // Reader 读取的数据来源
  i        int // current reading index（当前读的索引位置）
  prevRune int // index of previous rune; or < 0（前一个读取的 rune 索引位置）
}
```

可见 Reader 结构没有导出任何字段，而是提供一个实例化方法：
```gotemplate
func NewReader(s string) *Reader
```

该方法接收一个字符串，返回的 Reader 实例就是从该参数字符串读数据。如果只是为了读取，NewReader 会更高效。

## Builder 类型

该类型实现了 io 包下的 Writer, ByteWriter, StringWriter 等接口，可以向该对象内写入数据，Builder 没有实现 Reader 等接口，所以该类型不可读，但提供了 String 方法可以获取对象内的数据。

结构如下：
```gotemplate
type Builder struct {
    addr *Builder // of receiver, to detect copies by value
    buf  []byte
}
```