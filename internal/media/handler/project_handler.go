package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/charitan-go/media-server/internal/media/dto"
	"github.com/charitan-go/media-server/internal/media/service"
	restpkg "github.com/charitan-go/media-server/pkg/rest"
	"github.com/labstack/echo/v4"
)

type MediaHandler struct {
	svc service.MediaService
}

func (h *MediaHandler) CheckHealth() string {
	return "OK"
}

func NewMediaHandler(svc service.MediaService) *MediaHandler {
	return &MediaHandler{svc: svc}
}

func (h *MediaHandler) CreateMedia(c echo.Context) error {
	jwtPayload, err := restpkg.GetJwtPayload(c)
	if err != nil {
		log.Println("Not found header payload")
		return c.JSON(http.StatusNonAuthoritativeInfo, dto.ErrorResponseDto{Message: "Not authorized"})
	}

	req := new(dto.CreateMediaRequestDto)
	if err := c.Bind(req); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, dto.ErrorResponseDto{Message: "Invalid request body", StatusCode: http.StatusBadRequest})
	}

	res, errRes := h.svc.HandleCreateMediaRest(req, jwtPayload)
	if errRes != nil {
		return c.JSON(int(errRes.StatusCode), *errRes)
	}

	return c.JSON(http.StatusCreated, *res)
}

func (h *MediaHandler) GetMediaById(c echo.Context) error {
	mediaId := c.Param("mediaId")

	res, errRes := h.svc.HandleGetMediaByIdRest(mediaId)
	if errRes != nil {
		return c.JSON(int(errRes.StatusCode), *errRes)
	}

	return c.JSON(http.StatusCreated, *res)
}

func (h *MediaHandler) GetMedias(c echo.Context) error {
	pageStr, sizeStr := c.Param("page"), c.Param("size")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		if pageStr != "" {
			return c.JSON(http.StatusBadRequest, dto.ErrorResponseDto{Message: "Invalid request param", StatusCode: http.StatusBadRequest})
		}
		page = 1
	}

	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		if sizeStr != "" {
			return c.JSON(http.StatusBadRequest, dto.ErrorResponseDto{Message: "Invalid request param", StatusCode: http.StatusBadRequest})
		}
		size = 10
	}

	if page <= 0 || size <= 0 {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponseDto{Message: "Invalid request param", StatusCode: http.StatusBadRequest})
	}

	res, errRes := h.svc.HandleGetMediasRest(page, size)
	if errRes != nil {
		return c.JSON(int(errRes.StatusCode), *errRes)
	}

	return c.JSON(http.StatusCreated, res)
}
