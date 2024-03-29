package controllers

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/denzelIgnatius/auth_server_with_go/database"

	"github.com/denzelIgnatius/auth_server_with_go/encryption"

	"github.com/denzelIgnatius/auth_server_with_go/models"
	"github.com/denzelIgnatius/auth_server_with_go/validation"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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
		log.Println(err)
		c.JSON(400, errorMsg)
		return
	}

	var isValidPassword bool = encryption.CompareHashedPassword(result.Password, request.Password)

	if !isValidPassword {
		const errorMsg string = "Error: Invalid credentials"
		log.Println(errorMsg)
		c.JSON(400, errorMsg)
		return
	}
	expiration := time.Now().Add(time.Hour * 24 * 30)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Subject":    result.Username,
		"Expiration": expiration.Unix(),
	})

	tokenString, error := token.SignedString([]byte(os.Getenv("SECRET")))
	log.Println(tokenString)
	if error != nil {
		const errorMsg string = "Error: Token created failed with error"
		log.Println(error)
		c.JSON(500, errorMsg)
		return
	}

	c.JSON(200, gin.H{
		"token":      tokenString,
		"expiration": expiration,
	})
}

func ValidateAuth(c *gin.Context) {
	var request struct {
		Token string
	}
	c.Bind(&request)
	token, err := jwt.Parse(request.Token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		c.AbortWithError(500, err)
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		log.Println(claim["Expiration"])
		if float64(time.Now().Unix()) > claim["Expiration"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		filter := bson.D{{Key: "username", Value: claim["Subject"].(string)}}
		var result models.User
		err := database.Collection.FindOne(context.TODO(), filter).Decode(&result)
		if err != nil || len(result.Username) == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
