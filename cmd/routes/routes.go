package routes

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/valianx/gin-template/pkg/api/handler"
	"log"
	"net/http"
	"strconv"
)

func SetupRoutes(port int) *gin.Engine {
	portStr := strconv.Itoa(port)

	fmt.Printf("connect to port %s\n", portStr)

	r := gin.New()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	api := r.Group("/user")
	{
		api.GET("/:id", handler.GetUserName)
		api.POST("/", handler.PostHandler)
	}

	if err := http.ListenAndServe(":"+portStr, r); err != nil {
		log.Fatal(err)
	}
	return r
}
