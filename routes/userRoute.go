package routes

import (
	"github.com/Rashad-Muntar/soundproof/controllers"
	"github.com/Rashad-Muntar/soundproof/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/auth/signup", controllers.Signup)
	router.POST("/auth/login", controllers.Login)
	router.PUT("/user/:userId", middleware.AuthCheck, controllers.UpdateProfile)
	router.GET("/user/:userId",  middleware.AuthCheck, controllers.GetProfile)
}
