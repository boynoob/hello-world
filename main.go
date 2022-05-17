package main

import "github.com/gin-gonic/gin"

func main() {
	// declaring the router variable
	router := gin.Default()

	/*
	 * HTTP GET request on the root(/) resource and define a function to be called when HTTP request
	 * hits the root end point. The function sends a JSON response with status code of 200(OK) with
	 * a body of "message" : "Hello Vishal"
	 */
	router.GET("/", hello)

	/*
	 * Deploy the router on port 8080 using router.Run() method.
	 * By default HTTP server is listening on port 8080. You can define a different port by adding
	 * an argument to the Run method. router.Run(":5000"). Note that the port parameter should
	 * be passed as a string, prepended by a colon punctuation.
	 */
	router.Run()
}

func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello Vishal",
	})
}
