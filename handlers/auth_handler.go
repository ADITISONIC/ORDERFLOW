package handlers

import (
	"net/http"
	"orderflow/dto"
	"orderflow/services"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {

	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	err := services.Register(
		req.Name,
		req.Email,
		req.Password,
	)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "User registered successfully",
	})
}

func Login(c *gin.Context) {

	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	token, err := services.Login(
		req.Email,
		req.Password,
	)

	if err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"token": token,
	})
}

func Profile(c *gin.Context) {

	userID, _ := c.Get("userID")

	c.JSON(200, gin.H{
		"success": true,
		"userID": userID,
	})
}