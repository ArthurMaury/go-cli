// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage the config files or path",
	Long:  "Manage the config files or path",
	Run: func(cmd *cobra.Command, args []string) {
		getCustomConfig()
		if len(args) == 0 {
			showConfig()
		} else {
			for _, entry := range args {
				if value := customViper.Get(entry); value != nil {
					fmt.Println(entry, ":", customViper.Get(entry))
				} else {
					fmt.Println("L'entrée", entry, "n'existe pas dans le fichier de configuration")
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}

func showConfig() {
	err := customViper.ReadInConfig()
	check(err)

	fmt.Println("Using config file:", customViper.ConfigFileUsed())
	for key, value := range customViper.AllSettings() {
		fmt.Println(green(key), ":", value)
	}
}
