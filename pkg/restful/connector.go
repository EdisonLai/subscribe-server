package restful

import (
	"context"
	"io"
	"net/http"

	"server/pkg/connector"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func ConnectorGet(c echo.Context) error {
	ctx := context.Background()

	data, err := connector.GetConnector(ctx)
	if err != nil {
		return NewResponse(c).WriteResponse(http.StatusBadRequest, nil)
	}

	// 将编码后的内容作为响应
	return c.String(http.StatusOK, data)
}

func ConnectorPatch(c echo.Context) error {
	ctx := context.Background()

	bodyReader := c.Request().Body
	buf, err := io.ReadAll(bodyReader)
	if err != nil {
		return NewResponse(c).WriteResponse(http.StatusBadRequest, err)
	}

	logrus.Info(string(buf))
	if err = connector.WriteConnector(ctx, string(buf)); err != nil {
		return NewResponse(c).WriteResponse(http.StatusBadRequest, err)
	}
	return NewResponse(c).WriteResponse(http.StatusOK, nil)
}
