package controllers

import (
	"net/http"

	"github.com/faanrm/CRUD-api/models"
	"github.com/faanrm/CRUD-api/services"
	"github.com/gin-gonic/gin"
)

// * : permet de d'acceder a la valeur pointé par un pointeur
type UserController struct {
	userService *services.UserServices
}

// & en Go est l'opérateur "adresse de". Il renvoie l'adresse mémoire d'une variable.
// Lorsque vous utilisez & devant une variable, il vous donne un pointeur vers cette variable.
func UserControllers() *UserController {
	return &UserController{
		userService: &services.UserServices{},
	}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var user models.Users
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := c.Param("userId")
	if userId == "" {
		newUser, err := uc.userService.CreateUser(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, newUser)
	} else {

		updatedUser, err := uc.userService.CreateUser(&user, userId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, updatedUser)
	}
}

func (uc *UserController) GetUser(c *gin.Context) {
	userId := c.Param("userId")
	var users []*models.Users
	var err error
	if userId == "" {
		users, err = uc.userService.GetUser(nil)
	} else {
		users, err = uc.userService.GetUser(nil, userId)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	userId := c.Param("userId")
	err := uc.userService.DeleteUser(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
