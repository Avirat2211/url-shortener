package handler

import (
	"net/http"

	"github.com/Avirat2211/url-shortener/shortener"
	"github.com/Avirat2211/url-shortener/store"
	"github.com/gin-gonic/gin"
)

type UrlReq struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	var req UrlReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	short := shortener.GenerateShortenUrl(req.LongUrl, req.UserId)
	store.SaveUrlMapping(short, req.LongUrl, req.UserId)
	host := "http://localhost:9808/"
	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host + short,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	short := c.Param("shortUrl")
	original := store.RetriveInitialUrl(short)
	c.Redirect(302, original)
}
