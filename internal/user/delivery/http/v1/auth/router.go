package auth

import (
	"github.com/hexley21/fixup/internal/common/middleware"
	"github.com/labstack/echo/v4"
)

func (h *authHandler) MapRoutes(e *echo.Group, refreshSecretKey string) *echo.Group {
	registerGroup := e.Group("/register")

	registerGroup.POST("/customer", h.registerCustomer)
	registerGroup.POST("/provider", h.registerProvider)

	e.POST("/refresh", h.refresh, middleware.JWT(refreshSecretKey))
	e.POST("/login", h.login)
	e.POST("/logout", h.logout)

	return registerGroup
}
