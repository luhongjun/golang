# Go 常用的命令

Go 详细的命令见[官网](http://docscn.studygolang.com/cmd/go/)

- [build](http://docscn.studygolang.com/cmd/go/#hdr-Compile_packages_and_dependencies): 编译包和依赖

go build 在编译开始时，会搜索当前目录的 go 源码。这个例子中，go build 会找到指定的go文件。编译go文件后，生成当前目录名的可执行文件并放置于当前目录下。
```text
go build a.go b.go //可以同时编译多个文件
go build my_package//编译某个包
go build -n //打印编译时会用到的所有命令，但不真正执行
go build -x //打印编译时会用到的所有命令
go build -v //编译时显示包名
```

- [clean](http://docscn.studygolang.com/cmd/go/#hdr-Remove_object_files_and_cached_files): 移除对象文件

移除当前源码包和关联源码包里面编译生成的文件

- doc: 显示包或者符号的文档
- env: 打印go的环境信息
- bug: 启动错误报告
- fix: 运行go tool fix
- fmt: 运行gofmt进行格式化
- generate: 从processing source生成go文件
- get: 下载并安装包和依赖
- install: 编译并安装包和依赖
- list: 列出包

- run: 编译并运行go程序
- test: 运行测试
- tool: 运行go提供的工具
- version: 显示go的版本
- vet: 运行go tool vet