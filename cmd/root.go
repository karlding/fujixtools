package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "fujixtools",
	Short: "fujixtools provides tooling for working with a Fujifilm X-mount camera",
	Long:  `A suite of tooling that simpliflies working with a Fujifilm X-mount camera.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
