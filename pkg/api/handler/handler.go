package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "test"})
}

func PostHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "new data"})
}
