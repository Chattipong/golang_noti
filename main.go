package main

import (
	"fmt"
	"golang_noti/routes"

	// "net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	// "go.mongodb.org/mongo-driver/bson"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	r := gin.Default()
	routes.SetupRouter(r)

	myPort := os.Getenv("LIVE_PORT")
	r.Run(":" + myPort)
}
