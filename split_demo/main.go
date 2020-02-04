package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	str := "hello"
	index := 2
	for index >= 0 {
		fmt.Println("before:", str)
		str := str[index+1:]
		fmt.Println("after", str)
		index = strings.Index(str, "l")
		time.Sleep(time.Second)
	}
}
