package responses

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/rs/zerolog/log"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Massage string `json:"massage,omitempty"`
	Error   string `json:"error"`
}
type Response struct {
	Error *ErrorResponse `json:"error,omitempty"`
	Data  interface{}    `json:"data,omitempty"`
}

func (r2 Response) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func Error(w http.ResponseWriter, r *http.Request, errorCode int, err error, message string) {
	render.Status(r, errorCode)
	errs := render.Render(w, r, &Response{
		Error: &ErrorResponse{
			Code:    errorCode,
			Massage: message,
			Error:   err.Error(),
		},
	})
	if errs != nil {
		log.Error().Err(err).Msg("Cannot render error response.")
	}
	return
}
func Success(w http.ResponseWriter, r *http.Request, statusCode int, data interface{}) {
	render.Status(r, statusCode)
	err := render.Render(w, r, Response{
		Data: data,
	})
	if err != nil {
		log.Error().Err(err).Interface("data", data).Msg("Could not get response successfully ")
		return
	}
}
