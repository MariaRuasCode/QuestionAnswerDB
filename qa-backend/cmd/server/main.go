package main

import (
	"log"

	"github.com/MariaRuasCode/docanswers/internal/api"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	api.RegisterRoutes(r)

	log.Println("Listening on :8080...")
	r.Run(":8080")
}
