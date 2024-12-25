package auth

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
	Role     string `json:"role"`
	KeySize  int    `json:"keySize"`
}

func RegisterUser(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := CreateUser(req.Username, req.Password, req.Email, req.Role, req.KeySize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func RegisterRoutes(r *gin.Engine) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", RegisterUser)
	}
}