package user

import (
	"clean-arch/internal/middleware"

	"github.com/gin-gonic/gin"
)

// This function accepts gin.Routergroup to define a group route
func (h *handler) Router(g *gin.RouterGroup) {
	g.GET("", h.FindAll)
	g.POST("/create", h.CreateUser)
	g.GET("/:ID", h.GetUserByID)
	g.PUT("/update/:ID", h.UpdateUser)
}

func (h *handler) InRouter(g *gin.RouterGroup) {
	g.Use(middleware.AuthorizeSignature())
	// define route here
}
