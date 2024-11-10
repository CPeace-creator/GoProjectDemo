package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"goDemo/global"
	"net/http"
)

// 实现点赞功能
func LikeArticle(c *gin.Context) {
	articleId := c.Param("id")
	likeKey := "article:" + articleId + ":likes"

	if err := global.RedisDB.Incr(likeKey).Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successful article"})
}

func GetArticleLikes(c *gin.Context) {
	articleId := c.Param("id")
	likeKey := "article:" + articleId + ":likes"
	likes, err := global.RedisDB.Get(likeKey).Result()
	if err == redis.Nil {
		likes = "0"
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"likes": likes})

}
