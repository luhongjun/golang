# go test 指令详解

go test有非常丰富的参数，一些参数用于控制测试的编译，另一些参数控制测试的执行

## 控制编译的参数

| 参数 | 说明 | 示例 |
| --- | --- | --- |
| -args | -args后面可以附带多个参数，所有参数都将以字符串形式传入，每个参数做为一个string，并存放到字符串切片中 | go test -run TestArgs -v -args "cloud" |
| -json | 指示go test将结果输出转换成json格式，以方便自动化测试解析使用 | - |
| -o | 指定生成的二进制可执行程序，并执行测试，测试结束不会删除该程序；没有此参数时，go test生成的二进制可执行程序存放到临时目录，执行结束便删除 | go test -run TestAdd -o TestAdd |

## 控制测试的参数

| 参数 | 说明 | 示例 |
| --- | --- | --- |
| -bench regexp | go test默认不执行性能测试，使用-bench参数才可以运行，而且只运行性能测试函数 | go test -bench . |
| -benchtime s | 指定每个性能测试的执行时间，如果不指定，则使用默认时间1s | go test -bench Sub/A=1 -benchtime 2s |
| -cpu 1,2,4 | 提供一个CPU个数的列表，提供此列表后，那么测试将按照这个列表指定的CPU数设置GOMAXPROCS并分别测试 | go test -bench Sub/A=1 -cpu 1,2 (表示每个测试将执行两次，一次是用1个CPU执行，一次是用2个CPU执行) |
| -count n | 指定每个测试执行的次数，默认执行一次 | go test -bench Sub/A=1 -count 2 | 
| -failfast | 如果有测试出现失败，则立即停止测试 | go test -failfast |
| -list regexp | 列出匹配成功的测试函数，并不真正执行。而且，不会列出子函数 | go test -list Sub |
| -parallel n | 测试的最大并发数。当测试使用t.Parallel()方法将测试转为并发时，将受到最大并发数的限制，默认情况下最多有GOMAXPROCS个测试并发，其他的测试只能阻塞等待 | - |
| -run regexp | 跟据正则表达式执行单元测试和示例测试。正则匹配规则与-bench 类似。 | - |
| -timeout d | 默认情况下，测试执行超过10分钟就会超时而退出。 | go test -timeout=1s/5xm/1xh |
| -v | 默认情况下，测试结果只打印简单的测试结果，-v 参数可以打印详细的日志 | - |
| -benchmem | 默认情况下，性能测试结果只打印运行次数、每个操作耗时。使用-benchmem则可以打印每个操作分配的字节数、每个操作分配的对象数 | 
