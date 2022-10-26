package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
)

func NewSingleConsumer() {

}

func NewConsumerGroup() {
	cfg := kafka.ConsumerGroupConfig{
		ID:                     "test-1",
		Brokers:                []string{"127.0.0.1:9092"},
		Dialer:                 nil,
		Topics:                 []string{"test"},
		GroupBalancers:         nil,   // 组平衡策略
		HeartbeatInterval:      0,     // 心跳包间隔
		PartitionWatchInterval: 0,     // 监听分区变化
		WatchPartitionChanges:  false, // 是否监听分区变化
		SessionTimeout:         0,     // 心跳包超时时间
		RebalanceTimeout:       0,     // 等待新成员加入组的超时时间
		JoinGroupBackoff:       0,     // 加入失败后，等待多久后重新加入
		RetentionTime:          0,     //
		StartOffset:            0,     // 开始消费的位移
		Logger:                 nil,   // 自定义日志
		ErrorLogger:            nil,   // 自定义错误日志
		Timeout:                0,     // 与消费者组协调器的超时时间
	}

	group, err := kafka.NewConsumerGroup(cfg)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	next, err := group.Next(context.Background())
	next.Start()
	fmt.Println(next)
}

func Consumer() {
	// make a new reader that consumes from topic-A
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"localhost:9092"},
		GroupID:     "consumer-group-1",
		Topic:       "test",
		MinBytes:    10e3, // 10KB
		MaxBytes:    10e6, // 10MB
		StartOffset: kafka.FirstOffset,
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}

func main() {
	Consumer()
}
