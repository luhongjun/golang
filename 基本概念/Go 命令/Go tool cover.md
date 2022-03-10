# Go tool cover

我们在对go文件进行测试的时候，会输出一些文件报告
```shell
go test -coverprofile=c.out
```

此时，我们可以通过cover工具来友好地展示单元测试的覆盖率
```shell
go tool cover -html=coverprofile.cov
go tool cover -func=coverprofile.cov
```

## 指令说明

| 指令参数 | 概要 |
| --- | --- |
| -V | 输出go tool cover 的版本并结束 | 
| -func [string] | 输出每个函数的代码覆盖率报告 | 
| -html [string] | 以HTML格式生成代码覆盖率报告 |
| -mode [string] | 覆盖模式：set，count，atomic |
| -o [string] | 输出为文件形式，默认是终端/控制台输出 |
| -var [string] | 生成的覆盖率变量名称，默认值：GoCover |
