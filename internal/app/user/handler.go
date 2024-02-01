package user

import (
	"clean-arch/internal/dto"
	"clean-arch/internal/factory"
	"clean-arch/pkg/tracer"
	"clean-arch/pkg/util"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service Service
}

func NewHandler(f *factory.Factory) *handler {
	return &handler{
		service: NewService(f),
	}
}

func (h *handler) FindAll(c *gin.Context) {
	// Call the service to get the data
	data, err := h.service.FindAll(c)
	if err != nil {
		response := util.APIResponse("Internal server error", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := util.APIResponse("Data user", http.StatusOK, "success", data)
	tracer.Log(c, "info", "Get all user")
	c.JSON(http.StatusOK, response)
	return
}

func (h *handler) CreateUser(c *gin.Context) {
	var input dto.InsertUserRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := gin.H{"errors": "please fill data"}
		if err != io.EOF {
			errors := util.FormatValidationError(err)
			errorMessage = gin.H{"errors": errors}
		}
		response := util.APIResponse("User created failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}
	err = h.service.CreateUser(c, input)
	if err != nil {
		response := util.APIResponse("User created failed", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := util.APIResponse("User create success", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
	return
}

func (h *handler) GetUserByID(c *gin.Context) {
	paramId := c.Param("ID")
	userId, _ := strconv.Atoi(paramId)
	user, err := h.service.GetById(c, userId)
	if err != nil {
		response := util.APIResponse("Data Not Found", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	if user.ID <= 0 {
		response := util.APIResponse("Data Not Found", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := util.APIResponse("Data User", http.StatusOK, "success", user)
	c.JSON(http.StatusOK, response)
	return
}

func (h *handler) UpdateUser(c *gin.Context) {
	var input dto.UpdateUserRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := gin.H{"errors": "please fill data"}
		if err != io.EOF {
			errors := util.FormatValidationError(err)
			errorMessage = gin.H{"errors": errors}
		}
		response := util.APIResponse("User created failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}
	paramId := c.Param("ID")
	userId, _ := strconv.Atoi(paramId)
	user, err := h.service.GetById(c, userId)
	if err != nil {
		response := util.APIResponse("Data Not Found", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	if len(user.Name) <= 0 {
		response := util.APIResponse("Data Not Found", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	err = h.service.UpdateUser(c, input, userId)
	if err != nil {
		response := util.APIResponse("User update failed", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := util.APIResponse("User update success", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
	return

}
