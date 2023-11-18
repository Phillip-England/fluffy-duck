package main

import (
	"fluffy-duck/pkg/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
    	log.Fatalf("Error loading .env file: %v", err)
	}

	r := gin.Default()

	routes.MountRoutes(r)
	r.Run(":8080")
}
