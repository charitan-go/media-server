package repository

import (
	"context"
	"log"

	"github.com/charitan-go/media-server/ent"
	"github.com/charitan-go/media-server/ent/media"
	"github.com/charitan-go/media-server/internal/media/dto"
	"github.com/charitan-go/media-server/pkg/database"
	"github.com/google/uuid"
)

type MediaRepository interface {
	// Save(mediaModel *model.Media) (*model.Media, error)
	Save(*dto.SaveMediaEntDto) (*ent.Media, error)

	FindAll(page, size int) ([]*ent.Media, error)
	FindOneByReadableId(readableId uuid.UUID) (*ent.Media, error)
}

type mediaRepositoryImpl struct {
	client *ent.Client
	// db *gorm.DB
}

func NewMediaRepository() MediaRepository {
	// db := database.DB
	client := database.Client
	if client == nil {
		log.Println("db is nil")
	} else {
		log.Println("db is not nil")
	}

	return &mediaRepositoryImpl{client}
}

func (r *mediaRepositoryImpl) Save(entDto *dto.SaveMediaEntDto) (*ent.Media, error) {
	media, err := r.client.Media.
		Create().
		SetName(entDto.Name).
		SetDescription(entDto.Description).
		SetGoal(entDto.Goal).
		SetStartDate(entDto.StartDate).
		SetEndDate(entDto.EndDate).
		SetCategory(entDto.Category).
		SetCountryCode(entDto.CountryCode).
		SetOwnerCharityReadableID(entDto.OwnerCharityReadableId).
		SetStatus(entDto.Status).
		Save(context.Background())

	if err != nil {
		return nil, err
	}

	return media, nil
}

// FindAll implements MediaRepository.
func (r *mediaRepositoryImpl) FindAll(page int, size int) ([]*ent.Media, error) {
	mediaList, err := r.client.Media.
		Query().
		Offset((page - 1) * size).
		Order(
			media.ByStartDate(),
		).
		All(context.Background())

	if err != nil {
		return nil, err
	}

	return mediaList, err
}

func (r *mediaRepositoryImpl) FindOneByReadableId(readableId uuid.UUID) (*ent.Media, error) {
	media, err := r.client.Media.
		Query().
		Where(media.ReadableID(readableId)).
		Only(context.Background())

	if err != nil {
		return nil, err
	}

	return media, nil
}

// func (r *mediaRepositoryImpl) FindOneByEmail(email string) (*model.Media, error) {
// 	var media model.Media
//
// 	result := r.db.Where("email = ?", email).First(&media)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
//
// 	return &media, nil
// }
//
// // FindOneByReadableId implements MediaRepository.
// func (r *mediaRepositoryImpl) FindOneByReadableId(readableId string) (*model.Media, error) {
// 	var media model.Media
//
// 	result := r.db.Where("readable_id = ?", readableId).First(&media)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
//
// 	return &media, nil
// }
