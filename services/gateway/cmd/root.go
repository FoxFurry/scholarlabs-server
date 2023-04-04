/*
Package cmd
*/
package cmd

import (
	"os"

	"github.com/FoxFurry/scholarlabs/services/gateway/internal/config"
	"github.com/FoxFurry/scholarlabs/services/gateway/internal/server"
	"github.com/FoxFurry/scholarlabs/services/user/client"
	"github.com/caarlos0/env/v7"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go.elastic.co/ecslogrus"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gateway",
	Short: "Run HTTP(s) server for gateway service",
	Run: func(cmd *cobra.Command, args []string) {
		log := logrus.New()
		log.SetFormatter(&ecslogrus.Formatter{})
		log.ReportCaller = true

		cfg := config.Config{}

		if err := env.Parse(&cfg); err != nil {
			log.WithError(err).Fatal("failed to parse environment variables")
		}

		user, err := client.NewUserClient(cfg.UserServiceBaseURL)
		if err != nil {
			log.WithError(err).Fatal("failed to connect to user service")
		}

		gateway, err := server.New(cfg, log, user)
		if err != nil {
			log.WithError(err).Fatal("failed to create a gateway server")
		}

		log.Infof("starting gateway server at: %s", cfg.Host)

		gateway.Run()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
