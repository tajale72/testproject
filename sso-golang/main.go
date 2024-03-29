package main

import (
	"net/http"
	"sso-golang/controllers"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	http.HandleFunc("/signin", controllers.Signin)
	http.HandleFunc("/callback", controllers.Callback)
	http.ListenAndServe(":3000", nil)
}
