# Go1.18新特性

Go 1.18 是一个大型版本，其中包括新功能、性能改进以及对该语言的最大更改，具体来说，包含以下几大特性：
- 泛型
- 模糊测试（Fuzzing）：使用随机输入来测试程序功能的测试技术，是对常规测试手段的重要补充，可发现某些在人工生成的输入数据下难以发现的Bug

- 工作空间（Workspaces）
- 20% 性能提升：

### 模糊测试(Fuzzing)



```go
// app.go
package main

import "fmt"

func Reverse(s string) string {
    b := []byte(s)
    for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
        b[i], b[j] = b[j], b[i]
    }
    return string(b)
}

func main() {
    input := "The quick brown fox jumped over the lazy dog"
    rev := Reverse(input)
    doubleRev := Reverse(rev)
    fmt.Printf("original: %q\n", input)
    fmt.Printf("reversed: %q\n", rev)
    fmt.Printf("reversed again: %q\n", doubleRev)
}
```

测试代码：
```go
package main

import (
    "testing"
)

// 常规测试函数
func TestReverse(t *testing.T) {
    testcases := []struct {
        in, want string
    }{
        {"Hello, world", "dlrow ,olleH"},
        {" ", " "},
        {"!12345", "54321!"},
    }
    for _, tc := range testcases {
        rev := Reverse(tc.in)
        if rev != tc.want {
                t.Errorf("Reverse: %q, want %q", rev, tc.want)
        }
    }
}

// 模糊测试函数
func FuzzReverse(f *testing.F) {
    testcases := []string{"Hello, world", " ", "!12345"}
    for _, tc := range testcases {
        f.Add(tc)  // Use f.Add to provide a seed corpus
    }
    f.Fuzz(func(t *testing.T, orig string) {
        rev := Reverse(orig)
        doubleRev := Reverse(rev)
        if orig != doubleRev {
            t.Errorf("Before: %q, after: %q", orig, doubleRev)
        }
        if utf8.ValidString(orig) && !utf8.ValidString(rev) {
            t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
        }
    })
}
```

测试结果如下：
```shell
$ go test -fuzz=Fuzz
fuzz: elapsed: 0s, gathering baseline coverage: 0/3 completed
fuzz: elapsed: 0s, gathering baseline coverage: 3/3 completed, now fuzzing with 8 workers
fuzz: minimizing 38-byte failing input file...
--- FAIL: FuzzReverse (0.01s)
    --- FAIL: FuzzReverse (0.00s)
        reverse_test.go:20: Reverse produced invalid UTF-8 string "\x9c\xdd"

    Failing input written to testdata/fuzz/FuzzReverse/af69258a12129d6cbba438df5d5f25ba0ec050461c116f777e77ea7c9a0d217a
    To re-run:
    go test -run=FuzzReverse/af69258a12129d6cbba438df5d5f25ba0ec050461c116f777e77ea7c9a0d217a
FAIL
exit status 1
FAIL    example/fuzz  0.030s
```