/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/FoxFurry/scholarlabs/services/common/db"
	"github.com/FoxFurry/scholarlabs/services/course/internal/config"
	"github.com/caarlos0/env/v7"
	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Config{}

		if err := env.Parse(&cfg); err != nil {
			log.Fatalf("failed to load config")
		}

		database, err := db.NewDB(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName)
		if err != nil {
			log.Fatalf("failed to connect to database: %s", err.Error())
		}

		err = db.Migrate(database, "migrations")
		if err != nil {
			log.Fatalf("failed to run migrations: %s", err.Error())
		}

		log.Println("Successfully ran migrations")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
