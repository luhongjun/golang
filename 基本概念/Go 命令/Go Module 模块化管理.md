# Go Module (go mod指令)

Golang 的包管理一直被大众所诟病的一个点，但是我们可以看到现在确实是在往好的方向进行发展。下面是官方的包管理工具的发展历史：

1. 在 Go1.5 版本之前，所有的依赖包都是存放在 GOPATH 下，没有版本控制。这个类似 Google 使用单一仓库来管理代码的方式。这种方式的最大的弊端就是无法实现包的多版本控制，比如项目 A 和项目 B 依赖于不同版本的 package，如果 package 没有做到完全的向前兼容，往往会导致一些问题。

2. Go1.5 版本推出了 vendor 机制。所谓 vendor 机制，就是每个项目的根目录下可以有一个 vendor 目录，里面存放了该项目的依赖的 package。go build 的时候会先去 vendor 目录查找依赖，如果没有找到会再去 GOPATH 目录下查找

3. Go1.9 版本推出了实验性质的包管理工具 dep，这里把 dep 归结为 Golang 官方的包管理方式可能有一些不太准确。关于 dep 的争议颇多，比如为什么官方后来没有直接使用 dep 而是弄了一个新的 modules，具体细节这里不太方便展开。

4. 1.11 版本推出 modules 机制，简称 mod，它正在开始进行重大改变Go生态系统。它是GOPATH的替代品，集成了版本控制和软件包分发支持。

GOPATH将在1.13版本中弃用，建议后面安装 Golang 1.13或者更高的版本。 Go模块主要解决以下用例：
- 用于获取外部包的依赖关系管理。
- 解析自定义包作为GOPATH的替代。
- 包版本和发布

## go.mod 文件

我们可以通过`go mod init [your_module_name]`可以在当前模块（即当前项目目录下）创建 go.mod 文件，内容如下：
``` 
module your_module_name

go 1.13
```
后面如果开发者在当前模块中有依赖其他模块包，可以再通过执行`go mod tidy`重新更新依赖，此时查看 go.mod 文件，会看到新的内容如下：
``` 
module your_module_name

go 1.13

require rsc.io/quote v1.5.2
```

---------

go.mod 文件一般都不需要开发者手动去编辑，直接通过命令的方式就可以让程序自动地进行修改。但是我们还是有必要清楚 go.mod 文件里面的语法。

关于 go.mod 文件的一些说明可以查看[官网说明-The go.mod file](http://docscn.studygolang.com/cmd/go/#hdr-The_go_mod_file)

- 示例
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

//
//那么【什么时候应该使用replace】呢？ 答：https://github.com/golang/go/wiki/Modules#when-should-i-use-the-replace-directive
replace (
    golang.org/x/text => github.com/golang/text v0.3.0
)
```


## go.sum 文件

资料参阅 
- [Golang 语言中文网 - 谈谈go.sum](https://studygolang.com/articles/25658)
- [go.mod 文件中的indirect准确含义](https://blog.csdn.net/juzipidemimi/article/details/104441398)


另外,我们可以发现当前目录新增了 go.sum 文件，该文件内容如下：
``` 
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c h1:qgOY6WgZOaTkIIMiVjBQcw93ERBE4m30iBm00nkL0i8=
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c/go.mod h1:NqM8EUOU14njkJ3fqMW+pc6Ldnwhi/IjpwHt7yyuwOQ=
rsc.io/quote v1.5.2 h1:w5fcysjrx7yqtD/aO+QwRjYZOKnaM9Uh2b40tElTs3Y=
rsc.io/quote v1.5.2/go.mod h1:LzX7hefJvL54yjefDEDHNONDjII0t9xZLPXsUe+TKr0=
rsc.io/sampler v1.3.0 h1:7uVkIFmeBqHfdjD+gZwtXXI+RODJ2Wc4O7MPEh/QiW4=
rsc.io/sampler v1.3.0/go.mod h1:T1hPZKmBbMNahiBKFy5HrXp6adAjACjK9JXDnKaTXpA=
```
Go会给当前模块依赖的包以及版本打上标记并将相关的唯一值写入 go.sum 中，以防数据被偷偷篡改。

## go mod 

```text
//初始化 modules，会生成一个 go.mod 文件
go mod init [hojun_module_name]
//下载 modules 到本地缓存
go mod download
//提供一种命令行交互修改 go.mod 的方式
go mod edit -replace xxx-rep/uri/xx@v1.2    
//将 module 的依赖图在命令行打印出来
go mod graph
//清理 go.mod 中的依赖，会添加缺失的依赖，同时移除没有用到的依赖
go mod tidy
//将依赖包打包拷贝到项目的 vendor 目录下，值得注意的是并不会将 test code 中的依赖包打包到 vendor 中
go mod vendor
//用来检测依赖包自下载之后是否被改动过
go mod verify
//解释为什么 package 或者 module 是需要
go mod why
```


## 资料参考

- [go mod - 官方文档](http://docscn.studygolang.com/cmd/go/#hdr-Module_maintenance)
- [Go Modules 使用教程](https://www.cnblogs.com/klsw/p/11537850.html)
- [Go Modules详解](https://objcoding.com/2018/09/13/go-modules/)
- [Github golang/go 关于 Module 的 wiki](https://github.com/golang/go/wiki/Modules#go-modules)
- [Golang中文网 - 使用Go Module构建项目](https://studygolang.com/articles/21348?fr=sidebar)
- [关于 Golang 的项目管理](http://golang.iswbm.com/en/latest/chapters/p03.html)
- [何处安放我们的 Go 代码](https://liujiacai.net/blog/2019/10/24/go-modules/)
- [Go Module 终极入门](https://mp.weixin.qq.com/s?__biz=MzUxMDI4MDc1NA==&mid=2247483713&idx=1&sn=817ffef56f8bc5ca09a325c9744e00c7&source=41#wechat_redirect)

### 环境变量

用环境变量 GO111MODULE 开启或关闭模块支持，它有三个可选值：off、on、auto，默认值是 auto。

- GO111MODULE=off 无模块支持，go 会从 GOPATH 和 vendor 文件夹寻找包。
- GO111MODULE=on 模块支持，go 会忽略 GOPATH 和 vendor 文件夹，只根据 go.mod 下载依赖。
- GO111MODULE=auto 在 $GOPATH/src 外面且根目录有 go.mod 文件时，开启模块支持。
