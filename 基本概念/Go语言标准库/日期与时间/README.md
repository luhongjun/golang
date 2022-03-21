# 日期与时间

`time` 包提供了时间的显示和计量用的功能。日历的计算采用的是公历;

主要提供一下主要类型：

| 主要类型 | 说明 |
| --- | --- |
| Location | 代表一个地区，并表示该地区所在的时区（可能多个）。Location 通常代表地理位置的偏移，比如 CEST 和 CET 表示中欧 |
| Time | 代表一个纳秒精度的时间点，是公历时间 |
| Duration | 代表两个时间点之间经过的时间，以纳秒为单位。可表示的最长时间段大约 290 年，也就是说如果两个时间点相差超过 290 年，会返回 290 年，也就是 minDuration(-1 << 63) 或 maxDuration(1 << 63 - 1) |
| Timer与Ticker | 定时器相关类型 |
| Weekday与Month | 这两个类型的原始类型都是 int，定义它们，语义更明确，同时，实现 fmt.Stringer 接口，方便输出 |

一些比较常用的函数：

| 函数签名 | 说明 |
| --- | --- | --- |
| func LoadLocation(name string) (*Location, error) | (a)、根据名称获取特定时区的实例； (b)、LoadLocation 函数需要的时区数据库可能不是所有系统都提供，特别是非 Unix 系统。此时 LoadLocation 会查找环境变量 ZONEINFO 指定目录或解压该变量指定的 zip 文件（如果有该环境变量）；然后查找 Unix 系统的惯例时区数据安装位置，最后查找 $GOROOT/lib/time/zoneinfo.zip； (c)、通常我们使用`time.Local`即可获取当前系统时区 |
| func Now() Time | 获取当前时间 |
| func (t Time) IsZero() bool | 判断 Time 表示的时间是否是 0 值 |
| func Unix(sec int64, nsec int64) Time | 通过 Unix 时间戳生成 time.Time 实例 |
| time.Time.Unix() | 得到 Unix 时间戳 |
| ime.Time.UnixNano() | 得到 Unix 时间戳的纳秒表示 |
| time.Parse | time.Now() 的时区是 time.Local，而 time.Parse 解析出来的时区却是 time.UTC（可以通过 Time.Location() 函数知道是哪个时区）。在中国，它们相差 8 小时。 |
| time.Time().Truncate

```gotemplate

time = time.Time()
// 向下取整
time.Truncate(time.Hour)
//
time.Round(
```

## 时区

**问题：Local是如何做到表示本地时区的？**

时区信息既浩繁又多变，Unix 系统以标准格式存于文件中，这些文件位于 `/usr/share/zoneinfo`，而本地时区可以通过 `/etc/localtime` 获取，这是一个符号链接，指向 /usr/share/zoneinfo 中某一个时区。比如我本地电脑指向的是：/usr/share/zoneinfo/Asia/Shanghai。

因此，在初始化 Local 时，通过读取 /etc/localtime 可以获取到系统本地时区


## Time 类型详解

程序中应使用 Time 类型值来保存和传递时间，而不是指针。就是说，表示时间的变量和字段，应为 time.Time 类型，而不是 *time.Time. 类型。

Time 零值代表时间点 January 1, year 1, 00:00:00.000000000 UTC。因为本时间点一般不会出现在使用中，IsZero 方法提供了检验时间是否是显式初始化的一个简单途径。

一个 Time 类型值可以被多个 go 协程同时使用。时间点可以使用 Before、After 和 Equal 方法进行比较。Sub 方法让两个时间点相减，生成一个 Duration 类型值（代表时间段）。Add 方法给一个时间点加上一个时间段，生成一个新的 Time 类型时间点。

每一个 Time 都具有一个地点信息（即对应地点的时区信息），当计算时间的表示格式时，如 Format、Hour 和 Year 等方法，都会考虑该信息。Local、UTC 和 In 方法返回一个指定时区（但指向同一时间点）的 Time。修改地点 / 时区信息只是会改变其表示；不会修改被表示的时间点，因此也不会影响其计算。

通过 == 比较 Time 时，Location 信息也会参与比较，因此 Time 不应该作为 map 的 key。

- Time 的内部结构
```gotemplate
type Time struct {
    // sec gives the number of seconds elapsed since
    // January 1, year 1 00:00:00 UTC.
    sec int64

    // nsec specifies a non-negative nanosecond
    // offset within the second named by Seconds.
    // It must be in the range [0, 999999999].
    nsec int32

    // loc specifies the Location that should be used to
    // determine the minute, hour, month, day, and year
    // that correspond to this Time.
    // Only the zero Time has a nil Location.
    // In that case it is interpreted to mean UTC.
    loc *Location
}
```

要讲解 time.Time 的内部结构，得先看 time.Now() 函数。
```gotemplate
// Now returns the current local time.
func Now() Time {
    sec, nsec := now()
    return Time{sec + unixToInternal, nsec, Local}
}
```

**重点来罗！！**

now() 的具体实现在 runtime 包中，以 linux/amd64 为例，在 sys_linux_amd64.s 中的 `time · now`，这是汇编实现的：
- 调用系统调用 `clock_gettime` 获取时钟值（这是 POSIX 时钟）。其中 clockid_t 时钟类型是 CLOCK_REALTIME，也就是可设定的系统级实时时钟。得到的是 struct timespec 类型。（可以到纳秒）
- 如果 clock_gettime 不存在，则使用精度差些的系统调用 gettimeofday。得到的是 struct timeval 类型。（最多到微秒）

注意： 这里使用了 Linux 的 vdso 特性，不了解的，可以查阅相关知识。

虽然 timespec 和 timeval 不一样，但结构类似。因为 now() 函数返回两个值：sec( 秒 ) 和 nsec( 纳秒 )，所以，time · now 的实现将这两个结构转为需要的返回值。需要注意的是，Linux 系统调用返回的 sec( 秒 ) 是 Unix 时间戳，也就是从 1970-1-1 算起的。

回到 time.Now() 的实现，现在我们得到了 sec 和 nsec，从 Time{sec + unixToInternal, nsec, Local} 这句可以看出，Time 结构的 sec 并非 Unix 时间戳，实际上，加上的 unixToInternal 是 1-1-1 到 1970-1-1 经历的秒数。也就是 Time 中的 sec 是从 1-1-1 算起的秒数，而不是 Unix 时间戳。

Time 的最后一个字段表示地点时区信息。本章后面会专门介绍。




