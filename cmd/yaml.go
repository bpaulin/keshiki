/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// yamlCmd represents the yaml command
var yamlCmd = &cobra.Command{
	Use:   "yaml",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("yaml called")
		// Yaml load
		yfile, err := ioutil.ReadFile("config.yaml")

		if err != nil {
			log.Fatal(err)
		}

		data := make(map[interface{}]interface{})

		err2 := yaml.Unmarshal(yfile, &data)

		if err2 != nil {
			log.Fatalf("error: %v", err)
		}
		fmt.Printf("--- t:\n%v\n\n", data)
	},
}

func init() {
	rootCmd.AddCommand(yamlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// yamlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// yamlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
