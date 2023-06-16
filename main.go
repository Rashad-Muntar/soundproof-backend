package main

import (
	"github.com/Rashad-Muntar/soundproof/config"
	"github.com/gin-gonic/gin"
)

func init() {
config.LoadInitializers()
}

func main() {
r := gin.New()
r.Run(":8080")
}