package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func f(ctx context.Context) {
	wg.Done()
LOOP:
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("f")
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
}
func main() {
	ctx, cancle := context.WithCancel(context.Background())
	wg.Add(1)
	go f(ctx)
	time.Sleep(time.Second * 5)
	cancle()
	wg.Wait()
}
