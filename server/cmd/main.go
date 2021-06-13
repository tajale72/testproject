package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// same as
	// config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	// router.Use(cors.New(config))
	router.Use(cors.Default())
	fmt.Println("New Project")
	router.POST("/main", Main)
	router.GET("/home", Home)
	// router.GET("/main", GetData)
	// router.GET("/main", InsertData)
	router.Run()
}

//Main displays hello world
func Main(c *gin.Context) {
	fmt.Println("heollo")
	c.JSON(http.StatusOK, gin.H{"data": "hello world"})
}

//Home displays hello world
func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "home world"})
}
