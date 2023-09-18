package main

import (
	"fmt"
	"net/http"
	"web-service-gin/config"

	"github.com/gin-gonic/gin"
)

type success_result struct {
	Distance float64 `json:"distance"`
	Address  string  `json:"address"`
}

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		fmt.Printf("Error load config")
		return
	}
	router := gin.Default()
	router.GET("/distance", get_distance)
	address := cfg.Server.Host + ":" + cfg.Server.Port
	router.Run(address)
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
