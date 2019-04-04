package main

import (
	"github.com/gin-gonic/gin"
	"github.com/poom3d/go-boilerplate/audiences"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/health", healthCheck)

	provider := &audiences.SampleProvider{}
	audiences.Initialize(provider, r)

	r.Run(":8080")
}

func healthCheck(c *gin.Context) {
	c.Status(http.StatusOK)
}
