package handlers

import (
	"encoding/json"
	"github.com/as7446/lxd/internal/service"
	"log"

	"github.com/streadway/amqp"
)

var imageService = service.NewImageService()

func HandleImageExport(msg amqp.Delivery) {
	var data map[string]interface{}
	json.Unmarshal(msg.Body, &data)
	log.Println("收到导出镜像请求:", data)
	imageService.Export(data)
}

func HandleImageDelete(msg amqp.Delivery) {
	var data map[string]interface{}
	json.Unmarshal(msg.Body, &data)
	log.Println("收到删除镜像请求:", data)
	imageService.Delete(data)
}

func HandleImageCopy(msg amqp.Delivery) {
	var data map[string]interface{}
	json.Unmarshal(msg.Body, &data)
	log.Println("收到复制镜像请求:", data)
	imageService.Copy(data)
}
