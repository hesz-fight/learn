package main

import (
	"context"
	"log"
	"os"

	"github.com/olivere/elastic/v7"
)

type User struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Mrried bool   `json:"married"`
}

func main() {
	url := "http://127.0.0.1:9200"

	opts := []elastic.ClientOptionFunc{
		elastic.SetURL(url),
		elastic.SetErrorLog(log.New(os.Stderr, "", log.LstdFlags)),
	}
	cli, err := elastic.NewClient(opts...)
	if err != nil {
		panic(err)
	}
	ctx := context.Background()

	info, code, err := cli.Ping(url).Do(ctx)
	if err != nil {
		panic(err)
	}

	log.Printf("version:%v code:%v", info.Version.Number, code)

	user := User{
		Name:   "Xiaoming",
		Age:    10,
		Mrried: false,
	}
	// 7.0 建议移除 _type 统一为 _doc
	p1, err := cli.Index().Index("user").OpType("create").BodyJson(user).Do(context.Background())
	if err != nil {
		log.Printf("create err:%v\n", err)
	}

	log.Printf("_index:%v _type:%v _id:%v", p1.Index, p1.Type, p1.Id)
}
