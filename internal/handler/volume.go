package handlers

import (
	"encoding/json"
	"github.com/as7446/lxd/internal/service"
	"log"

	"github.com/streadway/amqp"
)

var volumeService = service.NewVolumeService()

func HandleLocalVolumeCreate(msg amqp.Delivery) {
	var data map[string]interface{}
	json.Unmarshal(msg.Body, &data)
	log.Println("收到创建存储卷请求:", data)
	volumeService.Create(data)
}

func HandleLocalVolumeClone(msg amqp.Delivery) {
	var data map[string]interface{}
	json.Unmarshal(msg.Body, &data)
	log.Println("收到克隆存储卷请求:", data)
	volumeService.Clone(data)
}

func HandleLocalVolumeDelete(msg amqp.Delivery) {
	var data map[string]interface{}
	json.Unmarshal(msg.Body, &data)
	log.Println("收到删除存储卷请求:", data)
	volumeService.Delete(data)
}

func HandleLocalVolumeAttach(msg amqp.Delivery) {
	var data map[string]interface{}
	json.Unmarshal(msg.Body, &data)
	log.Println("收到绑定存储卷请求:", data)
	volumeService.Attach(data)
}

func HandleLocalVolumeDetach(msg amqp.Delivery) {
	var data map[string]interface{}
	json.Unmarshal(msg.Body, &data)
	log.Println("收到解绑存储卷请求:", data)
	volumeService.Detach(data)
}

func HandleLocalVolumeChangeSize(msg amqp.Delivery) {
	var data map[string]interface{}
	json.Unmarshal(msg.Body, &data)
	log.Println("收到调整存储卷大小请求:", data)
	volumeService.ChangeSize(data)
}
