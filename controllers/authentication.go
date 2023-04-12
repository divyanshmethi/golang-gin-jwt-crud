package controllers

import (
	"net/http"

	"RESTAPI_Gin/helper"
	"RESTAPI_Gin/models"
	"github.com/gin-gonic/gin"
)

func Register(context *gin.Context) {
	var input models.AuthenticationInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.User{
		Username: input.Username,
		Password: input.Password,
	}
	savedUser, err := user.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusCreated, gin.H{"user": savedUser})
}

func Login(context *gin.Context) {
	var input models.AuthenticationInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := models.FindUserByUsername(input.Username)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = user.ValidatePassword(input.Password)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	jwt, err := helper.GenerateJWT(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusOK, gin.H{"jwt": jwt})
}
