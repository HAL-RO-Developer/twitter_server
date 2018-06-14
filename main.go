package main

import (
	"net/http"

	"github.com/HAL-RO-Developer/twitter_server/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok!")
	})
	r.POST("/api/tweet", controller.Tweet)
	r.GET("/api/tweets", controller.GetTweet)
}
