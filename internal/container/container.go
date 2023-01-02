package container

import (
	"github.com/ShuffleBoy/goapp"
	"github.com/ShuffleBoy/goapp/logger"
	"github.com/my-epoch/article_service/internal/config"
	"github.com/my-epoch/article_service/internal/grpc_server"
)

func Initialize(app *goapp.App) {
	app.AddProvider(logger.NewLogger)
	app.AddProvider(config.NewConfigProvider)
	app.AddProvider(grpc_server.NewArticleGrpcServer)

	app.AddInvoker(grpc_server.StartGrpcServer)
}
