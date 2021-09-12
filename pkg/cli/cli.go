/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	addCommand "github.com/awile/datamkr/pkg/cli/add"
	initCommand "github.com/awile/datamkr/pkg/cli/init"
	listCommand "github.com/awile/datamkr/pkg/cli/list"
	makeCommand "github.com/awile/datamkr/pkg/cli/make"
	"github.com/awile/datamkr/pkg/config"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "datamkr",
	Short: "A Command line tool for mock data generation",
	Long:  "A Command line tool for mock data generation",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func addCommands() {
	c, _ := config.NewDatamkrConfigFactory()
	rootCmd.AddCommand(initCommand.NewInitCmd(c))
	rootCmd.AddCommand(makeCommand.NewMakeCmd(c))
	rootCmd.AddCommand(addCommand.NewAddCmd(c))
	rootCmd.AddCommand(listCommand.NewListCmd(c))
}

func init() {
	cobra.OnInitialize(initConfig)

	addCommands()

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./.datamkr.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".datamkr" (without extension).
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".datamkr")
	}

	viper.AutomaticEnv() // read in environment variables that match

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
