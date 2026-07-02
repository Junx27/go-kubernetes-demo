package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/Junx27/go-kubernetes-demo/internal/metrics"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {

		metrics.HttpRequests.Inc()

		c.JSON(http.StatusOK, gin.H{
			"message": "Hello Kubernetes V2",
		})
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "UP",
		})
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "UP",
		})
	})

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.Run(":8080")
	metrics.Register()
}
