package app

import (
	"context"
	"net/http"

	"gitlabee.chehejia.com/gopkg/lsego"
	"gitlabee.chehejia.com/gopkg/lsego/pkg/log"
	"go.uber.org/zap"

	"gitlabee.chehejia.com/k8s/liks-gitops/internal/config"
)

type Server struct {
	cfg *config.Config

	handler http.Handler
}

func NewServer(ctx context.Context, cfg *config.Config, h http.Handler) *Server {
	if cfg.Debug {
		log.SetDefaultLoggerLevel(zap.DebugLevel)
	}

	return &Server{
		cfg:     cfg,
		handler: h,
	}
}

func (s *Server) Run() error {
	server := lsego.New()
	server.RootHandle(s.handler)
	server.HTTPServe(s.cfg.Addr)
	return nil
}
