package main

import (
	"context"
	"fmt"
	"log"
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

	resp, err := cli.Grant(context.TODO(), 5)
	if err != nil {
		log.Fatal(err)
	}

	_, err = cli.Put(context.TODO(), "marwi", "dsb", clientv3.WithLease(resp.ID))
	if err != nil {
		log.Fatal(err)
	}
}
