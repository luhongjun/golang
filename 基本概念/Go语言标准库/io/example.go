package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 从标准输入读取
	dd, _ := ReadFrom(os.Stdin, 11)

	fmt.Println("the string is:" + string(dd))
}

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if err != nil {
		return nil, err
	}

	return p[:n], nil
}
