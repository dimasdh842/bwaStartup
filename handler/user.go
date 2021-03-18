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

	if true {

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
