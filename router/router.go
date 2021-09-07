package router

import (
	"gin-temp/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)
func LoadRouter(g *gin.Engine,mw ...gin.HandlerFunc) *gin.Engine {
	// 加载中间件
	g.Use(mw...)

	// 404
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	test:=g.Group("/test")
	{
		test.GET("/v1",controllers.TestV1)
	}
	return g
}
