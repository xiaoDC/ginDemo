package main

import "github.com/gin-gonic/gin"

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", startPage)
	router.GET("/:user", userPage)

	return router
}

func startPage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "gong",
	})
}

func userPage(c *gin.Context) {
	name := c.Param("user")

	c.JSON(200, gin.H{
		"message": name,
		"name":    "hahah",
	})
}
