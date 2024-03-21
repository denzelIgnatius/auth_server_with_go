package main

import (
	"github.com/denzelIgnatius/auth_server_with_go/database"

	"github.com/gin-gonic/gin"

	"github.com/denzelIgnatius/auth_server_with_go/controllers"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load(".env")
	database.ConnectToMongo()
}

func main() {
	r := gin.Default()
	r.POST("/registerUser", controllers.AddUsers)
	r.DELETE("/deleteUser", controllers.DeleteUsers)
	r.POST("/getAuthentication", controllers.GetAuth)
	r.Run()
}
