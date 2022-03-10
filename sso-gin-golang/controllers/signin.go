package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var ssogolang *oauth2.Config
var RandomString = "random-string"

func init() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ssogolang = &oauth2.Config{
		RedirectURL:  os.Getenv("REDIRECT_URL"),
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

func Signin(c *gin.Context) {
	url := ssogolang.AuthCodeURL(RandomString)
	fmt.Println(url)
	c.Redirect(http.StatusTemporaryRedirect, url)
}
