package handler

import (
	"github.com/gin-gonic/gin"
	"go-web-template/internal/app/server/config"
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
	hd.install("/api")
	return hd
}