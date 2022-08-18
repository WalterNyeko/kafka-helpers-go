package consumer

import (
	"context"
	. "github.com/WalterNyeko/kafka-helpers-go/models"
	. "github.com/WalterNyeko/kafka-helpers-go/utils"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	failedToConnect = "Failed to dial leader:"
	failedToCloseBatchReader = "Failed to close batch reader:"
	failedToCloseWriter = "Failed to close writer:"
)

func ConsumeMessage(topic string) ServiceMessage {

	writeTimeout, _ := os.LookupEnv("KAFKA_WRITE_TIMEOUT")
	address, _ := os.LookupEnv("KAFKA_ADDRESS")
	network, _ := os.LookupEnv("KAFKA_NETWORK")
	partition, _ := os.LookupEnv("KAFKA_PARTITION")

	partitionValue, _ := strconv.Atoi(partition)
	conn, err := kafka.DialLeader(context.Background(), network, address, topic, partitionValue)
	if err != nil {
		log.Println(failedToConnect, err)
		return ErrorMessage(err, failedToConnect)
	}

	writeTimeoutValue, _ := strconv.Atoi(writeTimeout)
	conn.SetReadDeadline(time.Now().Add(time.Duration(writeTimeoutValue)*time.Second))
	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	b := make([]byte, 10e3) // 10KB max per message
	for {
		n, err := batch.Read(b)
		if err != nil {
			break
		}
		log.Println(string(b[:n]))
		return ServiceMessage{
			ResponseData: string(b[:n]),
		}
	}

	if err := batch.Close(); err != nil {
		log.Println(failedToCloseBatchReader, err)
		return ErrorMessage(err, failedToCloseBatchReader)
	}

	if err := conn.Close(); err != nil {
		log.Println(failedToCloseWriter, err)
		return ErrorMessage(err, failedToCloseWriter)
	}
	return ServiceMessage{}
}