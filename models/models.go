package models

import "reflect"

type ServiceMessage struct {
	StatusMessage  string `json:"statusMessage,omitempty"`
	SupportMessage string `json:"supportMessage,omitempty"`
	StatusCode     int    `json:"statusCode,omitempty"`
	ResponseCode   string `json:"responseCode,omitempty"`
	ResponseData   string `json:"responseData"`
}

func (s ServiceMessage) IsEmpty() bool {
	return reflect.DeepEqual(s, ServiceMessage{})
}

func (s ServiceMessage) IsMissingResponseData() bool {
	return s.ResponseData == ""
}