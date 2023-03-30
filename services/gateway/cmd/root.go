/*
Package cmd
*/
package cmd

import (
	"log"
	"os"

	"github.com/FoxFurry/scholarlabs/gateway/internal/config"
	"github.com/FoxFurry/scholarlabs/gateway/internal/server"
	"github.com/caarlos0/env/v7"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gateway",
	Short: "Run HTTP(s) server for gateway service",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Config{}

		if err := env.Parse(&cfg); err != nil {
			log.Fatalf("failed to load config")
		}

		gateway, err := server.New(cfg)
		if err != nil {
			log.Fatalf("failed to create a gateway server: %v", err)
		}

		gateway.Run()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
