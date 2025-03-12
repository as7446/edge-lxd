package api

import (
	"github.com/as7446/lxd/internal/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var opsService = service.NewOpsService()

func PingEdge(c *gin.Context) {
	log.Println("API 请求 Ping Edge")
	err := opsService.PingEdge(nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ping Edge 失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Ping Edge 成功"})
}

func GetEdgeInfo(c *gin.Context) {
	log.Println("API 请求获取 Edge 信息")
	err := opsService.GetEdgeInfo(map[string]interface{}{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取 Edge 信息失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "data"})
}
