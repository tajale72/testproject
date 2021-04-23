package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//SwaggerInput is a struct or
type SwaggerInput struct {
	FName string
	LName string
}

//swaggerAPIis a function.
func swaggerAPI(c *gin.Context) {
	//Creating an object for SwaggerInput struct
	x := SwaggerInput{"romit", "tajale"}

	c.JSON(http.StatusOK, gin.H{"data": x})
}

func main() {
	fmt.println("hello")
	//Creating a gin router
	r := gin.Default()

	//GET method for swagger endpoint.

	r.GET("/swagger", swaggerAPI)

	//This will run in port : 8080
	r.Run()
}
