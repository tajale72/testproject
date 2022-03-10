package controller

import (
	"encoding/json"
	"fmt"
	"go-dev/db"
	"go-dev/model"
	"log"
)

func Adduser(body []byte) error {
	var userinfo model.Signin
	json.Unmarshal(body, &userinfo)
	err := db.Adduser(userinfo)
	if err != nil {
		log.Println("error from Adduser", err)
		return err
	}
	return nil

}

func Getuser(name string) (*model.Signin, error) {
	userinfo, err := db.Getuser(name)
	if err != nil {
		log.Println("error from Get user")
		return nil, err
	}
	return userinfo, nil

}

func Updateuser(body []byte) error {
	var userinfo model.Signin
	json.Unmarshal(body, &userinfo)
	err := db.Adduser(userinfo)
	if err != nil {
		fmt.Println("error from PostSignin", err)
		return err
	}
	return nil

}

func Deleteuser(body []byte) error {
	var userinfo model.Signin
	json.Unmarshal(body, &userinfo)
	err := db.Adduser(userinfo)
	if err != nil {
		fmt.Println("error from PostSignin", err)
		return err
	}
	return nil

}
