package routes

import (
	"github.com/Rashad-Muntar/soundproof/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine){
	router.POST("/auth/signup", controllers.Signup)
}