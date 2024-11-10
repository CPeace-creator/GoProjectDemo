package controller

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"goDemo/global"
	"goDemo/model"
	"gorm.io/gorm"
	"net/http"
	"time"
)

var cacheKey = "articles"

func CreateArticle(c *gin.Context) {
	var article model.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := global.DB.AutoMigrate(&article); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := global.DB.Create(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := global.RedisDB.Del(cacheKey); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Err().Error()})
		return
	}
	c.JSON(http.StatusOK, article)
}

func GetArticles(c *gin.Context) {
	cachedData, err := global.RedisDB.Get(cacheKey).Result()
	if err == redis.Nil {
		var articles []model.Article
		if err := global.DB.Find(&articles).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			articleJson, err := json.Marshal(articles)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			if err := global.RedisDB.Set(cacheKey, articleJson, 10*time.Second); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Err().Error()})
			}
		}
		c.JSON(http.StatusOK, articles)
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	} else {
		var articles []model.Article
		if err := json.Unmarshal([]byte(cachedData), &articles); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, articles)
	}
}

func GetArticleById(c *gin.Context) {
	id := c.Param("id")
	var article model.Article
	if err := global.DB.Where("id = ?", id).First(&article).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, article)
}
