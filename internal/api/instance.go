package api

import (
	"github.com/as7446/lxd/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var instanceService = service.NewInstanceService()

func CreateInstance(c *gin.Context) {
	var requestData map[string]interface{}
	if err := c.ShouldBindJSON(&requestData); err != nil {
		log.Printf("CreateInstance 参数解析失败: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求参数"})
		return
	}
	log.Println("API 请求创建实例:", requestData)
	err := instanceService.Create(requestData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "实例创建失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "实例创建成功"})
}

func DeleteInstance(c *gin.Context) {
	instanceID := c.Param("id")
	log.Println("API 请求删除实例:", instanceID)
	err := instanceService.Delete(instanceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "实例删除失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "实例删除成功"})
}
