package handler

import (
	"bwastartup/auth"
	"bwastartup/helper"
	"bwastartup/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
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

	if err != nil {
		response := helper.APIResponse("Account Has Not Created", http.StatusUnprocessableEntity, "failed", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
	}

	token, err := h.authService.GenerateToken(newUser.ID)

	if err != nil {
		response := helper.APIResponse("Account Has Not Created", http.StatusUnprocessableEntity, "failed", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
	}

	formatter := user.FormatUser(newUser, token)
	response := helper.APIResponse("Account Has Created", http.StatusOK, "success", formatter)
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
	token, err := h.authService.GenerateToken(loggedinUser.ID)
	if err != nil {
		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	formatter := user.FormatUser(loggedinUser, token)

	response := helper.APIResponse("Successfully Loggedin", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {

	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Email is not available", http.StatusUnprocessableEntity, "error", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)

	if err != nil {
		errorMessage := "Server Error"

		response := helper.APIResponse("Email is not available", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	metaMessage := "Email has been registered"

	if isEmailAvailable {
		metaMessage = "Email is Available"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) UploudAvatar(c *gin.Context) {

	file, err := c.FormFile("avatar")

	if err != nil {
		data := gin.H{"is_uplouded": false}
		response := helper.APIResponse("Failed to uploud avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// JWT
	userId := 1

	//images name
	path := fmt.Sprintf("images/%d-%s", userId, file.Filename)
	err = c.SaveUploadedFile(file, path)

	if err != nil {
		data := gin.H{"is_uplouded": false}
		response := helper.APIResponse("Failed to uploud avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userId, path)
	if err != nil {
		data := gin.H{"is_uplouded": false}
		response := helper.APIResponse("Failed to uploud avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uplouded": true}
	response := helper.APIResponse("Avatar succcesfully uplouded", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)

}
