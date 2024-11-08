package http_blog

import (
	"app/app/delivery/http/middleware"
	"app/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type blogRouteHandler struct {
	Usecase   domain.BlogAppUsecase
	Route     *gin.RouterGroup
	Middleware middleware.Middleware
}

func NewRouteHandler(route *gin.RouterGroup, middleware middleware.Middleware, u domain.BlogAppUsecase) {
	handler := &blogRouteHandler{
		Usecase:    u,
		Route:      route,
		Middleware: middleware,
	}

	handler.handleBlogRoutes("/blog")
}

func (h *blogRouteHandler) handleBlogRoutes(path string) {
	blog := h.Route.Group(path)
	blog.GET("/public/list", h.getAllBlogs)
	blog.GET("/list", h.Middleware.Auth(), h.getAllBlogs)
	blog.POST("/create", h.Middleware.Auth(), h.createBlog)
}

func (h *blogRouteHandler) createBlog(c *gin.Context) {
	var req domain.BlogRequest 
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := h.Usecase.CreateBlog(c.Request.Context(), req) 
	c.JSON(response.Status, response)
}

func (h *blogRouteHandler) getAllBlogs(c *gin.Context) {
	response := h.Usecase.GetAllBlogs(c.Request.Context())
	c.JSON(response.Status, response)
}
