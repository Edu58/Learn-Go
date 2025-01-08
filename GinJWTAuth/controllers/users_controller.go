package controllers

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Edu58/Learn-Go/GinJWTAuth/initializers"
	"github.com/Edu58/Learn-Go/GinJWTAuth/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var body struct {
		Name     string
		Email    string
		Password string
	}

	err := c.Bind(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Please provide an email and password",
		})

		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 13)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not hash password",
		})
		log.Fatalln("Could not hash ", body.Password)
	}

	new_user := models.User{Name: body.Name, Email: body.Email, Password: string(hashedPassword)}
	result := initializers.DB.Create(&new_user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create user",
		})
		log.Fatalf("Could not create %v", new_user)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": gin.H{
			"name":  body.Name,
			"email": body.Email,
		},
	})
}

func Login(c *gin.Context) {
	var body models.User
	err := c.Bind(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	var user models.User
	initializers.DB.Where("email = ?", body.Email).First(&user)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Invalid email or password",
		})
	}

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Email,
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	secret := os.Getenv("HMACSECRETKEY")
	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ := token.SignedString([]byte(secret))

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
