package controllers

import (
	"net/http"

	"github.com/Rashad-Muntar/soundproof/config"
	"github.com/Rashad-Muntar/soundproof/models"
	"github.com/gin-gonic/gin"
)

func UpdateProfile(c *gin.Context) {
	var body struct {
		Name  string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required"`
	}
	var user models.User

    if err := config.DB.Where("id = ?", c.Param("userId")).First(&user).Error; err != nil {
        c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "record not found"})
        return
    }

    if err := c.BindJSON(&body); err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    updatedPost := models.User{Name: body.Name, Email: body.Email}
    config.DB.Model(&user).Updates(&updatedPost)
    c.JSON(http.StatusOK, gin.H{"data": user})

}

func GetProfile(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	config.DB.First(&user, id)
	c.JSON(200, &user)
}



