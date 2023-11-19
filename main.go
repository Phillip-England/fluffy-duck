package main

import (
	"fluffy-duck/pkg/database"
	"fluffy-duck/pkg/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load(".env")

	pool := database.ConnectDb()
	// database.Clear(pool)
	err := database.Init(pool)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	r := gin.Default()

	routes.MountRoutes(r, pool)
	r.Run()
}
