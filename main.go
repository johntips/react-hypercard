package main

import (
	"github.com/gin-gonic/gin"
    
)

func main() {
	// setup godo hotloader
	// main 
	r := gin.Default()
	r.GET("/astras", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/show", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "boom",
		})
	})

	r.Run()
}
