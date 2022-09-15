package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goctl",
	Short: "A simple Go cli",
	Long:  `A simple Go cli applicaton for experimenting with Go cobra`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
