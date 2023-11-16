package arrays

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BinarySearch struct {
	Array  []int `json:"array"`
	Target int   `json:"target"`
}

func HandleBinarySearch(c *gin.Context) {
	var bs BinarySearch
	if err := c.BindJSON(&bs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	result := binarySearch(bs)
	c.JSON(http.StatusOK, gin.H{"index": result})
}

type ProductExceptSelf struct {
	Array []int `json:"nums"`
}

func HandleProductExceptSelf(c *gin.Context) {
	var nums ProductExceptSelf

	if err := c.BindJSON(&nums); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	result := productExceptSelf(nums)
	c.JSON(http.StatusOK, gin.H{"result": result})
}
