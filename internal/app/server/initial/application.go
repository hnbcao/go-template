package initial

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go-web-template/internal/app/server/config"
	"net/http"
)

type application struct {
	host    string
	port    int
	handler http.Handler
}

type Application interface {
	Run(ctx context.Context)
	InitHttpHandler(handler http.Handler)
}

func (app application) Run(ctx context.Context) {
	//ctx, cancel := context.WithCancel(ctx)
	app.startHttpServer(ctx)
}

func (app *application) InitHttpHandler(handler http.Handler) {
	app.handler = handler
}

func (app application) startHttpServer(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	addr := fmt.Sprintf("%s:%d", app.host, app.port)
	srv := http.Server{
		Addr:    addr,
		Handler: app.handler,
	}

	// 新开启goroutine,运行http server
	logrus.Info(fmt.Sprintf("http server start at %s", addr))
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			logrus.WithError(err).Error("http server get some error")
			cancel()
		}
	}()

	// 新开启goroutine,监听context状态
	select {
	case <-ctx.Done():
		_ = srv.Shutdown(ctx)
		logrus.Info("stop application ...")
	}
}

func NewApplication(config config.Config) Application {
	app := application{
		host: config.Host,
		port: config.Port,
	}
	return &app
}
