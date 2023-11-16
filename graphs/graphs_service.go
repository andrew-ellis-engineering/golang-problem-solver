package graphs

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IslandCount struct {
	Grid [][]string `json:"grid"`
}

func HandleIslandCount(c *gin.Context) {
	var ic IslandCount
	if err := c.BindJSON(&ic); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	result := numIslands(ic)
	c.JSON(http.StatusOK, gin.H{"count": result})
}
