package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/Rashad-Muntar/soundproof/config"
	"github.com/Rashad-Muntar/soundproof/models"
	"github.com/Rashad-Muntar/soundproof/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Sign up       godoc
// @Summary      Create a new user
// @Description  Takes a User JSON and store in DB. Return message.
// @Tags         users
// @Produce      json
// @Param        user  body      models.User  true  "User JSON"
// @Success      200   {object}  models.User
// @Router       /auth [signup]
func Signup(c *gin.Context) {
	var body struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}
	user := models.User{Name: body.Name, Email: body.Email, Password: string(bcryptPassword)}
	newUser := config.DB.Create(&user)

	if newUser.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": newUser.Error,
		})
		return
	}
	c.JSON(200, &user)
}

func Login(c *gin.Context) {

	var body struct {
		Email    string
		Password string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	var user models.User
	config.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email is not found",
		})
		return
	}
	cryptErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if cryptErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	utils.SignAddress("0xfcb5128c24acfd307843fbe40a1cf8c12c5c7d5f2268bda66783e1129bb531e35f8cbfc27628e75cd8a55ff013fc51eb0d5ccecc87693bc77fc3bf2e248846961b")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(200, tokenString)
}
