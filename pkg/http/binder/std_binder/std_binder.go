package std_binder

import (
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"

	"github.com/hexley21/fixup/pkg/http/binder"
	"github.com/hexley21/fixup/pkg/http/json"
	"github.com/hexley21/fixup/pkg/http/rest"
)

type standardBinder struct {
	JSONDeserializer json.Deserializer
}

func New(JSONDeserializer json.Deserializer) *standardBinder {
	return &standardBinder{
		JSONDeserializer,
	}
}

func (b *standardBinder) BindJSON(r *http.Request, i any) *rest.ErrorResponse {
	if r.ContentLength == 0 {
		return binder.ErrEmptyBody
	}

	ctype := r.Header.Get("Content-Type")
	if !strings.HasPrefix(ctype, "application/json") {
		return binder.ErrUnsupportedMediaType
	}

	if err := b.JSONDeserializer.Deserialize(r.Body, i); err != nil {
		return rest.NewInvalidArgumentsError(err)
	}

	return nil
}

func (b *standardBinder) BindMultipartForm(r *http.Request, maxSize int64) (*multipart.Form, *rest.ErrorResponse) {
	if !strings.HasPrefix(r.Header.Get("Content-Type"), "multipart/form-data") {
		return nil, binder.ErrUnsupportedMediaType
	}

	if err := r.ParseMultipartForm(maxSize); err != nil {
		return nil, rest.NewReadFileError(err)
	}

	return r.MultipartForm, nil
}

func (b *standardBinder) BindForm(r *http.Request) (url.Values, *rest.ErrorResponse) {
	if !strings.HasPrefix(r.Header.Get("Content-Type"), "application/x-www-form-urlencoded") {
		return nil, binder.ErrUnsupportedMediaType
	}

	if err := r.ParseForm(); err != nil {
		return nil, rest.NewInternalServerError(err)
	}

	return r.Form, nil
}
