## go install

## 参考资料

- [go install命令——编译并安装](http://c.biancheng.net/view/122.html)

## 概要

go install 只是将编译的中间文件放在 GOPATH 的 pkg 目录下，以及固定地将编译结果放在 GOPATH 的 bin 目录下。

这个命令在内部实际上分成了两步操作：
1. 生成结果文件（可执行文件或者 .a 包）;
2. 第二步会把编译好的结果移到 $GOPATH/pkg 或者 $GOPATH/bin;

**编译过程说明**
- go install 是建立在 GOPATH 上的，无法在独立的目录里使用 go install。
- GOPATH 下的 bin 目录放置的是使用 go install 生成的可执行文件，可执行文件的名称来自于编译时的包名。
- go install 输出目录始终为 GOPATH 下的 bin 目录，无法使用-o附加参数进行自定义。
- GOPATH 下的 pkg 目录放置的是编译期间的中间文件。