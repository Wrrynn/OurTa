package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int     `json:"id"`
	Username string  `json:"username"`
	Password string  `json:"-"`
	Saldo    float64 `json:"saldo"`
}

var users = map[string]*User{
	"user1": {ID: 1, Username: "user1", Password: "password1", Saldo: 0.0},
	"user2": {ID: 2, Username: "user2", Password: "password2", Saldo: 0.0},
}

func main() {
	router := gin.Default()

	router.POST("/login", loginHandler)
	router.POST("/topup/:id", topupHandler)
	router.GET("/users", getUser)

	router.Run(":8080")
}

func loginHandler(c *gin.Context) {
	var requestBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user, ok := users[requestBody.Username]; ok && user.Password == requestBody.Password {

		c.JSON(http.StatusOK, gin.H{"message": "Login berhasil", "user_id": user.ID})
	} else {

		c.JSON(http.StatusUnauthorized, gin.H{"error": "Login gagal"})
	}
}

func topupHandler(c *gin.Context) {
	userID := c.Param("id")

	if user, ok := users[userID]; ok {
		var topupAmount float64
		if err := c.ShouldBindJSON(&topupAmount); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user.Saldo += topupAmount

		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Top-up sebesar %.f berhasil", topupAmount), "saldo_terkini": user.Saldo})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pengguna tidak ditemukan"})
	}
}

func getUser(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}
