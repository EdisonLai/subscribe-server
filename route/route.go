package route

import (
	"server/pkg/restful"

	"github.com/labstack/echo"
)

func SetRoute(e *echo.Echo) {
	e.GET("/server", restful.ConnectorGet)
	e.PATCH("/server", restful.ConnectorPatch)
}
