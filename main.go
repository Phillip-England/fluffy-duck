package main

import (
	"fluffy-duck/pkg/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load(".env")


	r := gin.Default()

	routes.MountRoutes(r)
	r.Run()
}
