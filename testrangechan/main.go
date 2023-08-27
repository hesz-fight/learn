package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 10)
	go func() {
		// range 遍历 channel
		// channel 被关闭后，for 结束，最终不会再进入循环体
		for c := range ch {
			fmt.Println("read:", c)
		}
		// 关闭的通道还可以继续读，读到的是零值
		fmt.Printf("read afrer for-range: %v\n", <-ch)

		fmt.Println("exit")
	}()

	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	close(ch)
	time.Sleep(5 * time.Second)
}
