package cmd

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"restful-api-demo/apps"
	"restful-api-demo/apps/host/http"
	"restful-api-demo/apps/host/impl"
	"restful-api-demo/conf"
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
		// host service specify impl
		// service := impl.NewHostServiceImpl()
		apps.HostService = impl.NewHostServiceImpl()
		// 通过 Host Api Handler 提供 HTTP RestFul接口
		api := http.NewHostHandler()
		api.Config()
		g := gin.Default()
		api.Registry(g)
		g.Run(conf.C().App.HttpAddr())
		return errors.New("no flags find")
	},
}

func init() {
	StartCmd.PersistentFlags().StringVarP(&confFile, "config", "f", "etc/demo.toml",
		"demo config file")
	RootCmd.AddCommand(StartCmd)
}
