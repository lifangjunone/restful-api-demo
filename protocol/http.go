package protocol

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lifangjunone/restful-api-demo/apps"
	"github.com/lifangjunone/restful-api-demo/conf"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type HttpService struct {
	server *http.Server
	l      *zap.SugaredLogger
	r      *gin.Engine
}

func NewHttpService() *HttpService {
	//gin.SetMode(gin.ReleaseMode)
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
	s.l.Infof("HTTP 服务监听地址: %s", s.server.Addr)
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
