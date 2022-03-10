package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Callback(c *gin.Context) {
	var res interface{}
	state := c.Request.FormValue("state")
	code := c.Request.FormValue("code")
	data, err := getUserData(state, code)
	json.Unmarshal(data, &res)
	if err != nil {
		log.Fatal("error getting user data")
	}
	c.JSON(http.StatusAccepted, res)
}
func getUserData(state, code string) ([]byte, error) {
	if state != RandomString {
		return nil, errors.New("invalid user state")
	}
	token, err := ssogolang.Exchange(context.Background(), code)
	if err != nil {
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
