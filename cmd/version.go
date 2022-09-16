package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"restful-api-demo/version"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "print version",
	Long:  "print project version info ",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version:", version.Version)
	},
}

func init() {
	RootCmd.AddCommand(VersionCmd)
}
