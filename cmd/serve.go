/*
The MIT License (MIT)

Copyright (c) 2018 Steve Nicholls

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stevexnicholls/next/internal"
	log "github.com/stevexnicholls/next/logger"
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

	viper.AutomaticEnv()

	serveCmd.PersistentFlags().StringVarP(&port, "port", "p", "localhost:3000", "port to listen on")
	serveCmd.PersistentFlags().StringVar(&storePath, "db_path", "", "path to store db file")
	serveCmd.PersistentFlags().StringVar(&storeBucket, "db_bucket", "", "name of bucket in db")
	serveCmd.PersistentFlags().StringVar(&apiKey, "api_key", "", "api key")

	viper.BindPFlag("port", serveCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("db_path", serveCmd.PersistentFlags().Lookup("db_path"))
	viper.BindPFlag("db_bucket", serveCmd.PersistentFlags().Lookup("db_bucket"))
	viper.BindPFlag("api_key", serveCmd.PersistentFlags().Lookup("api_key"))

	viper.SetDefault("port", "localhost:3000")
	viper.SetDefault("db_path", "data.db")
	viper.SetDefault("db_bucket", "bucket")
	viper.SetDefault("api_key", "")
	viper.SetDefault("log_file", "next.log")

	err := log.Setup()
	if err != nil {
		log.Fatal(err)
	}
}
