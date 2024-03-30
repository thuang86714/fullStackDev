package main

import (
	"github.com/thuang86714/fullStackDev/initializers"
	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVariables()
}

func main() { 
	r := gin.Default()//default router
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message" : "pong",
		})
	})

	r.Run()
}