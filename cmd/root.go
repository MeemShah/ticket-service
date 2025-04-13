package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use:   "ticket-service",
		Short: "ticket-service server binary",
	}
)

func init() {
	RootCmd.AddCommand(serveRestCmd)
	RootCmd.AddCommand(TestRestCmd)
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
