package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	api := router.Group("/api")
	{
		api.Get("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}
}
