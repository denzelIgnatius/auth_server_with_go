package controllers

import (
	"context"
	"time"

	"github.com/denzelIgnatius/auth_server_with_go/database"
	"github.com/denzelIgnatius/auth_server_with_go/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddUsers(c *gin.Context) {
	var request struct {
		Username string
		Password string
	}
	c.Bind(&request)
	user := models.User{
		ID:        primitive.NewObjectID(),
		Username:  request.Username,
		Password:  request.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err := database.Collection.InsertOne(context.TODO(), user)
	if err != nil {
		panic(err)
	}
	c.JSON(200, gin.H{})
}
