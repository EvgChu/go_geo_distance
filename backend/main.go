package main

import (
	"net/http"
	"web-service-gin/config"
	"web-service-gin/yageoservice"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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
	initLoger(cfg.Server.LogLevel)
	router := setupRouter()
	log.WithFields(log.Fields{
		"Server":       cfg.Server.LogLevel,
		"Host":         cfg.Server.Host,
		"Port":         cfg.Server.Port,
		"YandexApiKey": cfg.Server.YaApiKey,
	}).Info("Start server")
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
	ygs, err := yageoservice.Coordinates(address)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad address"})
		return
	}
	log.Info("Received request:", ygs.Addr)
	var distance float64 = 64 //todo add calculate
	result := success_result{
		Distance: distance,
		Address:  address,
	}

	c.IndentedJSON(http.StatusOK, result)
}

func initLoger(log_level string) {

	// setup logrus
	logLevel, err := log.ParseLevel(log_level)
	if err != nil {
		logLevel = log.InfoLevel
	}

	log.SetLevel(logLevel)
	log.SetFormatter(&log.JSONFormatter{})
}
