package handlers

import (
	"encoding/json"
	"github.com/as7446/lxd/internal/service"
	"github.com/as7446/lxd/internal/utils"
	"log"

	"github.com/streadway/amqp"
)

var instanceService = service.NewInstanceService()

type data map[string]interface{}

func HandleInstanceCreate(msg amqp.Delivery) {
	var data map[string]interface{}
	message, _ := utils.ParseMessage[data](msg)
	json.Unmarshal(msg.Body, &data)
	log.Println("收到创建实例请求:", message)
	instanceService.Create(data)
}

func HandleInstanceDelete(msg amqp.Delivery) {
	var data map[string]interface{}
	json.Unmarshal(msg.Body, &data)
	instanceID := data["id"].(string)
	log.Println("收到删除实例请求:", instanceID)
	instanceService.Delete(instanceID)
}

func HandleInstanceStart(msg amqp.Delivery) {
	var data map[string]interface{}
	json.Unmarshal(msg.Body, &data)
	instanceID := data["id"].(string)
	log.Println("收到启动实例请求:", instanceID)
	instanceService.Start(instanceID)
}

func HandleInstanceShutdown(msg amqp.Delivery) {
	var data map[string]interface{}
	json.Unmarshal(msg.Body, &data)
	instanceID := data["id"].(string)
	log.Println("收到关闭实例请求:", instanceID)
	instanceService.Shutdown(instanceID)
}

func HandleInstanceAttachGPU(msg amqp.Delivery) {
	var data map[string]interface{}
	json.Unmarshal(msg.Body, &data)
	instanceID := data["id"].(string)
	log.Println("收到绑定 GPU 请求:", instanceID)
	instanceService.AttachGPU(instanceID)
}

func HandleInstanceDetachGPU(msg amqp.Delivery) {
	var data map[string]interface{}
	json.Unmarshal(msg.Body, &data)
	instanceID := data["id"].(string)
	log.Println("收到解绑 GPU 请求:", instanceID)
	instanceService.DetachGPU(instanceID)
}

func HandleInstanceModifyConfig(msg amqp.Delivery) {
	var data map[string]interface{}
	json.Unmarshal(msg.Body, &data)
	instanceID := data["id"].(string)
	log.Println("收到修改配置请求:", instanceID)
	instanceService.ModifyConfig(instanceID, data)
}

func HandleInstanceEnableNetworkProxy(msg amqp.Delivery) {
	var data map[string]interface{}
	json.Unmarshal(msg.Body, &data)
	instanceID := data["id"].(string)
	log.Println("收到启用网络代理请求:", instanceID)
	instanceService.EnableNetworkProxy(instanceID)
}

func HandleInstanceDisableNetworkProxy(msg amqp.Delivery) {
	var data map[string]interface{}
	json.Unmarshal(msg.Body, &data)
	instanceID := data["id"].(string)
	log.Println("收到禁用网络代理请求:", instanceID)
	instanceService.DisableNetworkProxy(instanceID)
}

func HandleInstanceSetupApplications(msg amqp.Delivery) {
	var data map[string]interface{}
	json.Unmarshal(msg.Body, &data)
	instanceID := data["id"].(string)
	apps := data["apps"].([]string)
	log.Println("收到安装应用请求:", instanceID, apps)
	instanceService.SetupApplications(instanceID, apps)
}

func HandleInstanceSetupApplicationsV2(msg amqp.Delivery) {
	var data map[string]interface{}
	json.Unmarshal(msg.Body, &data)
	instanceID := data["id"].(string)
	apps := data["apps"].([]string)
	log.Println("收到安装应用 V2 请求:", instanceID, apps)
	instanceService.SetupApplicationsV2(instanceID, apps)
}

func HandleInstanceModifyPassword(msg amqp.Delivery) {
	var data map[string]interface{}
	json.Unmarshal(msg.Body, &data)
	instanceID := data["id"].(string)
	password := data["password"].(string)
	log.Println("收到修改密码请求:", instanceID)
	instanceService.ModifyPassword(instanceID, password)
}
