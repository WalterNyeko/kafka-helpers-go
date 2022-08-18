package consumer

import (
	"context"
	. "github.com/WalterNyeko/kafka-helpers-go/models"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
	"strconv"
)

const (
	groupId = "G1"
)

func ConsumeMessage(topic string) ServiceMessage {

	address, _ := os.LookupEnv("KAFKA_ADDRESS")
	partition, _ := os.LookupEnv("KAFKA_PARTITION")
	partitionValue, _ := strconv.Atoi(partition)


	config := kafka.ReaderConfig{
		Brokers: []string{address},
		Topic: topic,
		GroupID: groupId,
		MinBytes: 10e3,
		MaxBytes: 1e6,
		Partition: partitionValue,
	}

	reader := kafka.NewReader(config)

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			continue
			return ServiceMessage{}
		}
		log.Println(string(m.Value))
		return ServiceMessage{
			ResponseData: string(m.Value),
		}
	}
}