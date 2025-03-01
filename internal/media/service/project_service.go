package service

import (
	"context"
	"log"
	"net/http"

	"github.com/charitan-go/media-server/ent"
	"github.com/charitan-go/media-server/internal/media/dto"
	"github.com/charitan-go/media-server/internal/media/repository"
	restpkg "github.com/charitan-go/media-server/pkg/rest"
	"github.com/google/uuid"
)

type MediaService interface {
	HandleCreateMediaRest(req *dto.CreateMediaRequestDto, jwtPayload *restpkg.JwtPayload) (*dto.MediaResponseDto, *dto.ErrorResponseDto)

	HandleGetMediaByIdRest(mediaId string) (*dto.MediaResponseDto, *dto.ErrorResponseDto)
	HandleGetMediasRest(page, size int) ([]*dto.MediaResponseDto, *dto.ErrorResponseDto)
}

type mediaServiceImpl struct {
	r        repository.MediaRepository
	redisSvc MediaRedisService
}

func NewMediaService(
	r repository.MediaRepository,
	redisSvc MediaRedisService,
) MediaService {
	return &mediaServiceImpl{r, redisSvc}
}

func (svc *mediaServiceImpl) toMediaResponseDto(p *ent.Media) *dto.MediaResponseDto {
	return &dto.MediaResponseDto{
		ReadableId:             p.ReadableID.String(),
		Name:                   p.Name,
		Description:            p.Description,
		Goal:                   p.Goal,
		StartDate:              p.StartDate.UnixMilli(),
		EndDate:                p.EndDate.UnixMilli(),
		Category:               string(p.Category),
		CountryCode:            p.CountryCode,
		Status:                 string(p.Status),
		OwnerCharityReadableId: p.OwnerCharityReadableID,
	}
}

// HandleCreateMediaRest implements MediaService.
func (svc *mediaServiceImpl) HandleCreateMediaRest(
	req *dto.CreateMediaRequestDto,
	jwtPayload *restpkg.JwtPayload,
) (*dto.MediaResponseDto, *dto.ErrorResponseDto) {
	ctx := context.Background()

	// Validation for date

	// Create media dto to save to db
	saveMediaEntDto := req.ToSaveMediaEntDto(jwtPayload)

	mediaDto, err := svc.r.Save(saveMediaEntDto)
	if err != nil {
		log.Printf("Cannot save media to db: %v\n", err)
		return nil, &dto.ErrorResponseDto{StatusCode: http.StatusInternalServerError, Message: "Cannot create media."}
	}

	// Cache to redis
	err = svc.redisSvc.SetById(ctx, mediaDto)
	if err != nil {
		log.Printf("Cannot save to redis: %v\n", err)
	}

	responseDto := svc.toMediaResponseDto(mediaDto)

	return responseDto, nil
}

func (svc *mediaServiceImpl) HandleGetMediaByIdRest(
	mediaIdStr string,
) (*dto.MediaResponseDto, *dto.ErrorResponseDto) {
	ctx := context.Background()

	log.Printf("Media id: %s", mediaIdStr)
	mediaId, err := uuid.Parse(mediaIdStr)
	if err != nil {
		log.Printf("Cannot convert str to id: %v\n", err)
		return nil, &dto.ErrorResponseDto{StatusCode: http.StatusBadRequest, Message: "Media not found."}
	}

	// TODO: Find in redis first
	p, err := svc.redisSvc.GetById(ctx, mediaIdStr)
	if err != nil {
		log.Printf("Cannot get by id from redis: %v\n", err)
	}

	// Can find string from redis
	if p != nil {
		log.Println("CACHED HIT!!!")
		responseDto := svc.toMediaResponseDto(p)
		return responseDto, nil
	}
	log.Println("CACHED MISS!!!")

	// Find media by id
	mediaDto, err := svc.r.FindOneByReadableId(mediaId)
	if err != nil {
		log.Printf("Cannot find media: %v\n", err)
		return nil, &dto.ErrorResponseDto{StatusCode: http.StatusBadRequest, Message: "Media not found."}
	}

	err = svc.redisSvc.SetById(ctx, mediaDto)
	if err != nil {
		log.Printf("Cannot save to redis: %v\n", err)
	}

	responseDto := svc.toMediaResponseDto(mediaDto)

	return responseDto, nil
}

// HandleGetMediasRest implements MediaService.
func (svc *mediaServiceImpl) HandleGetMediasRest(page int, size int) ([]*dto.MediaResponseDto, *dto.ErrorResponseDto) {
	// Query media
	mediaDtoList, err := svc.r.FindAll(page, size)
	if err != nil {
		log.Printf("Cannot query: %v\n", err)
		return nil, &dto.ErrorResponseDto{StatusCode: http.StatusBadRequest, Message: "Media not found."}
	}

	// Map to response dto
	responseDto := make([]*dto.MediaResponseDto, len(mediaDtoList))
	for i, dto := range mediaDtoList {
		responseDto[i] = svc.toMediaResponseDto(dto)
	}

	return responseDto, nil
}
