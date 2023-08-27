package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

var etcdCli *clientv3.Client

func initEtcd() {
	endpointds := []string{"127.0.0.1:2379"}
	dialTimeout := 3 * time.Second
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpointds,
		DialTimeout: dialTimeout,
	})
	if err != nil {
		panic(fmt.Sprintf("init etcd error:%v", err))
	}
	etcdCli = cli
}

func putOp(k string, v string) {
	_, err := etcdCli.Put(context.Background(), k, v)
	if err != nil {
		fmt.Printf("get resp for form ctcd error:%v", err)
		return
	}
}

func getOp(k string) {
	resp, err := etcdCli.Get(context.Background(), k)
	if err != nil {
		fmt.Printf("get resp for form ctcd error:%v", err)
		return
	}
	for _, kv := range resp.Kvs {
		fmt.Printf("getOp key:%v, value:%v\n", string(kv.Key), string(kv.Value))
	}
}

func watchOp() {
	// 监视键的变化
	respChan := etcdCli.Watch(context.Background(), "test")
	for resp := range respChan {
		for _, event := range resp.Events {
			fmt.Println("watchOp type:", event.Type, "key:", string(event.Kv.Key), "val:", string(event.Kv.Value))
		}
	}
}

func main() {
	initEtcd()
	go watchOp()
	time.Sleep(1 * time.Second)
	k := "test"
	getOp(k)
	putOp(k, "world")
	getOp(k)
	putOp(k, "hello")
	time.Sleep(3 * time.Second)
}
