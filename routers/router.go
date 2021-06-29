package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mzit/order_api/pkg/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.Runmode)
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test231231",
		})
	})

	return r
}
