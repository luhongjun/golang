# Go Doc

## 参考资料

- [go doc与godoc](https://www.bookstack.cn/read/go_command_tutorial/0.5.md)

## 概要

`go doc`命令可以打印附于Go语言程序实体上的文档。我们可以通过把程序实体的标识符作为该命令的参数来达到查看其文档的目的。

**解释说明：**
所谓Go语言的程序实体，是指变量、常量、函数、结构体以及接口。而程序实体的标识符即是代表它们的名称。标识符又分非限定标识符和限定标识符。

其中，限定标识符一般用于表示某个代码包中的程序实体或者某个结构体类型中的方法或字段。例如，标准库代码包`io`中的名为`EOF`的变量用限定标识符表示即`io.EOF`。又例如，如果我有一个`sync.WaitGroup`类型的变量wg并且想调用它的Add方法，那么可以这样写`wg.Add()`。其中，`wg.Add`就是一个限定标识符，而后面的()则代表了调用操作。

## 指令说明

展示当前包(目录)的文档
```shell
go doc 
```

| 命令参数 | 概要 |
| --- | --- |
| -c | 加入此标记后会使go doc命令区分参数中字母的大小写。默认情况下，命令是大小写不敏感的 |
| -cmd | 加入此标记后会使go doc命令同时打印出main包中的可导出的程序实体（其名称的首字母大写）的文档。默认情况下，这部分文档是不会被打印出来的 |
| -u | 加入此标记后会使go doc命令同时打印出不可导出的程序实体（其名称的首字母小写）的文档。默认情况下，这部分文档是不会被打印出来的 |


## godoc

### 指令说明
usage: `godoc -http=localhost:6060`

| 命令参数 | 概要 |
| --- | --- |
| -analysis string | comma-separated list of analyses to perform when in GOPATH mode (supported: type, pointer). See https://golang.org/lib/godoc/analysis/help.html |
| -goroot string | Go root directory (default "/usr/local/go") |
| -http string | HTTP service address (default "localhost:6060") |
| -index | enable search index |
| -index_files string | glob pattern specifying index files; if not empty, the index is read from these files in sorted order |
| -index_interval duration | interval of indexing; 0 for default (5m), negative to only index once at startup |
| -index_throttle float | index throttle value; 0.0 = no time allocated, 1.0 = full throttle (default 0.75) |
| -links | link identifiers to their declarations (default true) |
| -maxresults int | maximum number of full text search results shown (default 10000) |
| -notes string | regular expression matching note markers to show (default "BUG") |
| -play | enable playground |
| -templates string | load templates/JS/CSS from disk in this directory |
| -timestamps | show timestamps with directory listings |
| -url string | print HTML for named URL |
| -v | verbose mode |
| -write_index | write index to a file; the file name must be specified with -index_files |
| -zip string | zip file providing the file system to serve; disabled if empty |
