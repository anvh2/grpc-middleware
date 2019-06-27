package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

//RootCmd ...
var RootCmd = &cobra.Command{
	Use:   "Grpc middleware",
	Short: "Grpc Middleware",
}

//Execute ...
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
