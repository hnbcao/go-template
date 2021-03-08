package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"go-template/cmd/config"
	"go-template/internal/app/facade/handler"
	"go-template/internal/pkg/core/signal"
)

func main() {
	cfg, err := config.InitializeConfig()

	if err != nil {
		logrus.WithError(err).Fatalln("main: invalid configuration")
	}

	ctx := signal.WithContext(context.Background())

	apiHandler := handler.CreateHTTPAPIHandler(cfg)

	app := NewApplication(cfg)
	app.InitHttpHandler(apiHandler)
	app.Run(ctx)
}
