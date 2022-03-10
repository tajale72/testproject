package db

import (
	"context"
	"go-dev/model"
	"go-dev/mongo"

	"log"

	"github.com/globalsign/mgo/bson"
)

func Adduser(userinfo model.Signin) error {
	client, err := mongo.MongoConnection()
	if err != nil {
		log.Println("error : mongo cinet")
	}
	collection := client.Database("myapp").Collection("userinfo")
	_, err = collection.InsertOne(context.Background(), userinfo)
	if err != nil {
		log.Println("error Inserting client from mongodb", err)
		return err
	}
	return nil

}

func Getuser(name string) (*model.Signin, error) {
	var userinfo model.Signin
	client, err := mongo.MongoConnection()
	if err != nil {
		log.Println("error retriving client from mongodb", err)
	}
	collection := client.Database("myapp").Collection("userinfo")
	query := bson.M{"name": name}
	err = collection.FindOne(context.Background(), query).Decode(&userinfo)
	if err != nil {
		log.Println("error getting datafrom mongodb", err)
		return nil, err
	}
	return &userinfo, nil

}
