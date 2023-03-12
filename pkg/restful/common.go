package restful

import (
	"github.com/labstack/echo"
)

type responseWriter struct {
	echo.Context
}

func NewResponse(e echo.Context) *responseWriter {
	return &responseWriter{e}
}

func (w *responseWriter) WriteResponse(status int, data interface{}) error {
	return w.JSON(status, data)
}
