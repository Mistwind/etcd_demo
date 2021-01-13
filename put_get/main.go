package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("Connect to etcd failed, err: %v\n", err)
		return
	}
	fmt.Printf("Connect to etcd success\n")
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "marwi", "huang")
	cancel()
	if err != nil {
		fmt.Printf("Put to etcd failed, err: %v\n", err)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "marwi")
	cancel()
	if err != nil {
		fmt.Printf("Get from etcd failed, err: %v\n", err)
		return
	}

	for _, ev := range resp.Kvs {
		fmt.Printf("%s: %s\n", ev.Key, ev.Value)
	}
}
