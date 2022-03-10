package main

import (
	"go-dev/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	c := cors.New(config)
	r.Use(c)
	//creating end point
	r.POST("/adduser", router.Adduser)
	r.GET("/adduser/:name", router.Getuser)
	r.PUT("/adduser/:id", router.Updateuser)
	r.DELETE("/adduser", router.Deleteuser)

	//sso handlres
	//r.StaticFS("/public/index.html", http.Dir("public"))
	r.GET("/signin", router.Signin)
	r.GET("/callback", router.Callback)

	r.Run(":3000")

}
