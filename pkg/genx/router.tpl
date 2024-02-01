package {{.PackageName}}

import (
	"github.com/gin-gonic/gin"
)

// This function accepts gin.Routergroup to define a group route
// ex path "/api/v1"
func (h *handler) Router(g *gin.RouterGroup) {
    // Define a middleware here

    // Define a route here
}

// This function accepts gin.Routergroup to define a internal group route
// ex path "/in/api/v1"
func (h *handler) InRouter(g *gin.RouterGroup) {
    // Define a middleware here

    // Define a route here
}
