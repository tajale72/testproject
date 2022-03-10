package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Callback(c *gin.Context) {
	fmt.Println(200, "callback")
}
