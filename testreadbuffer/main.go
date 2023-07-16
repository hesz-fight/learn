package main

import (
	"bufio"
	"fmt"
	"os"
)

func testScanln() {
	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
}

func testBuffer() {
	input := bufio.NewReader(os.Stdin)
	str, _ := input.ReadString('\n')
	fmt.Println(str)

}

func main() {
	testBuffer()
}
