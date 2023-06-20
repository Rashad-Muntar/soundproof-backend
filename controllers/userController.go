package controllers

import (
	"encoding/hex"
	"net/http"

	"github.com/Rashad-Muntar/soundproof/config"
	"github.com/Rashad-Muntar/soundproof/models"
	"github.com/Rashad-Muntar/soundproof/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UpdateProfile
// @Summary      Update user profile
// @Description  Updates the profile information of a user.
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id path string true "User ID"
// @Param        body body models.User  true "User profile details"
// @Success      200 {object} models.User
// @Router       /user/{id} [put]
func UpdateProfile(c *gin.Context) {
	var body struct {
		gorm.Model
		Name      string `json:"name" binding:"required"`
		Email     string `json:"email" binding:"required"`
		Signature string `json:"public"`
	}
	var user models.User
	var SignatureKey []byte
	if err := config.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if body.Signature != "" {
		key, err := utils.SignAddress(body.Signature)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		SignatureKey = []byte(key)
	}

	updatedPost := models.User{Name: body.Name, Email: body.Email, Signature: hex.EncodeToString(SignatureKey)}
	config.DB.Model(&user).Updates(&updatedPost)
	c.JSON(http.StatusOK, gin.H{"data": user})
}


// GetUserbyID   godoc
// @Summary      Get single user by ID
// @Description  Returns the user whose ID value matches the userId.
// @Tags         users
// @Produce      json
// @Param        userId  path      string  true  "search user by userId"
// @Success      200  {object}  models.User
// @Router       /user/{userId} [get]
func GetProfile(c *gin.Context) {
	id := c.Param("userId")
	var user models.User
	config.DB.First(&user, id)
	c.JSON(200, &user)
}
