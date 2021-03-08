package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-template/cmd/config"
	"net/http"
)

type handler struct {
	engine *gin.Engine
}

func (self *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	self.engine.ServeHTTP(w, req)
}

func (self handler) install(groupPath string) {
	baseGroup := self.engine.Group(groupPath)
	baseGroup.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "hello")
	})
}

func CreateHTTPAPIHandler(cfg config.Config) *handler {
	gin.SetMode(cfg.Mode)
	hd := &handler{
		engine: gin.Default(),
	}
	hd.engine.Use(middleware())
	hd.install("/api")
	return hd
}

func middleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		logrus.Info("middleware")
	}
}
