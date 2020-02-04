package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("err1:", err)
		return
	}
	fmt.Println("etcd connect success")
	defer cli.Close()
	ch := cli.Watch(context.Background(), "dave")
	for wresp := range ch {
		for _, evt := range wresp.Events {
			fmt.Printf("type:%v,key:%v,value:%v \n", evt.Type, string(evt.Kv.Key), string(evt.Kv.Value))
		}
	}
}
