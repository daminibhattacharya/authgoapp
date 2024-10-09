package router

import (
	"auth-go-app/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.File("./index.html")
	})
	r.POST("/register", controller.RegisterUser)
	r.GET("/health", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"data": "API is healthy"}) })
	return r
}
