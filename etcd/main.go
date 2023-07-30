package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {
	endpointds := []string{"127.0.0.1:2379"}
	dialTimeout := 3 * time.Second
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpointds,
		DialTimeout: dialTimeout,
	})
	if err != nil {
		panic(fmt.Sprintf("init etcd error:%v", err))
	}
	defer cli.Close()
	ctx := context.Background()
	resp, err := cli.Get(ctx, "test")
	if err != nil {
		fmt.Printf("get resp for form ctcd error:%v", err)
		return
	}
	for _, kv := range resp.Kvs {
		fmt.Printf("key:%v, value:%v", string(kv.Key), string(kv.Value))
	}
}
