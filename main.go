package main

import (
	"fmt"
	"restful-api-demo/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
