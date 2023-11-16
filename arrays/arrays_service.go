package arrays

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleBinarySearch(c *gin.Context) {
	var bs BinarySearch
	if err := c.BindJSON(&bs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	result := binarySearch(bs)
	c.JSON(http.StatusOK, gin.H{"index": result})
}
