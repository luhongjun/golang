# Unicode 码点、UTF-8/16 编码

世界中的字符有许许多多，有英文，中文，韩文等。我们强烈需要一个大大的映射表把世界上的字符映射成计算机可以阅读的二进制数字（字节）。 这样，每个字符都给予一个独一无二的编码，就不会出现写文字的人和阅读文字的人编码不同而出现无法读取的乱码现象了。

于是 Unicode 就出现了，它是一种所有符号的编码映射。最开始的时候，unicode 认为使用两个字节，也就是 16 位就能包含所有的字符了。 但是非常可惜，两个字节最多只能覆盖 65536 个字符，这显然是不够的，于是附加了一套字符编码，即 unicode4.0，附加的字符用 4 个字节表示。 现在为止，大概 Unicode 可以覆盖 100 多万个字符了。

----------------

Unicode 只是定义了一个字符和一个编码的映射，但是呢，对应的存储却没有制定。 比如一个编码 0x0041 代表大写字母 A，那么可能有一种存储至少有 4 个字节，那可能 0x00000041 来存储代表 A。 这个就是 unicode 的具体实现。unicode 的具体实现有很多种，UTF-8 和 UTF-16 就是其中两种。

- UTF-8 表示最少用一个字节就能表示一个字符的编码实现。它采取的方式是对不同的语言使用不同的方法，将 unicode 编码按照这个方法进行转换。 我们只要记住最后的结果是**英文占一个字节，中文占三个字节**。这种编码实现方式也是现在应用最为广泛的方式了。

- UTF-16 表示最少用两个字节能表示一个字符的编码实现。同样是对 unicode 编码进行转换，它的结果是英文占用两个字节，中文占用两个或者四个字节。 实际上，UTF-16 就是最严格实现了 unicode4.0。但由于英文是最通用的语言，所以推广程度没有 UTF-8 那么普及。

-----------

回到 go 对 unicode 包的支持，由于 UTF-8 的作者 Ken Thompson 同时也是 go 语言的创始人，所以说，在字符支持方面，几乎没有语言的理解会高于 go 了。 go 对 unicode 的支持包含三个包 :

- unicode: unicode 包包含基本的字符判断函数
- unicode/utf8: 主要负责 rune 和 byte 之间的转换
- unicode/utf16: 负责 rune 和 uint16 数组之间的转换

由于字符的概念有的时候比较模糊，比如字符（小写 a）普通显示为 a，在重音字符中（grave-accented）中显示为à。 这时候字符（character）的概念就有点不准确了，因为 a 和à显然是两个不同的 unicode 编码，但是却代表同一个字符，所以引入了 rune。 一个 rune 就代表一个 unicode 编码，所以上面的 a 和à是两个不同的 rune。

这里有个要注意的事情，go 语言的所有代码都是 UTF8 的，所以如果我们在程序中的字符串都是 utf8 编码的，但是我们的单个字符（单引号扩起来的）却是 unicode 的。

## unicode 包

包函数有如下：
```gotemplate
func IsControl(r rune) bool  // 是否控制字符
func IsDigit(r rune) bool  // 是否阿拉伯数字字符，即 0-9
func IsGraphic(r rune) bool // 是否图形字符
func IsLetter(r rune) bool // 是否字母
func IsLower(r rune) bool // 是否小写字符
func IsMark(r rune) bool // 是否符号字符
func IsNumber(r rune) bool // 是否数字字符，比如罗马数字Ⅷ也是数字字符
func IsOneOf(ranges []*RangeTable, r rune) bool // 是否是 RangeTable 中的一个
func IsPrint(r rune) bool // 是否可打印字符
func IsPunct(r rune) bool // 是否标点符号
func IsSpace(r rune) bool // 是否空格
func IsSymbol(r rune) bool // 是否符号字符
func IsTitle(r rune) bool // 是否 title case
func IsUpper(r rune) bool // 是否大写字符
func Is(rangeTab *RangeTable, r rune) bool // r 是否为 rangeTab 类型的字符
func In(r rune, ranges ...*RangeTable) bool  // r 是否为 ranges 中任意一个类型的字符
```

## utf8 包

utf8 里面的函数就有一些字节和字符的转换。

判断是否符合 utf8 编码的函数：
```gotemplate
func Valid(p []byte) bool
func ValidRune(r rune) bool
func ValidString(s string) bool
```

判断 rune 所占字节数：
```gotemplate
func RuneLen(r rune) int
```

判断字节串或者字符串的 rune 数：
```gotemplate
func RuneCount(p []byte) int
func RuneCountInString(s string) (n int)
```

编码和解码到 rune：
```gotemplate
func EncodeRune(p []byte, r rune) int
func DecodeRune(p []byte) (r rune, size int)
func DecodeRuneInString(s string) (r rune, size int)
func DecodeLastRune(p []byte) (r rune, size int)
func DecodeLastRuneInString(s string) (r rune, size int)
```

是否为完整 rune：
```gotemplate
func FullRune(p []byte) bool
func FullRuneInString(s string) bool
```

是否为 rune 第一个字节：
```gotemplate
func RuneStart(b byte) bool
```


## utf16 包

略...