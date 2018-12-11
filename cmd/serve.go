package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stevexnicholls/next/internal"
)

var author string
var port, storePath, storeBucket, apiKey string

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start http server with configured api",
	Long:  `Starts a http server and serves the configured api`,
	Run: func(cmd *cobra.Command, args []string) {
		s, err := api.NewServer()
		if err != nil {
			panic(err)
		}

		s.Start()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.PersistentFlags().StringVarP(&port, "port", "p", "localhost:3000", "port to listen on")
	serveCmd.PersistentFlags().StringVar(&storePath, "db_path", "", "path to store db file")
	serveCmd.PersistentFlags().StringVar(&storeBucket, "db_bucket", "", "name of bucket in db")
	serveCmd.PersistentFlags().StringVar(&apiKey, "api_key", "", "api key")

	viper.BindPFlag("port", serveCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("db_path", serveCmd.PersistentFlags().Lookup("db_path"))
	viper.BindPFlag("db_bucket", serveCmd.PersistentFlags().Lookup("db_bucket"))
	viper.BindPFlag("api_key", serveCmd.PersistentFlags().Lookup("api_key"))

	viper.SetDefault("port", "localhost:3000")
	viper.SetDefault("db_path", "./data.db")
	viper.SetDefault("db_bucket", "bucket")
	viper.SetDefault("api_key", "")
}
