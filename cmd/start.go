package cmd

import (
	"github.com/lifangjunone/restful-api-demo/apps"
	_ "github.com/lifangjunone/restful-api-demo/apps/service_registry"
	"github.com/lifangjunone/restful-api-demo/conf"
	"github.com/lifangjunone/restful-api-demo/protocol"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

var (
	confFile string
)

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "start demo backend api",
	Long:  "start demo impl backend api",
	RunE: func(cmd *cobra.Command, args []string) error {
		// load config
		err := conf.LoadConfigFromToml(confFile)
		if err != nil {
			panic(err)
		}
		// r := gin.Default()
		// apps.Init(r)
		apps.InitImpl()
		httpSvc := NewManager()
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP, syscall.SIGINT)
		go httpSvc.WaitStop(ch)
		go httpSvc.grpc.Start()
		go httpSvc.restful.Start()
		return httpSvc.Start()
	},
}

func NewManager() *manager {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Named("CLI")
	return &manager{
		restful: protocol.NewRestfulService(),
		http:    protocol.NewHttpService(),
		grpc:    protocol.NewGRPCService(),
		l:       sugar,
	}
}

type manager struct {
	restful *protocol.RestfulService
	http    *protocol.HttpService
	grpc    *protocol.GRPCService
	l       *zap.SugaredLogger
}

func (m *manager) Start() error {
	return m.http.Start()
}

func (m *manager) WaitStop(ch <-chan os.Signal) {
	for v := range ch {
		switch v {
		case syscall.SIGTERM:
			m.l.Infof("received signal【sigterm】: %s", v)
		case syscall.SIGQUIT:
			m.l.Infof("received signal【sigquit】: %s", v)
		case syscall.SIGHUP:
			m.l.Infof("received signal【sighup】: %s", v)
		case syscall.SIGINT:
			m.l.Infof("received signal【sigint】: %s", v)
		default:
			m.l.Infof("received signal: %s", v)
		}
		m.grpc.Stop()
		m.http.Stop()
		m.restful.Stop()
	}
}

func init() {
	StartCmd.PersistentFlags().StringVarP(&confFile, "config", "f", "etc/demo.toml",
		"demo config file")
	RootCmd.AddCommand(StartCmd)
}
