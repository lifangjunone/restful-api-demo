package cmd

import (
	"errors"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "demo",
	Short: "demo backend api",
	Long:  "demo impl backend api",
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("no flags find")
	},
}
