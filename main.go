package main

import (
	"github.com/Rashad-Muntar/soundproof/config"
	_ "github.com/Rashad-Muntar/soundproof/docs"
	"github.com/Rashad-Muntar/soundproof/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadInitializers()
	config.ConnectDB()
}

// @title    Soundproof API
// @version         1.0
// @description     Soundproof service API in Go using Gin framework
// @host      localhost:8080
// @BasePath  /
func main() {
	r := gin.New()
	routes.UserRoute(r)
	r.Run(":8080")
}
