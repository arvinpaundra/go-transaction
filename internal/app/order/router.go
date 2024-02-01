package order

import "github.com/gin-gonic/gin"

func (h *Handler) RouterV1(g *gin.RouterGroup) {
	g.POST("", h.HandlerInsert)
}
