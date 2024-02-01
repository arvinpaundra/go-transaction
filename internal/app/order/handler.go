package order

import (
	"clean-arch/internal/dto"
	"clean-arch/internal/factory"
	"clean-arch/pkg/tracer"
	"clean-arch/pkg/util"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(f *factory.Factory) *Handler {
	return &Handler{
		service: NewService(f),
	}
}

func (h *Handler) HandlerInsert(c *gin.Context) {
	var req dto.CreateOrderReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		errorMessage := gin.H{"errors": "please fill data"}
		if err != io.EOF {
			errors := util.FormatValidationError(err)
			errorMessage = gin.H{"errors": errors}
		}
		response := util.APIResponse("Create order failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}
	err = h.service.Insert(c, &req)
	if err != nil {
		response := util.APIResponse("Create order failed", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := util.APIResponse("Create order success", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
	return
}

func (h *Handler) HandlerFindAll(c *gin.Context) {
	orders, err := h.service.FindAll(c)
	if err != nil {
		response := util.APIResponse("Internal server error", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := util.APIResponse("Data orders", http.StatusOK, "success", orders)
	tracer.Log(c, "info", "Get all orders")
	c.JSON(http.StatusOK, response)
	return
}
