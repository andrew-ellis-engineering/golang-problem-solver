package main

import (
	// "net/http"

	"andrew-ellis-engineering/golang-problem-solver/arrays"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", ping)
	router.POST("/binarySearch", arrays.HandleBinarySearch)
	router.POST("/productExceptSelf", arrays.HandleProductExceptSelf)

	router.Run("localhost:8080")
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
