package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version string

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Returns the version of aws-role",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("aws-role " + Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
