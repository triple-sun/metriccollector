package responses

import (
	"net/http"

	"github.com/go-chi/render"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Status  int `json:"-"`
}

func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.SetContentType(render.ContentTypeJSON)
	render.Status(r, e.Status)
	return nil
}

func NewErrorResponse(status int, err error) render.Renderer {
	return &ErrorResponse{
		Status:       status,
		Message:      err.Error(),
	}
}