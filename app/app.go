package app

import (
	"log"

	"github.com/charitan-go/media-server/internal/media"
	"github.com/charitan-go/media-server/pkg/database"
	redispkg "github.com/charitan-go/media-server/pkg/redis"
	"github.com/charitan-go/media-server/rest"
	"github.com/charitan-go/media-server/rest/api"

	"go.uber.org/fx"
)

// Run both servers concurrently
func runServers(restSrv *rest.RestServer) {
	log.Println("In invoke")

	// Start REST server
	go func() {
		log.Println("In goroutine of rest")
		restSrv.Run()
	}()

	// Start RabbitMQ server
	// go func() {
	// 	log.Println("In goroutine of rabbitmq")
	// 	rabbitmqSrv.Run()
	// }()

}

func Run() {
	// Connect to db
	database.SetupDatabase()

	redispkg.SetupRedis()

	fx.New(
		media.MediaModule,
		redispkg.RedisModule,
		fx.Provide(
			rest.NewRestServer,
			rest.NewEcho,
			api.NewApi,
		),
		fx.Invoke(runServers),
	).Run()
}
