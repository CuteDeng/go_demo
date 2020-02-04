package main

import (
	"fmt"

	"github.com/hpcloud/tail"
)

func main() {
	t, err := tail.TailFile("./my.log", tail.Config{
		Follow:    true,
		ReOpen:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	})
	if err != nil {
		fmt.Println("tail file err:", err)
		return
	}
	for line := range t.Lines {
		fmt.Println(line.Text)
	}
}
