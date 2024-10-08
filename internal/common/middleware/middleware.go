package middleware

import (
	"github.com/hexley21/fixup/pkg/http/binder"
	"github.com/hexley21/fixup/pkg/http/writer"
)

const(
	MsgInsufficientRights = "Insufficient rights"
	MsgUserIsVerified     = "User has to be not-verified"
	MsgUserIsNotVerified  = "User is not verified"

	MsgMissingAuthorizationHeader = "Authorization header is missing"
	MsgMissingBearerToken         = "Bearer token is missing"
)

type Middleware struct {
	binder binder.FullBinder
	writer writer.HTTPErrorWriter
}

func NewMiddleware(binder binder.FullBinder, writer writer.HTTPErrorWriter) *Middleware{
	return &Middleware{
		binder, writer,
	}
}
