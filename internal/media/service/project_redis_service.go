package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/charitan-go/media-server/ent"
	redispkg "github.com/charitan-go/media-server/pkg/redis"
)

const (
	SET_BY_ID_KEY = "media:id"
)

type MediaRedisService interface {
	SetById(ctx context.Context, p *ent.Media) error
	GetById(ctx context.Context, id string) (*ent.Media, error)
}

type mediaRedisServiceImpl struct {
	redisSvc redispkg.RedisService
}

func NewMediaRedisService(redisSvc redispkg.RedisService) MediaRedisService {
	return &mediaRedisServiceImpl{redisSvc}
}

func (svc *mediaRedisServiceImpl) getByIdKey(id string) string {
	return fmt.Sprintf("%s:%s", SET_BY_ID_KEY, id)
}

// SetById implements MediaRedisService.
func (svc *mediaRedisServiceImpl) SetById(ctx context.Context, p *ent.Media) error {
	key := svc.getByIdKey(p.ReadableID.String())

	value, err := json.Marshal(p)
	if err != nil {
		log.Printf("Cannot parse media to string: %v\n", p)
		return err
	}

	err = svc.redisSvc.Set(ctx, key, value)
	if err != nil {
		log.Printf("Cannot set to redis: %v\n", err)
		return err
	}

	return nil
}

func (svc *mediaRedisServiceImpl) GetById(ctx context.Context, id string) (*ent.Media, error) {
	key := svc.getByIdKey(id)

	result, err := svc.redisSvc.Get(ctx, key)
	if err != nil {
		log.Printf("Cannot get from redis: %v\n", err)
		return nil, err
	}

	// Parse result to *ent.Media
	p := &ent.Media{}
	err = json.Unmarshal([]byte(result), p)
	if err != nil {
		log.Printf("Cannot parse from string to *ent.Media: %v\n", err)
		return nil, err
	}

	return p, nil
}
