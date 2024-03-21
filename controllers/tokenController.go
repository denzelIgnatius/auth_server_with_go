package controllers

import (
	"context"
	"log"

	"github.com/denzelIgnatius/auth_server_with_go/database"

	"github.com/denzelIgnatius/auth_server_with_go/encryption"

	"github.com/denzelIgnatius/auth_server_with_go/models"
	"github.com/denzelIgnatius/auth_server_with_go/validation"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAuth(c *gin.Context) {

	var request models.Request

	c.Bind(&request)

	if !validation.IsValidRequest(request) {
		const errorMsg string = "Error: Missing Input"
		log.Println(errorMsg)
		c.JSON(400, errorMsg)
		return
	}

	filter := bson.D{{Key: "username", Value: request.Username}}
	var result models.User
	err := database.Collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil || len(result.Username) == 0 {
		const errorMsg string = "Error: User does not exists"
		log.Println(errorMsg)
		c.JSON(400, errorMsg)
		return
	}

	var isValidPassword bool = encryption.CompareHashedPassword(result.Password, request.Password)

	c.JSON(200, isValidPassword)
}
