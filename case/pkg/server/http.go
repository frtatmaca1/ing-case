package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type httpServer struct {
	Engine *gin.Engine
	server *http.Server
}

func NewHTTPServer(port int) *httpServer {
	router := gin.New()

	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{SkipPaths: []string{"/healthcheck"}}))
	router.Use(gin.Recovery())

	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, "Healthy")
	})

	httpSrv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	return &httpServer{Engine: router, server: httpSrv}
}

func (srv *httpServer) Run() error {
	return srv.server.ListenAndServe()
}

func (srv *httpServer) Close(ctx context.Context) error {
	return srv.server.Shutdown(ctx)
}
