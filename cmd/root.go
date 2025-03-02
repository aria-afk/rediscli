/*
Copyright © 2025 ARIA LOPEZ <aria.lopez.dev@proton.me>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rediscli",
	Short: "Feature rich redis-cli",
	Long:  "Feature rich redis-cli",
	Run:   RunRedis,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}

// "Main" function to establish a connection to the redis client
// with the provided args and render the CLI.
func RunRedis(cmd *cobra.Command, args []string) {
	fmt.Println("Hello world")
}
