package api

import "github.com/gin-gonic/gin"

func setupRouter(router *gin.Engine) {
	ws := router.Group("/api/v1/ws/")
	{
		ws.GET("/", func(c *gin.Context) {
			wsHandler(c, c.Writer, c.Request)
		})
	}
}
