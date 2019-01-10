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
var port, tlsCertPath, tlsKeyPath, storePath, storeBucket, apiKey string

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
	serveCmd.PersistentFlags().StringVar(&tlsCertPath, "tls_cert", "", "path to tls certificate")
	serveCmd.PersistentFlags().StringVar(&tlsKeyPath, "tls_key", "", "path to tls private key")
	serveCmd.PersistentFlags().StringVar(&storePath, "store_path", "next.store", "path to keystore")
	serveCmd.PersistentFlags().StringVar(&storeBucket, "store_bucket", "key", "name of bucket in keystore")
	serveCmd.PersistentFlags().StringVar(&apiKey, "api_key", "", "api key")

	viper.BindPFlag("port", serveCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("tls_cert", serveCmd.PersistentFlags().Lookup("tls_cert"))
	viper.BindPFlag("tls_key", serveCmd.PersistentFlags().Lookup("tls_key"))
	viper.BindPFlag("store_path", serveCmd.PersistentFlags().Lookup("store_path"))
	viper.BindPFlag("store_bucket", serveCmd.PersistentFlags().Lookup("store_bucket"))
	viper.BindPFlag("api_key", serveCmd.PersistentFlags().Lookup("api_key"))

	serveCmd.MarkPersistentFlagRequired("tls_cert")
	serveCmd.MarkPersistentFlagRequired("tls_key")

	viper.SetDefault("port", "localhost:3000")
	viper.SetDefault("tls_cert", "")
	viper.SetDefault("tls_key", "")
	viper.SetDefault("store_path", "data.db")
	viper.SetDefault("store_bucket", "bucket")
	viper.SetDefault("api_key", "")
	viper.SetDefault("log_path", "next.log")

	err := log.Setup()
	if err != nil {
		log.Fatal(err)
	}
}
