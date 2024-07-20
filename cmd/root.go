package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "codebase",
	Short: "Codebase skeleton for golang",
	Long:  "Codebase skeleton for golang.",
}

func Run(args []string) error {
	RootCmd.SetArgs(args)
	return RootCmd.Execute()
}
