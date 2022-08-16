# go test 工作机制

Go 有多个命令行工具，go test只是其中一个。go test命令的函数入口在`src\cmd\go\internal\test\test.go:runTest()`，这个函数就是go test的大脑

runTest()函数场景如下：
```gotemplate
func runTest(cmd *base.Command, args []string)
```
GO 命令行工具的实现中，都遵循这种函数声明，其中args即命令行输入的全部参数。

runTest首先会分析所有需要测试的包，为每个待测包生成一个二进制文件，然后执行。