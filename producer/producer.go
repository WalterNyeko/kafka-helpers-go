package producer

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
	failedToWrite = "Failed to write message:"
	failedToCloseWriter = "Failed to close writer:"
)

func PublishToTopic(message string, topic string) ServiceMessage {

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
	conn.SetWriteDeadline(time.Now().Add(time.Duration(writeTimeoutValue)*time.Second))

	_, err = conn.WriteMessages(kafka.Message{Value: []byte(message)})
	if err != nil {
		log.Println(failedToWrite, err)
		return ErrorMessage(err, failedToWrite)
	}

	if err := conn.Close(); err != nil {
		log.Println(failedToCloseWriter, err)
		return ErrorMessage(err, failedToCloseWriter)
	}
	return ServiceMessage{}
}


