package api

import (
	"github.com/as7446/lxd/internal/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var volumeService = service.NewVolumeService()

func CreateVolume(c *gin.Context) {
	var requestData map[string]interface{}
	if err := c.ShouldBindJSON(&requestData); err != nil {
		log.Printf("CreateVolume 参数解析失败: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求参数"})
		return
	}
	log.Println("API 请求创建存储卷:", requestData)
	err := volumeService.Create(requestData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "存储卷创建失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "存储卷创建成功"})
}

func DeleteVolume(c *gin.Context) {
	volumeID := c.Param("id")
	log.Println("API 请求删除存储卷:", volumeID)
	err := volumeService.Delete(map[string]interface{}{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "存储卷删除失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "存储卷删除成功"})
}
