package http_user

import (
	"app/app/delivery/http/middleware"
	usecase_user "app/app/usecase/user"

	"github.com/gin-gonic/gin"
)

type routeHandler struct {
	Usecase    usecase_user.UserAppUsecase
	Route      *gin.RouterGroup
	Middleware middleware.Middleware
}

func NewRouteHandler(route *gin.RouterGroup, middleware middleware.Middleware, u usecase_user.UserAppUsecase) {
	handler := &routeHandler{
		Usecase:    u,
		Route:      route,
		Middleware: middleware,
	}

	handler.handleAuthRoute("/user")
}

func (h *routeHandler) handleAuthRoute(path string) {
	auth := h.Route.Group(path)
	auth.GET("/list", h.list)
}

func (h *routeHandler) list(c *gin.Context) {
	response := h.Usecase.FindAllUser(c.Request.Context())

	c.JSON(response.Status, response)
}
