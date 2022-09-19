package protocol

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"restful-api-demo/apps"
	"restful-api-demo/conf"
	"time"
)

type HttpService struct {
	server *http.Server
	l      *zap.SugaredLogger
	r      *gin.Engine
}

func NewHttpService() *HttpService {
	r := gin.Default()
	server := &http.Server{
		ReadHeaderTimeout: 60 * time.Second,
		ReadTimeout:       60 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20,
		Addr:              conf.C().App.HttpAddr(),
		Handler:           r,
	}
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Named("host")

	return &HttpService{
		server: server,
		l:      sugar,
		r:      r,
	}
}

func (s *HttpService) Start() error {
	// 加载Handler,
	apps.InitGin(s.r)
	apps := apps.LoadedGinApps()
	s.l.Infof("loaded gin apps: %v", apps)
	if err := s.server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			s.l.Info("service stopped")
			return nil
		}
		return fmt.Errorf("start service error, %s", err.Error())
	}
	return nil
}

func (s *HttpService) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		s.l.Warnf("shut down http service error, %s", err)
		return err
	}
	return nil
}
