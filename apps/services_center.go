package apps

import "fmt"

// 服务注册

var (
	ServicesCenter map[string]Service = make(map[string]Service, 10)
)

type Service interface {
	Name() string
	InitService()
}

func Registry(svc Service) {
	if _, ok := ServicesCenter[svc.Name()]; ok {
		panic(fmt.Sprintf("%s registry yet", svc.Name()))
	}
	ServicesCenter[svc.Name()] = svc
}

func Init() {
	for _, svc := range ServicesCenter {
		svc.InitService()
	}
}
