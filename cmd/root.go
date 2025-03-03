/*
Copyright © 2025 ARIA LOPEZ <aria.lopez.dev@proton.me>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/aria-afk/rediscli/redis"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

var (
	cfgFile string
	uri     string
)

func init() {
	// Config file bindings
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file to read from. Default is $HOME/.rediscli.[json/yaml/env/toml]")
	rootCmd.PersistentFlags().StringVar(&uri, "u", "redis://localhost:6379/0", "URI string for connection. Default is redis://localhost:6379/0")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		viper.AddConfigPath(home)
		viper.SetConfigName(".rediscli")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read in config file", err)
	}
}

// Builds redis opts object
// priority is cli flag -> config file -> defaults
func buildRedisOpts(cmd *cobra.Command) redis.RedisOpts {
	ro := redis.RedisOpts{}

	// URI
	uriFlag := cmd.PersistentFlags().Lookup("u")
	uriConf := viper.GetString("URI")
	if uriFlag.Changed {
		ro.URI = uriFlag.Value.String()
	} else if len(uriConf) > 0 {
		ro.URI = uriConf
	} else {
		ro.URI = uriFlag.Value.String()
	}

	return ro
}

// "Main" function to establish a connection to the redis client
// with the provided args and render the CLI.
func RunRedis(cmd *cobra.Command, args []string) {
	redisOpts := buildRedisOpts(cmd)
	_, err := redis.NewRedis(cmd.Context(), redisOpts)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
