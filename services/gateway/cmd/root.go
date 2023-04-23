/*
Package cmd
*/
package cmd

import (
	"os"

	course "github.com/FoxFurry/scholarlabs/services/course/client"
	"github.com/FoxFurry/scholarlabs/services/gateway/internal/config"
	"github.com/FoxFurry/scholarlabs/services/gateway/internal/server"
	user "github.com/FoxFurry/scholarlabs/services/user/client"
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

		user, err := user.NewUserClient(cfg.UserServiceBaseURL)
		if err != nil {
			log.WithError(err).Fatal("failed to connect to user service")
		}

		course, err := course.NewCourseClient(cfg.CourseServiceBaseURL)
		if err != nil {
			log.WithError(err).Fatal("failed to connect to course service")
		}

		gateway, err := server.New(cfg, log, user, course)
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
