package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/valianx/gin-template/pkg/domain/service"
	"net/http"
	"strconv"
)

func GetUserName(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	userName := service.GetUserName(id)

	fmt.Println(userName)
	if userName == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "User Not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"userName": userName})
}

func PostHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "new data"})
}
