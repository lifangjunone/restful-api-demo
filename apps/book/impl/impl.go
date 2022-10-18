package impl

import (
	"database/sql"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	"github.com/lifangjunone/restful-api-demo/apps/book"
	"github.com/lifangjunone/restful-api-demo/conf"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	db  *sql.DB
	log logger.Logger
	// 实现类必须嵌套类这个类
	book.UnimplementedServiceServer
}

func (s *service) Config() error {
	db := conf.C().MySQL.GetDB()
	s.log = zap.L().Named(s.Name())
	s.db = db
	return nil
}

func (s *service) Name() string {
	return book.AppName
}

func (s *service) Registry(server *grpc.Server) {
	book.RegisterServiceServer(server, svr)
}

//
//func init() {
//	app.RegistryGrpcApp(svr)
//}
