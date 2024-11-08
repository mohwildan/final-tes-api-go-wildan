package http_faq

import (
	"app/app/delivery/http/middleware"
	"app/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type faqRouteHandler struct {
	Usecase   domain.FAQAppUsecase
	Route     *gin.RouterGroup
	Middleware middleware.Middleware
}

func NewRouteHandler(route *gin.RouterGroup, middleware middleware.Middleware, u domain.FAQAppUsecase) {
	handler := &faqRouteHandler{
		Usecase:    u,
		Route:      route,
		Middleware: middleware,
	}

	handler.handleFAQRoutes("/faq")
}

func (h *faqRouteHandler) handleFAQRoutes(path string) {
	faq := h.Route.Group(path)
	faq.GET("/public/list", h.getAllFaqs)
	faq.GET("/list", h.Middleware.Auth(), h.getAllFaqs)
	faq.POST("/create",h.Middleware.Auth(), h.createFaq)
}

func (h *faqRouteHandler) createFaq(c *gin.Context) {
	var req domain.FaqRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := h.Usecase.CreateFaq(c.Request.Context(), req)
	c.JSON(response.Status, response)
}

func (h *faqRouteHandler) getAllFaqs(c *gin.Context) {
	response := h.Usecase.GetAllFaqs(c.Request.Context())
	c.JSON(response.Status, response)
}