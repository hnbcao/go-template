package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"go-web-template/internal/app/server/config"
	"go-web-template/internal/app/server/handler"
	"go-web-template/internal/app/server/initial"
	"go-web-template/internal/pkg/signal"
)

func main() {
	cfg, err := config.InitializeConfig()

	if err != nil {
		logrus.WithError(err).Fatalln("main: invalid configuration")
	}

	ctx := signal.WithContext(context.Background())

	apiHandler := handler.CreateHTTPAPIHandler(cfg)

	app := initial.NewApplication(cfg)
	app.InitHttpHandler(apiHandler)
	app.Run(ctx)
}
