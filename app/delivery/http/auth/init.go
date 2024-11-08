package http_auth

import (
	"app/app/delivery/http/middleware"
	"app/domain"

	"github.com/gin-gonic/gin"
)

type routeHandler struct {
	Usecase   domain.AuthAppUsecase 
	Route      *gin.RouterGroup
	Middleware middleware.Middleware
}

func NewRouteHandler(route *gin.RouterGroup, middleware middleware.Middleware, u domain.AuthAppUsecase) {
	handler := &routeHandler{
		Usecase:    u,
		Route:      route,
		Middleware: middleware,
	}

	handler.handleAuthRoute("/auth/admin")
}

func (h *routeHandler) handleAuthRoute(path string) {
	auth := h.Route.Group(path)
	auth.GET("/me", h.Middleware.Auth(), h.getMe)
	auth.POST("/register", h.register)
	auth.POST("/login", h.login)
}

func (h *routeHandler) register(c *gin.Context) {
	var req domain.RegisterUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response := h.Usecase.Register(c.Request.Context(), req)

	c.JSON(response.Status, response)
}

func (h *routeHandler) login(c *gin.Context) {
	var req domain.LoginUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response := h.Usecase.Login(c.Request.Context(), req)

	c.JSON(response.Status, response)
}

func (h *routeHandler) getMe(c *gin.Context) {

	response := h.Usecase.GetMe(c.Request.Context(), c.MustGet("token_data").(domain.JWTClaimUser))

	c.JSON(response.Status, response)
}