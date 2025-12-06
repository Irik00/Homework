package api

import "github.com/gin-gonic/gin"

func InitRouter() {
	r := gin.Default()
	r.GET("/ping", Ping1)
	r.POST("/login", Login)
	r.POST("/register", Register)
	r.POST("/revision", ReviseKey)
	r.Run(":8080")
}
