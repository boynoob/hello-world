package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/", hello) // what does this step do? It is not present in the documentation.
	router.Run()
}

func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello Vishal",
	})
}
