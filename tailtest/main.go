package main

import (
	"fmt"
	"os"
	"time"

	"github.com/hpcloud/tail"
)

// 从文件末尾开始读取
// 分隔符是换行符
func main() {
	filename := "./tail.log"
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: os.SEEK_END},
		MustExist: false,
		Poll:      true,
	}

	tail, err := tail.TailFile(filename, config)
	if err != nil {
		panic(err.Error())
	}

	for {
		msg, ok := <-tail.Lines
		if !ok {
			time.Sleep(time.Second)
			continue
		}
		fmt.Println(msg.Text)
		fmt.Println(msg.Time)
	}
}
