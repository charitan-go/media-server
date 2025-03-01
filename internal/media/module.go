package media

import (
	"github.com/charitan-go/media-server/internal/media/handler"
	"go.uber.org/fx"
)

var MediaModule = fx.Module("media",
	fx.Provide(
		handler.NewMediaHandler,
	),
)
