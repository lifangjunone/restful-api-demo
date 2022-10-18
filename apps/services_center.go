package apps

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

// 服务注册

var (
	ServicesCenter map[string]Service     = make(map[string]Service, 10)
	HttpApps       map[string]HttpService = make(map[string]HttpService, 10)
	GrpcApps       map[string]GrpcService = make(map[string]GrpcService, 10)
)

type Service interface {
	Name() string
	InitService()
}

func GetService(name string) Service {
	return ServicesCenter[name]
}

type HttpService interface {
	Name() string
	Registry(g gin.IRouter)
	InitService(r *gin.Engine)
	Config()
}

type GrpcService interface {
	Name() string
	Registry(r *grpc.Server)
	Config()
}

func GetHttpService(name string) HttpService {
	return HttpApps[name]
}

func Registry(svc Service) {
	if _, ok := ServicesCenter[svc.Name()]; ok {
		panic(fmt.Sprintf("%s registry yet", svc.Name()))
	}
	ServicesCenter[svc.Name()] = svc
}

func HttpRegistry(h HttpService) {
	if _, ok := HttpApps[h.Name()]; ok {
		panic(fmt.Sprintf("%s registry yet", h.Name()))
	}
	HttpApps[h.Name()] = h
}

func GrpcRegistry(h GrpcService) {
	if _, ok := GrpcApps[h.Name()]; ok {
		panic(fmt.Sprintf("%s registry yet", h.Name()))
	}
	GrpcApps[h.Name()] = h
}

func LoadGinApps() (names []string) {
	for name := range HttpApps {
		names = append(names, name)
	}
	return
}

func LoadGrpcApps() (names []string) {
	for name := range GrpcApps {
		names = append(names, name)
	}
	return
}

func InitGin(r *gin.Engine) {
	for _, app := range HttpApps {
		app.Config()
		app.InitService(r)
	}
}

func InitGrpc(r *grpc.Server) {
	for _, app := range GrpcApps {
		app.Config()
		app.Registry(r)
	}
}

func InitImpl() {
	for _, svc := range ServicesCenter {
		svc.InitService()
	}
}
