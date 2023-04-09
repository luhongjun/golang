# Go Template模板

Go语言的模板（template）是一种将数据结构与文本/HTML相结合的方式，实现动态生成输出的功能。在Go语言中，我们通过使用`text/template1`或`html/template`包来创建和执行模板。

- 示例
```go
package main

import (
	"html/template"
	"net/http"
)

func main() {
	// 定义模板
	tmpl, err := template.New("index").Parse(`
	<!DOCTYPE html>
	<html>
	<head>
		<title>{{.Title}}</title>
	</head>
	<body>
		<h1>{{.Heading}}</h1>
		<p>{{.Content}}</p>
	</body>
	</html>	
	`)
	if err != nil {
		panic(err)
	}

	// 定义处理函数
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 准备数据
		data := struct {
			Title   string
			Heading string
			Content string
		}{
			Title:   "欢迎来到我的网站",
			Heading: "这是我的网站",
			Content: "欢迎访问我的网站！",
		}

		// 渲染模板
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// 启动服务器
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
```