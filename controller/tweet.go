package controller

import (
	"net/http"

	"github.com/HAL-RO-Developer/twitter_server/model"
	"github.com/gin-gonic/gin"
)

func Tweet(c *gin.Context) {
	tweet := model.Tweet{}
	err := c.BindJSON(&tweet)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	if tweet.Image == "" {
		tweet.Image = "https://pbs.twimg.com/profile_images/778387920440823808/kw35Gu4T_400x400.jpg"
	}
	tweet.Save()
	c.JSON(http.StatusOK, "ok")
}

func GetTweet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"tweets": model.GetTweets(),
	})
}
