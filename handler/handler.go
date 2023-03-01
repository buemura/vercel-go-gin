package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"message": "pong",
	})
}

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"message": "hello world",
	})
}

func Time(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"message": time.Now(),
	})
}

func ErrRouter(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status": "error",
		"message": "Route not found",
	})
}

func Cors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Next()
}