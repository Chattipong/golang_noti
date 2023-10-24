package controllers

import (
	"net/http"

	"golang_noti/dto"
	"golang_noti/repository"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserAPI struct {
	UserRepository repository.UserRepository
}

func NewUserAPI() UserAPI {
	return UserAPI{
		UserRepository: repository.UserRepo(),
	}
}

func (api UserAPI) GetAllUser(c *gin.Context) {

	result, err := api.UserRepository.GetAllUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"success": false,
			"data":    err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"success": true,
		"data":    "All User",
		"obj":     result})

}

func (api UserAPI) CreateUser(c *gin.Context) {
	getUser := dto.CreateUser{}
	if err := c.ShouldBindJSON(&getUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "data": err.Error(), "success": false})
		return
	}

	err := api.UserRepository.CreateUser(getUser)

	if err.Status == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"success": false,
			"data":    err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"success": true,
		"data":    "Create User success",
	})
}

func (api UserAPI) UpdateUserById(c *gin.Context) {

	userID, errUserId := primitive.ObjectIDFromHex(c.Param("id"))
	if errUserId != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusBadRequest,
			"success": false,
			"data":    "User Id incorrect",
		})
		return
	}

	getUser := dto.UpdateUser{}
	if errDto := c.ShouldBindJSON(&getUser); errDto != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "data": errDto.Error(), "success": false})
		return
	}
	errApi := api.UserRepository.UpdateUserById(userID, getUser)

	if errApi.Status == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"success": false,
			"data":    errApi.Message,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"success": true,
		"data":    "Update User success",
	})
}

func (api UserAPI) DeleteUserById(c *gin.Context) {
	userID, errUserId := primitive.ObjectIDFromHex(c.Param("id"))
	if errUserId != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"success": false,
			"data":    "User Id incorrect",
		})
		return
	}

	errApi := api.UserRepository.DeleteUserById(userID)
	if errApi.Status == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"success": false,
			"data":    errApi.Message,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"success": true,
		"data":    "Delete User success",
	})
}
