package controllers

import (
	"context"
	"log"
	"time"

	"github.com/denzelIgnatius/auth_server_with_go/database"
	"github.com/denzelIgnatius/auth_server_with_go/encryption"
	"github.com/denzelIgnatius/auth_server_with_go/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddUsers(c *gin.Context) {

	var request struct {
		Username string
		Password string
	}

	c.Bind(&request)

	if len(request.Username) == 0 || len(request.Password) == 0 {
		const errorMsg string = "Error: Missing Input"
		log.Println(errorMsg)
		c.JSON(400, errorMsg)
		return
	}

	filter := bson.D{{"username", request.Username}}
	var result models.User
	err := database.Collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil || len(result.Username) != 0 {
		const errorMsg string = "Error: User already exists"
		log.Println(errorMsg)
		c.JSON(400, errorMsg)
		return
	}

	user := models.User{
		ID:        primitive.NewObjectID(),
		Username:  request.Username,
		Password:  encryption.SHA256Hashing(request.Password),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	res, err := database.Collection.InsertOne(context.TODO(), user)
	if err != nil || res != nil {
		var errorMsg string = "Error: " + err.Error()
		log.Println(errorMsg)
		c.JSON(500, errorMsg)
		return
	}
	c.JSON(200, gin.H{})
}

func DeleteUsers(c *gin.Context) {

	var request struct {
		Username string
	}

	c.Bind(&request)

	if len(request.Username) == 0 {
		const errorMsg string = "Error: Missing Input"
		log.Println(errorMsg)
		c.JSON(400, errorMsg)
		return
	}

	filter := bson.D{{"username", request.Username}}
	var findUser models.User
	err := database.Collection.FindOne(context.TODO(), filter).Decode(&findUser)

	if err != nil || len(findUser.Username) == 0 {
		const errorMsg string = "Error: User does not exist"
		log.Println(err)
		log.Println(findUser.Username)
		c.JSON(400, errorMsg)
		return
	}

	deleteUser, err := database.Collection.DeleteOne(context.TODO(), filter)
	if err != nil || deleteUser.DeletedCount != 1 {
		var errorMsg string = "Error: " + err.Error()
		log.Println(errorMsg)
		c.JSON(500, errorMsg)
		return
	}
	c.JSON(200, gin.H{})
}
