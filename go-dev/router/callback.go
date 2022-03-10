package router

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func Callback(c *gin.Context) {
	state := c.Request.FormValue("state")
	code := c.Request.FormValue("code")
	data, err := getUserData(state, code)
	if err != nil {
		log.Println("error getting user data")
	}
	fmt.Printf("Data: %s", data)

}

func getUserData(state, code string) ([]byte, error) {
	var ssogolang *oauth2.Config
	var RandomString = "random-string"

	if state != RandomString {
		return nil, errors.New("Invalid user state")
	}
	token, err := ssogolang.Exchange(context.Background(), code)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return data, nil

}
