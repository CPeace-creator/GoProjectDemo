package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"goDemo/global"
	"goDemo/model"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func CreateExchangeRate(c *gin.Context) {
	var exchangeRate model.ExchangeRate
	if err := c.ShouldBindJSON(&exchangeRate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	exchangeRate.Date = time.Now()
	if err := global.DB.AutoMigrate(exchangeRate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := global.DB.Create(&exchangeRate).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, exchangeRate)
}

func GetExchangeRate(c *gin.Context) {
	var exchangeRates []model.ExchangeRate
	if err := global.DB.Find(&exchangeRates).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
	c.JSON(http.StatusOK, exchangeRates)
}
