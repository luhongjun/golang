##  Go mod指令

- [go mod - 官方文档](http://docscn.studygolang.com/cmd/go/#hdr-Module_maintenance)
- [Go Modules 不完全教程](https://www.cnblogs.com/klsw/p/11537850.html)
- [Go Modules详解](https://objcoding.com/2018/09/13/go-modules/)

### 子命令
```text
//下载 modules 到本地缓存
download    download modules to local cache
//提供一种命令行交互修改 go.mod 的方式
edit        edit go.mod from tools or scripts
//将 module 的依赖图在命令行打印出来
graph       print module requirement graph
//初始化 modules，会生成一个 go.mod 文件
init        initialize new module in current directory
//清理 go.mod 中的依赖，会添加缺失的依赖，同时移除没有用到的依赖
tidy        add missing and remove unused modules
//将依赖包打包拷贝到项目的 vendor 目录下，值得注意的是并不会将 test code 中的依赖包打包到 vendor 中
vendor      make vendored copy of dependencies
//用来检测依赖包自下载之后是否被改动过
verify      verify dependencies have expected content
//解释为什么 package 或者 module 是需要
why         explain why packages or modules are needed
```

### 包管理历史

Golang 的包管理一直被大众所诟病的一个点，但是我们可以看到现在确实是在往好的方向进行发展。下面是官方的包管理工具的发展历史：

- 在 1.5 版本之前，所有的依赖包都是存放在 GOPATH 下，没有版本控制。这个类似 Google 使用单一仓库来管理代码的方式。这种方式的最大的弊端就是无法实现包的多版本控制，比如项目 A 和项目 B 依赖于不同版本的 package，如果 package 没有做到完全的向前兼容，往往会导致一些问题。

- 1.5 版本推出了 vendor 机制。所谓 vendor 机制，就是每个项目的根目录下可以有一个 vendor 目录，里面存放了该项目的依赖的 package。go build 的时候会先去 vendor 目录查找依赖，如果没有找到会再去 GOPATH 目录下查找

- 1.9 版本推出了实验性质的包管理工具 dep，这里把 dep 归结为 Golang 官方的包管理方式可能有一些不太准确。关于 dep 的争议颇多，比如为什么官方后来没有直接使用 dep 而是弄了一个新的 modules，具体细节这里不太方便展开。

- 1.11 版本推出 modules 机制，简称 mod，也就是本文要讨论的重点。

### 环境变量

用环境变量 GO111MODULE 开启或关闭模块支持，它有三个可选值：off、on、auto，默认值是 auto。

- GO111MODULE=off 无模块支持，go 会从 GOPATH 和 vendor 文件夹寻找包。
- GO111MODULE=on 模块支持，go 会忽略 GOPATH 和 vendor 文件夹，只根据 go.mod 下载依赖。
- GO111MODULE=auto 在 $GOPATH/src 外面且根目录有 go.mod 文件时，开启模块支持。

### go.mod 文件

- [go mod文件 - 官网](http://docscn.studygolang.com/cmd/go/#hdr-The_go_mod_file)

```
//可以通过 go mod init example.com/m 生成go.mod文件的同时，也定义模块路径
//模块路径指模块根目录的导入路径，也是其他子目录导入路径的前缀
module example.com/m

//表明我们期望的go语言版本
go 1.14

//表明当前模块所需要的 给定版本或更高版本的特定模块
require (
    golang.org/x/text v0.3.0
    gopkg.in/yaml.v2 v2.1.0 
)

//表明当前模块需要排除指定的模块
exclude (
    old/thing v1.2.3
)

replace (
    golang.org/x/text => github.com/golang/text v0.3.0
)
```