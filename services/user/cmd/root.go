/*
Package cmd
*/
package cmd

import (
	"log"
	"net"
	"os"

	"github.com/FoxFurry/scholarlabs/services/user/internal/config"
	"github.com/FoxFurry/scholarlabs/services/user/internal/server"
	"github.com/FoxFurry/scholarlabs/services/user/internal/store"
	"github.com/FoxFurry/scholarlabs/services/user/proto"
	"github.com/caarlos0/env/v7"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

		database, err := store.NewDataStore(cfg)
		if err != nil {
			log.Fatalf("failed to create datastore: %s", err.Error())
		}

		gateway, err := server.New(cfg, database)
		if err != nil {
			log.Fatalf("failed to create a gateway server: %v", err)
		}

		grpcListener, err := net.Listen("tcp", cfg.Host)
		if err != nil {
			log.Fatalf("failed to listen to a port: %v", err)
		}

		var opts []grpc.ServerOption
		grpcServer := grpc.NewServer(opts...)

		proto.RegisterUserServer(grpcServer, gateway)
		reflection.Register(grpcServer)

		log.Printf("Starting user server at: %s", cfg.Host)
		if err := grpcServer.Serve(grpcListener); err != nil {
			log.Fatalf("failed to run grpc server: %v", err)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
