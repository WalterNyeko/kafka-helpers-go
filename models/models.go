package models

import "reflect"

type ServiceMessage struct {
	StatusMessage  string `json:"statusMessage,omitempty"`
	SupportMessage string `json:"supportMessage,omitempty"`
	StatusCode     int    `json:"statusCode,omitempty"`
	ResponseCode   string `json:"responseCode,omitempty"`
}

func (s ServiceMessage) IsEmpty() bool {
	return reflect.DeepEqual(s, ServiceMessage{})
}
