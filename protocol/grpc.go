package protocol

import (
	"github.com/lifangjunone/restful-api-demo/apps"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/lifangjunone/restful-api-demo/conf"
)

// NewGRPCService todo
func NewGRPCService() *GRPCService {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Named("host")
	grpcServer := grpc.NewServer()

	return &GRPCService{
		svr: grpcServer,
		l:   sugar,
		c:   conf.C(),
	}
}

// GRPCService grpc服务
type GRPCService struct {
	svr *grpc.Server
	l   *zap.SugaredLogger
	c   *conf.Config
}

// Start 启动GRPC服务
func (s *GRPCService) Start() {
	// 装载所有GRPC服务
	apps.InitGrpc(s.svr)
	s.l.Infof("GRPC 服务监听地址: %s", s.c.App.GrpcAddr())
	// 启动HTTP服务
	lis, err := net.Listen("tcp", s.c.App.GrpcAddr())
	if err != nil {
		s.l.Errorf("listen grpc tcp conn error, %s", err)
		return
	}
	if err := s.svr.Serve(lis); err != nil {
		if err == grpc.ErrServerStopped {
			s.l.Info("service is stopped")
		}

		s.l.Error("start grpc service error, %s", err.Error())
		return
	}
}

// Stop 启动GRPC服务
func (s *GRPCService) Stop() error {
	s.svr.GracefulStop()
	return nil
}
