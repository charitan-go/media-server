package api

import (
	"net/http"

	mediahandler "github.com/charitan-go/media-server/internal/media/handler"

	"github.com/labstack/echo/v4"
)

type Api struct {
	MediaHandler *mediahandler.MediaHandler
}

func NewApi(mediaHandler *mediahandler.MediaHandler) *Api {
	return &Api{
		MediaHandler: mediaHandler,
	}
}

func (*Api) HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
