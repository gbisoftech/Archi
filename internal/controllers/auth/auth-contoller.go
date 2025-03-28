package auth

import (
	"net/http"

	"main/global/tokenutil"
	"main/internal/models"

	"github.com/gin-gonic/gin"

	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"bindingerror": err.Error()})
		return
	}

	if err := user.BeforeCreate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"before create error": err.Error()})
		return
	}

	if err := user.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"create error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func Login(c *gin.Context) {
	var loginRequest struct {
		Email    string `form:"email"`
		Password string `form:"password"`
	}

	if err := c.ShouldBind(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := models.GetUserByEmail(loginRequest.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := tokenutil.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
