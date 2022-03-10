package router

import (
	"fmt"
	"go-dev/controller"
	"go-dev/model"
	"net/http"

	"io"

	"github.com/gin-gonic/gin"
)

func Adduser(c *gin.Context) {
	var success model.Success
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, "unable to read data")
	}
	err = controller.Adduser(body)
	if err != nil {
		c.JSONP(http.StatusBadRequest, "error adding userinfo")
	}
	success.Success = "successfully added userinfo"
	c.JSONP(http.StatusAccepted, success)

}

func Getuser(c *gin.Context) {
	name, _ := c.Params.Get("name")
	fmt.Println("name", name)
	userinfo, err := controller.Getuser(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, "error getting userinfo")
	}
	c.JSON(http.StatusAccepted, userinfo)

}

func Updateuser(c *gin.Context) {
	var success model.Success
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSONP(400, "unable to read data")
	}
	err = controller.Updateuser(body)
	if err != nil {
		c.JSONP(404, "error updating userinfo")
	}
	success.Success = "successfully added userinfo"
	c.JSONP(200, success)

}

func Deleteuser(c *gin.Context) {
	var success model.Success
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSONP(400, "unable to read data")
	}
	err = controller.Deleteuser(body)
	if err != nil {
		c.JSONP(404, "error adding userinfo")
	}
	success.Success = "successfully added userinfo"
	c.JSONP(200, success)

}
