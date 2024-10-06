package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/hexley21/fixup/internal/common/auth_jwt"
	"github.com/hexley21/fixup/internal/common/middleware"
	"github.com/hexley21/fixup/internal/user/delivery/http/v1/auth"
	"github.com/hexley21/fixup/internal/user/delivery/http/v1/user"
	"github.com/hexley21/fixup/internal/user/service"
	"github.com/hexley21/fixup/internal/user/jwt/verifier"
	"github.com/hexley21/fixup/pkg/http/handler"
)

type RouterArgs struct {
	AuthService            service.AuthService
	UserService            service.UserService
	MiddlewareFactory      *middleware.MiddlewareFactory
	HandlerComponents      *handler.Components
	AccessJWTManager       auth_jwt.JWTManager
	RefreshJWTManager      auth_jwt.JWTManager
	VerificationJWTManager verifier.JWTManager
}

func MapV1Routes(args RouterArgs, router chi.Router) {
	authHandlerFactory := auth.NewFactory(
		args.HandlerComponents,
		args.AuthService,
	)

	userHandlerFactory := user.NewFactory(
		args.HandlerComponents,
		args.UserService,
	)

	accessJWTMiddleware := args.MiddlewareFactory.NewJWT(args.AccessJWTManager)
	onlyVerifiedMiddleware := args.MiddlewareFactory.NewAllowVerified(true)

	router.Route("/v1", func(r chi.Router) {
		auth.MapRoutes(args.MiddlewareFactory, authHandlerFactory, args.AccessJWTManager, args.RefreshJWTManager, args.VerificationJWTManager, r)
		user.MapRoutes(args.MiddlewareFactory, userHandlerFactory, accessJWTMiddleware, onlyVerifiedMiddleware, r)
	})
}
