/*
Package cmd
*/
package cmd

import (
	"context"
	"net"
	"os"

	"github.com/FoxFurry/scholarlabs/services/environment/internal/config"
	"github.com/FoxFurry/scholarlabs/services/environment/internal/server"
	"github.com/FoxFurry/scholarlabs/services/environment/internal/store"
	"github.com/FoxFurry/scholarlabs/services/environment/proto"
	"github.com/caarlos0/env/v7"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go.elastic.co/ecslogrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const serviceName = "environment"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "environment",
	Short: "Run GRPC server for environment service",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		log := logrus.New()
		log.SetFormatter(&ecslogrus.Formatter{})
		log.ReportCaller = true

		cfg := config.Config{}

		if err := env.Parse(&cfg); err != nil {
			log.WithError(err).Fatal("failed to parse environment variables")
		}

		database, err := store.NewDataStore(cfg, log)
		if err != nil {
			log.WithError(err).Fatal("failed to create datastore")
		}

		environment, err := server.New(ctx, cfg, database, log)
		if err != nil {
			log.WithError(err).Fatal("failed to create a environment server")
		}

		grpcListener, err := net.Listen("tcp", cfg.Host)
		if err != nil {
			log.WithError(err).Fatal("failed to listen to a port")
		}

		var opts []grpc.ServerOption
		grpcServer := grpc.NewServer(opts...)

		proto.RegisterEnvironmentServer(grpcServer, environment)
		reflection.Register(grpcServer)

		log.Infof("starting environment server at: %s", cfg.Host)
		if err := grpcServer.Serve(grpcListener); err != nil {
			log.WithError(err).Fatal("failed to run environment grpc server:")
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
