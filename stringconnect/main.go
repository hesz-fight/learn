package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	// 加号拼接字符串
	a := "aaa"
	b := "bbb"
	c := a + b
	fmt.Println(c)
	// bytes 数组拼接字符串
	var chs bytes.Buffer
	fmt.Fprint(&chs, "aaa", "bbb")
	fmt.Println(chs.String())
	// StringBuilder 拼接字符串
	var str strings.Builder
	fmt.Fprint(&str, "aaa", "bbb")
	fmt.Println(str.String())
}
