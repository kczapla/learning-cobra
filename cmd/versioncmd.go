package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print version string",
	Long:  "long string for version command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("learning-cobra v0.1 -- head")
	},
}
