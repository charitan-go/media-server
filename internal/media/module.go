package media

import (
	"github.com/charitan-go/media-server/internal/media/handler"
	"github.com/charitan-go/media-server/internal/media/repository"
	"github.com/charitan-go/media-server/internal/media/service"
	"go.uber.org/fx"
)

var MediaModule = fx.Module("media",
	fx.Provide(
		handler.NewMediaHandler,
		service.NewMediaService,
		service.NewMediaRedisService,
		repository.NewMediaRepository,
	),
)
