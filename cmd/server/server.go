package main

import (
	"github.com/ShuffleBoy/goapp"
	"github.com/my-epoch/article_service/internal/container"
)

func main() {
	app := goapp.NewApp()
	container.Initialize(app)
	app.Start()
}
