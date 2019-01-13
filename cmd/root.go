// Copyright © 2018 Kentaro Hiramatsu <kterada.0509sg@gmail.com>
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
	"log"
	"os"

	"github.com/kterada0509/datadog-terraformer/internal"
	middleware "github.com/kterada0509/datadog-terraformer/middleware/datadog"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile, accountName, appKey, apiKey string
var showVersion bool
var credential middleware.Credential

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "datadog-terraformer",
	Short: "Datadof terraformer command",
	Long:  "Datadof terraformer command",
	RunE: func(cmd *cobra.Command, args []string) error {
		if showVersion {
			internal.PrintVersion()
			return nil
		}
		if err := cmd.Usage(); err != nil {
			return err
		}
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.SetOutput(os.Stdout)
	if err := rootCmd.Execute(); err != nil {
		rootCmd.SetOutput(os.Stderr)
		rootCmd.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(NewCmdVersion())
	rootCmd.AddCommand(NewCmdMonitor())
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().BoolVarP(&showVersion, "version", "v", false, "show the version and exit")

	viper.SetEnvPrefix("datadog")
	viper.AutomaticEnv()
}

// initConfig reads ENV variables if set.
func initConfig() {
	appKey = viper.GetString("APP_KEY")
	apiKey = viper.GetString("api_key")
	creds, err := middleware.NewCredential(apiKey, appKey)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	credential = *creds
}
