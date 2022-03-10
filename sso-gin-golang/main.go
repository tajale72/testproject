package main

import (
	"net/http"
	"sso-gin-golang/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	router.GET("/signin", controllers.Signin)
	router.GET("/callback", controllers.Callback)
	router.Run(":3000")
}
