package routes

import (
	"basicapis/models"
	"basicapis/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func addUser(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error message": "error in binding user json",
		})
	}
	err = user.SaveUser()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error message": "error in saving user",
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "user saved",
		"user":    user,
	})
}

func getUsers(ctx *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error message": "error in getting users",
		})
	}
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"message": "user get sucessfully",
		"users":   users,
	})

}

func loginUser(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error message": "Could not parse requested data",
		})
		return
	}
	err = user.ValidateUser()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error message": "Could not authenticate user",
		})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error message": "Could not generate jwt token",
		})
		return
	}

	ctx.JSON(http.StatusBadRequest, gin.H{
		"message": "User logged in successfully",
		"token":   token,
	})

}
