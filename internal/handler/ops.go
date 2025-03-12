package handlers

import (
	"encoding/json"
	"github.com/as7446/lxd/internal/service"
	"log"

	"github.com/streadway/amqp"
)

var opsService = service.NewOpsService()

func HandleOpsPingEdge(msg amqp.Delivery) {
	var data map[string]interface{}
	json.Unmarshal(msg.Body, &data)
	log.Println("收到 Ping Edge 请求:", data)
	opsService.PingEdge(data)
}

func HandleOpsPingEdgeInstance(msg amqp.Delivery) {
	var data map[string]interface{}
	json.Unmarshal(msg.Body, &data)
	log.Println("收到 Ping Edge Instance 请求:", data)
	opsService.PingEdgeInstance(data)
}

func HandleOpsGetEdgeInfo(msg amqp.Delivery) {
	var data map[string]interface{}
	json.Unmarshal(msg.Body, &data)
	log.Println("收到获取 Edge 信息请求:", data)
	opsService.GetEdgeInfo(data)
}

func HandleOpsRunEdgeInstanceCommand(msg amqp.Delivery) {
	var data map[string]interface{}
	json.Unmarshal(msg.Body, &data)
	log.Println("收到运行 Edge 命令请求:", data)
	opsService.RunEdgeInstanceCommand(data)
}
