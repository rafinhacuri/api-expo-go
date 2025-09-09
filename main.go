package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rafinhacuri/api-expo-go/db"
	"github.com/rafinhacuri/api-expo-go/routes"
)

func init() {
	godotenv.Load()

	if err := db.InitDB(true, os.Getenv("MONGO_USERNAME"), os.Getenv("MONGO_URL"), os.Getenv("MONGO_PASSWORD"), os.Getenv("MONGO_DB_NAME")); err != nil {
		log.Fatal("Erro ao conectar no banco de dados:", err)
	}
}

func main() {
	server := gin.Default()

	server.SetTrustedProxies([]string{"127.0.0.1", "::1"})

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
