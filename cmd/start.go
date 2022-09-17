package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"restful-api-demo/apps"
	_ "restful-api-demo/apps/service_registry"
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
		apps.Init()
		return errors.New("no flags find")
	},
}

func init() {
	StartCmd.PersistentFlags().StringVarP(&confFile, "config", "f", "etc/demo.toml",
		"demo config file")
	RootCmd.AddCommand(StartCmd)
}
