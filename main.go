package main

import (
	"andrew-ellis-engineering/golang-problem-solver/arrays"
	"andrew-ellis-engineering/golang-problem-solver/graphs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/ping", ping)

	router.POST("/arrays/binarySearch", arrays.HandleBinarySearch)
	router.POST("/arrays/productExceptSelf", arrays.HandleProductExceptSelf)

	router.POST("/graphs/countIslands", graphs.HandleIslandCount)

	router.Run("localhost:8080")
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
