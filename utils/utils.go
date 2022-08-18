package utils

import (
	"fmt"
	. "github.com/WalterNyeko/kafka-helpers-go/models"
)

const (
	internalServerError = 500
	responseCode = "5000"
)
func ErrorMessage(err error, message string) ServiceMessage {
	return ServiceMessage{
		StatusMessage:  fmt.Sprintf(message + " %s", err.Error()),
		StatusCode:     internalServerError,
		ResponseCode:   responseCode,
		SupportMessage: err.Error(),
	}
}
