// Copyright © 2017 Steve Nicholls <stevexnicholls@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
