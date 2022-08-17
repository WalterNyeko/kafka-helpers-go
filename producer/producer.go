package producer

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	. "kafka-helpers-go/models"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	internalServerError = 500
	responseCode = "5000"
	failedToConnect = "Failed to dial leader:"
	failedToWrite = "Failed to write message:"
	failedToCloseWriter = "Failed to close writer:"
)

func PublishToTopic(message interface{}, topic string) ServiceMessage {

	writeTimeout, _ := os.LookupEnv("KAFKA_WRITE_TIMEOUT")
	address, _ := os.LookupEnv("KAFKA_ADDRESS")
	network, _ := os.LookupEnv("KAFKA_NETWORK")
	partition, _ := os.LookupEnv("KAFKA_PARTITION")

	partitionValue, _ := strconv.Atoi(partition)
	conn, err := kafka.DialLeader(context.Background(), network, address, topic, partitionValue)
	if err != nil {
		log.Println(failedToConnect, err)
		return errorMessage(err, failedToConnect)
	}

	writeTimeoutValue, _ := strconv.Atoi(writeTimeout)
	conn.SetWriteDeadline(time.Now().Add(time.Duration(writeTimeoutValue)*time.Second))

	_, err = conn.WriteMessages(kafka.Message{Value: message.([]byte)})
	if err != nil {
		log.Println(failedToWrite, err)
		return errorMessage(err, failedToWrite)
	}

	if err := conn.Close(); err != nil {
		log.Println(failedToCloseWriter, err)
		return errorMessage(err, failedToCloseWriter)
	}
	return ServiceMessage{}
}

func errorMessage(err error, message string) ServiceMessage {
	return ServiceMessage{
		StatusMessage:  fmt.Sprintf(message + " %s", err.Error()),
		StatusCode:     internalServerError,
		ResponseCode:   responseCode,
		SupportMessage: err.Error(),
	}
}
