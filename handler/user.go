package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {

		errors := helper.FormatError(err)
		errorMessages := gin.H{"errors": errors}
		response := helper.APIResponse("Account Has Not Created", http.StatusUnprocessableEntity, "failed", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	formatter := user.FormatUser(newUser, "tokentoekntokentoken")
	response := helper.APIResponse("Account Has Created", http.StatusOK, "success", formatter)
	if err != nil {
		response := helper.APIResponse("Account Has not Created", http.StatusBadRequest, "failed", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, response)

}

func (h *userHandler) Login(c *gin.Context) {

	var input user.LoginUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatError(err)
		errorMessages := gin.H{"errors": errors}
		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)

	if err != nil {
		errorMessages := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormatUser(loggedinUser, "tokentokentoken")

	response := helper.APIResponse("Successfully Loggedin", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
