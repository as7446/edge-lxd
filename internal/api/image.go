package api

import (
	"github.com/as7446/lxd/internal/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var imageService = service.NewImageService()

func ExportImage(c *gin.Context) {
	var requestData map[string]interface{}
	if err := c.ShouldBindJSON(&requestData); err != nil {
		log.Printf("ExportImage 参数解析失败: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求参数"})
		return
	}
	log.Println("API 请求导出镜像:", requestData)
	err := imageService.Export(requestData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "镜像导出失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "镜像导出成功"})
}

func DeleteImage(c *gin.Context) {
	imageID := c.Param("id")
	log.Println("API 请求删除镜像:", imageID)
	err := imageService.Delete(map[string]interface{}{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "镜像删除失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "镜像删除成功"})
}
