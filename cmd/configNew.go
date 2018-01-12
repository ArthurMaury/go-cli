// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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
)

// configNewCmd represents the configNew command
var configNewCmd = &cobra.Command{
	Use:   "new",
	Short: "Creates a new config file",
	Long:  "Creates a new config file",
	Run: func(cmd *cobra.Command, args []string) {
		newConfigFile()
	},
}

func init() {
	configCmd.AddCommand(configNewCmd)
}

func newConfigFile() {

	//Sets the name of the config file
	fmt.Print("config file name: ")
	name, _ := readClean()
	permanentViper.Set("configname", name)

	//Sets the path to config file
	fmt.Print("config file path (.): ")
	path, _ := readClean()
	if path == "." || path == "" {
		path, _ = os.Getwd()
	}
	if paths := permanentViper.GetStringSlice("configpaths"); !contains(paths, path) {
		newpaths := append(paths, path)
		permanentViper.Set("configpaths", newpaths)
	}

	os.MkdirAll(path, 0777)
	if _, err := os.Stat(path + "/" + name + ".yaml"); os.IsNotExist(err) {
		os.Create(path + "/" + name + ".yaml")
	}

	permanentViper.WriteConfig()
}
