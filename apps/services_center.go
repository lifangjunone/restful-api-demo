package apps

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 服务注册

var (
	ServicesCenter map[string]Service     = make(map[string]Service, 10)
	HttpApps       map[string]HttpService = make(map[string]HttpService, 10)
)

type Service interface {
	Name() string
	InitService()
}

type HttpService interface {
	Name() string
	Registry(g gin.IRouter)
	InitService()
	Config()
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

func Init() {
	for _, svc := range ServicesCenter {
		svc.InitService()
	}
	for _, app := range HttpApps {
		app.Config()
		app.InitService()
	}
}
