package main

import (
	"fmt"

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
	fmt.Println("Hello World!")
	r := gin.Default()
	r.POST("/createUser", controllers.AddUsers)
	r.Run()
}
