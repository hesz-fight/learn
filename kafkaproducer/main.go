package main

import (
	"fmt"
	"log"

	"github.com/shopify/sarama"
)

type KafkaSyncProducer struct {
	Producer sarama.SyncProducer
	Topic    string
}

func NewSyncProducer(addrs []string, topic string, config *sarama.Config) (*KafkaSyncProducer, error) {
	if config == nil {
		config = sarama.NewConfig()
		config.Producer.RequiredAcks = sarama.WaitForAll
		config.Producer.Partitioner = sarama.NewRandomPartitioner
		config.Producer.Return.Successes = true
		config.Producer.Return.Errors = true
	}
	p, err := sarama.NewSyncProducer(addrs, config)
	if err != nil {
		return nil, err
	}

	return &KafkaSyncProducer{Producer: p, Topic: topic}, nil
}

func (p *KafkaSyncProducer) SendMessag(value string) (int32, int64, error) {
	msg := &sarama.ProducerMessage{
		Topic: p.Topic,
		Value: sarama.StringEncoder(value),
	}

	return p.Producer.SendMessage(msg)
}

func main() {
	addrs := []string{"127.0.0.1:9092"}
	p, err := NewSyncProducer(addrs, "kafka-test-topic", nil)
	if err != nil {
		panic(fmt.Sprintf("new kafka producer error:%v", err))
	}

	partion, offset, err := p.SendMessag("hello world")
	if err != nil {
		log.Printf("send message error:%v\n", err)
	}
	log.Printf("send message successful! partion:%v offser:%v\n", partion, offset)
}
