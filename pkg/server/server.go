package server

import (
	"context"
	"net/http"
	"time"

	"github.com/cleidison-barradas/korvia.api/pkg/router"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	httpSrv *http.Server
}

func New(handlers ...router.Register) (*Server, error) {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api")

	api.GET("/health",func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"status": "ok",
			"uptime": time.Now().Format("2006-01-02 15:04:05"),
		})
	})

	for _, h := range handlers {
		h.RegisterRoutes(api)
	}

	return &Server{
		router: router,
	}, nil
}

func (s *Server) Listen(add string) error {
	s.httpSrv = &http.Server{
		Addr: add,
		Handler: s.router,
	}

	errCh := make(chan error, 1)

	go func() {
		if err := s.httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errCh <- err
		}
	}()

	select {
	case err := <-errCh:
		return err
	case <-time.After(200 *time.Microsecond):
		return nil
	}
}

func (s *Server) Shutdown(ctx context.Context) error {
	if s.httpSrv == nil {
		return nil
	}
	return s.httpSrv.Shutdown(ctx)
}