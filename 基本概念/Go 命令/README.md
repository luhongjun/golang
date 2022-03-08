# Go命令

| 命令 | 概要 |
| --- | --- |
| go | - |
| cgo | - |
| cover | - |
| fix | - |
| fmt | - |
| godoc | - |
| vet | - |


| go 命令 | 用法 |概要 |
| --- | --- | --- |
| go bug | `go bug` | 使用默认浏览器向官方发送bug报告 |
| go version | `go version` | 打印golang版本 |
| go env | `go env [-json] [-u] [-w] [var ...]` | 获取/修改golang的环境变量 |
| go fix | `go fix [packages]` | 封装了"go tool fix"的工具指令，修复 vendor 目录下依赖包的语法 |
| go fmt | `go fmt [-n] [-x] [packages]` | 封装了"gofmt"的工具指令 |
| go generate | `go generate [-run regexp] [-n] [-v] [-x] [build flags] [file.go... | packages]` | 扫描与当前包相关的源代码文件，找出所有包含"//go:generate"的特殊注释，提取并执行该特殊注释后面的命令，命令为可执行程序，形同shell下面执行 |
| go build | `go build [-o output] [build flags] [packages]` | 编译包与依赖，会编译当前模块 |
| go clean | `go clean [clean flags] [build flags] [packages]` | 移除当前源码包和关联源码包里面编译生成的文件 |
| go doc | `go doc <pkg> <sym>[.<methodOrField>]` | 打印附于Go语言程序实体上的文档 |
| go install | - | compile and install packages and dependencies |
| go list | `go list [-f format] [-json] [-m] [list flags] [build flags] [packages]` | 列出指定的代码包的信息 / 列出 vendor 目录下的包信息 |
| go get | `go get [-d] [-t] [-u] [-v] [-insecure] [build flags] [packages]` | 添加包依赖到当前模块并安装依赖，执行后依赖信息会自动添加至 go.mod 文件|
| go mod | `go mod <command> [arguments]` | 模块依赖 |
| go vet | `go vet [-n] [-x] [-vettool prog] [build flags] [vet flags] [packages]` | 用于检查Go语言源码中静态错误的简单工具 |
| go run | - | compile and run Go program |
| go test | - | test packages |
| go tool | - | run specified go tool |

