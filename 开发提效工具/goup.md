# goup

Go语言中的goup工具是一个用于更新Go模块依赖项的命令行工具。它可以自动解决版本冲突并将依赖项升级到其最新版本。

使用goup工具，您可以快速更新所有依赖项或仅更新特定模块的依赖项。此外，goup还支持锁定依赖项的版本，以确保不会在构建时出现意外的依赖项更改。

官网：https://owenou.com/goup/


## 安装
要安装goup工具，可以使用以下命令：
```shell
curl -sSf https://raw.githubusercontent.com/owenthereal/goup/master/install.sh | sh
## 或者
curl -sSf https://raw.githubusercontent.com/owenthereal/goup/master/install.sh | sh -s -- '--skip-prompt'
```


## goup指令
```shell
goup --help
The Go installer

Usage:
  goup [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  install     Install Go with a version
  list        List all installed Go
  remove      Remove Go with a version
  search      Search Go versions to install
  set         Set the default Go version
  upgrade     Upgrade goup
  version     Show goup version

Flags:
  -h, --help      help for goup
  -v, --verbose   Verbose
```