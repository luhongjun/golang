# Go命令

| 命令 | 概要 |
| --- | --- |
| go | Go程序 |
| go tool cgo | 支持创建可调用C代码的Go包 |
| go tool cover | 创建和分析“go test-coverprofile”生成的覆盖率配置文件 |
| go tool fix | 指定代码包的所有Go语言源码文件中的旧版本代码修正为新版本的代码 |
| fmt | 格式化Go包，与`gofmt`一样的作用 |
| godoc | 给Go包提取注释并生成文档 |
| go tool vet | 检查Go语言源码中静态错误的简单工具 |
| go tool pprof | 用于对Go程序进行性能分析，交互式的访问概要文件的内容 |


| go 命令 | 用法 |概要 |
| --- | --- | --- |
| go bug | `go bug` | 使用默认浏览器向官方发送bug报告 |
| go version | `go version` | 打印golang版本 |
| *go env | `go env [-json] [-u] [-w] [var ...]` | 获取/修改golang的环境变量 |
| go fix | `go fix [packages]` | 封装了"go tool fix"的工具指令，修复 vendor 目录下依赖包的语法 |
| go fmt | `go fmt [-n] [-x] [packages]` | 封装了"gofmt"的工具指令 |
| go generate | `go generate [-run regexp] [-n] [-v] [-x] [build flags] [file.go... | packages]` | 扫描与当前包相关的源代码文件，找出所有包含"//go:generate"的特殊注释，提取并执行该特殊注释后面的命令，命令为可执行程序，形同shell下面执行 |
| go clean | `go clean [clean flags] [build flags] [packages]` | 移除当前源码包和关联源码包里面编译生成的文件 |
| go doc | `go doc <pkg> <sym>[.<methodOrField>]` | 打印附于Go语言程序实体上的文档 |
| go list | `go list [-f format] [-json] [-m] [list flags] [build flags] [packages]` | 列出指定的代码包的信息 / 列出 vendor 目录下的包信息 |
| *go get | `go get [-d] [-t] [-u] [-v] [-insecure] [build flags] [packages]` | 添加包依赖到当前模块并安装依赖，执行后依赖信息会自动添加至 go.mod 文件|
| *go mod | `go mod <command> [arguments]` | 模块依赖 |
| go vet | `go vet [-n] [-x] [-vettool prog] [build flags] [vet flags] [packages]` | 用于检查Go语言源码中静态错误的简单工具 |
| go build | `go build [-o output] [build flags] [packages]` | 编译包与依赖，会编译当前模块 |
| go install | `go install [build flags] [packages]` | 编译并安装 |
| *go run | `go run [build flags] [-exec xprog] package [arguments...]` | 编译并运行 |
| *go test | `go test [build/test flags] [packages] [build/test flags & test binary flags]` | 测试包 |

