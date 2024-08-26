package routers

import (
	"mediahub/controller"

	"github.com/gin-gonic/gin"
)

func InitRouters(api *gin.RouterGroup, c *controller.Controller) {
	v1 := api.Group("/v1")
	fileGroup := v1.Group("/file")
	fileGroup.POST("/upload", c.Upload)
}
