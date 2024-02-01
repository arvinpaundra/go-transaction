package http

import (
	"clean-arch/internal/app/order"
	"clean-arch/internal/app/user"
	"clean-arch/internal/factory"
	"clean-arch/pkg/tracer"

	"github.com/gin-gonic/gin"
)

// Here we define route function for user Handlers that accepts gin.Engine and factory parameters
func NewHttp(g *gin.Engine, f *factory.Factory) {
	logger, err := tracer.InitLogger()
	if err != nil {
		panic(err)
	}
	if logger == nil {
		panic(err)
	}

	defer logger.Sync()

	// Here we use our custom logger middleware

	// Here we use logger middleware before the actual API to catch any api call from clients
	g.Use(tracer.LoggingMiddleware(logger))

	// Here we use the recovery middleware to catch a panic, if panic occurs recover the application witohut shutting it off
	g.Use(tracer.RecoverMiddleware(logger))

	//g.Use(gin.Logger())
	//g.Use(gin.Recovery())

	// Here we define a router group
	v1 := g.Group("/api/v1")
	// Here we register the route from user handler
	user.NewHandler(f).Router(v1.Group("/user"))
	order.NewHandler(f).RouterV1(v1.Group("/orders"))
}
