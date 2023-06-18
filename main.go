package main

import (

	"github.com/Rashad-Muntar/soundproof/config"
	"github.com/gin-gonic/gin"
	"github.com/Rashad-Muntar/soundproof/routes"
)

func init() {
	config.LoadInitializers()
	config.ConnectDB()
}

func main() {
	r := gin.New()
	routes.UserRoute(r)
	r.Run(":8080")
}
