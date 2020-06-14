package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type Joke struct {
	ID    int    `json:"id" binding:"required"`
	Likes int    `json:"likes"`
	Joke  string `json:"joke" binding:"required"`
}

func main() {
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	api.GET("/jokes", JokeHandler)
	api.POST("/jokes/like/:jokeID", LikeJoke)

	router.Run(":3000")
}

func JokeHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": "not implemented",
	})
}

func LikeJoke(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": "not implemented",
	})
}
