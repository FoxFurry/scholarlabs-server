/*
Package cmd
*/
package cmd

import (
	"net"
	"os"

	"github.com/FoxFurry/scholarlabs/services/user/internal/config"
	"github.com/FoxFurry/scholarlabs/services/user/internal/server"
	"github.com/FoxFurry/scholarlabs/services/user/internal/store"
	"github.com/FoxFurry/scholarlabs/services/user/proto"
	"github.com/caarlos0/env/v7"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go.elastic.co/ecslogrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const serviceName = "user"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gateway",
	Short: "Run HTTP(s) server for gateway service",
	Run: func(cmd *cobra.Command, args []string) {
		log := logrus.New()
		log.SetFormatter(&ecslogrus.Formatter{})

		cfg := config.Config{}

		if err := env.Parse(&cfg); err != nil {
			log.WithError(err).Fatal("failed to parse environment variables")
		}

		database, err := store.NewDataStore(cfg)
		if err != nil {
			log.WithError(err).Fatal("failed to create datastore")
		}

		user, err := server.New(cfg, database, log)
		if err != nil {
			log.WithError(err).Fatal("failed to create a user server")
		}

		grpcListener, err := net.Listen("tcp", cfg.Host)
		if err != nil {
			log.WithError(err).Fatal("failed to listen to a port")
		}

		var opts []grpc.ServerOption
		grpcServer := grpc.NewServer(opts...)

		proto.RegisterUserServer(grpcServer, user)
		reflection.Register(grpcServer)

		log.Infof("starting user server at: %s", cfg.Host)
		if err := grpcServer.Serve(grpcListener); err != nil {
			log.WithError(err).Fatal("failed to run user grpc server:")
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
