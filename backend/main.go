package main

import (
	"net/http"
	"web-service-gin/config"

	"github.com/gin-gonic/gin"
)

type success_result struct {
	Distance float64 `json:"distance"`
	Address  string  `json:"address"`
}

func main() {
	cfg, err := config.GetConfig("config.yaml")
	if err != nil {
		panic(err)
	}
	router := setupRouter()
	address := cfg.Server.Host + ":" + cfg.Server.Port
	router.Run(address)
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/distance", get_distance)
	return router
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
