package auth

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/hexley21/fixup/internal/common/auth_jwt"
	"github.com/hexley21/fixup/internal/user/jwt/refresh_jwt"
	"github.com/hexley21/fixup/internal/user/jwt/verify_jwt"
)

func MapRoutes(
	h *Handler,
	refreshJWTMiddleware func(http.Handler) http.Handler,
	accessJwtManager auth_jwt.Manager,
	refreshJwtManager refresh_jwt.Manager,
	vrfJWTManager verify_jwt.Manager,
	router chi.Router,
) chi.Router {
	router.Route("/auth", func(r chi.Router) {
		r.Post("/register/customer", h.RegisterCustomer(vrfJWTManager))
		r.Post("/register/provider", h.RegisterProvider(vrfJWTManager))
		r.Post("/resend-confirmation", h.ResendConfirmationLetter(vrfJWTManager))

		r.With(refreshJWTMiddleware).Post("/refresh", h.Refresh(accessJwtManager))
		r.Post("/login", h.Login(accessJwtManager, refreshJwtManager))
		r.Post("/logout", h.Logout)

		r.Get("/verify", h.VerifyEmail(vrfJWTManager))
	})

	return router
}
