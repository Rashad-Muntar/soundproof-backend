package routes

import (
	"github.com/Rashad-Muntar/soundproof/controllers"
	"github.com/Rashad-Muntar/soundproof/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func UserRoute(router *gin.Engine) {
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/auth/signup", controllers.Signup)
	router.POST("/auth/login", controllers.Login)
	router.PUT("/user/:userId", middleware.AuthCheck, controllers.UpdateProfile)
	router.GET("/user/:id", middleware.AuthCheck, controllers.GetProfile)
}
