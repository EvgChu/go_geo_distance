package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type success_result struct {
	Distance float64 `json:"distance"`
	Address  string  `json:"address"`
}

func main() {
	router := gin.Default()
	router.GET("/distance", get_distance)
	router.Run("127.0.0.1:8080")
}

func get_distance(c *gin.Context) {
	address := c.Query("address")
	if address == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Empty address"})
		return

	}
	var distance float64 = 64 //todo add calculate
	result := success_result{
		Distance: distance,
		Address:  address,
	}

	c.IndentedJSON(http.StatusOK, result)
}
