package utils

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
)

func ParseMessage[T any](msg amqp.Delivery) (*T, error) {
	var data T
	if err := json.Unmarshal(msg.Body, &data); err != nil {
		return nil, fmt.Errorf("解析消息体失败: %v", err)
	}
	return &data, nil
}
