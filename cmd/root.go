// Copyright © 2017 Arthur Maury arthur03.maury@gmail.com
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
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-cli",
	Short: "A simple cli",
	Long:  "A simple cli",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) {},
}

var permanentViper = viper.New()
var customViper = viper.New()

const cliConfigName = "cli_config"
const cliConfigPath = "."

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.testeur.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	permanentViper.SetConfigName("cli_config")
	permanentViper.AddConfigPath(".")

	// TODO fix viper watch

	// err := permanentViper.ReadInConfig()
	// check(err)

	// permanentViper.WatchConfig()
	// fmt.Println("WATCHING")
	// permanentViper.OnConfigChange(func(e fsnotify.Event) {
	// 	fmt.Println("Saving config")
	// 	permanentViper.WriteConfig()
	// 	getCustomConfig()
	// })
	// getCustomConfig()

}

func getCustomConfig() {
	err := permanentViper.ReadInConfig()
	check(err)

	configPaths := permanentViper.GetStringSlice("configpaths")
	configName := permanentViper.GetString("configname")
	// TODO if not exist
	customViper.SetConfigName(configName)
	for _, path := range configPaths {
		customViper.AddConfigPath(path)
	}
	err = customViper.ReadInConfig()
	check(err)
}

func check(e error) {
	if e != nil {
		if _, ok := e.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("----------- Not Found Error")
		}
		fmt.Println(e)
	}
}
