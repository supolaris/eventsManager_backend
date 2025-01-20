package routes

import (
	"basicapis/models"
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
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error message": "error in binding login user data",
		})
	}
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"message": "user loggedin successfully",
		"users":   user,
	})
}
