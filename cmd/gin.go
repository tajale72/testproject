package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/hello", Hello)
	r.POST("/hello", Hey)
	r.PATCH("/hello", Hi)
	r.Run()

}

//Hi is a function
func Hi(c *gin.Context) {
	c.JSONP(200, "hi")
}

//Hello is a function
func Hello(c *gin.Context) {

	var hi Hi
	hi.Name = []string{"Krischal", "romit"}
	hi.Adress = "413 village dr apt 120"
	hi.Number = 6127073040
	c.JSONP(200, hi)
}

type Hi struct {
	Name   []string
	Adress string
	Number int
}

func Hey(c *gin.Context) {
	var hi Hi
	hi.Name = []string{"Romit", "Krischal"}
	hi.Adress = "413 village dr apt 120"
	hi.Number = 6127073040
	c.JSON(200, hi)
}
