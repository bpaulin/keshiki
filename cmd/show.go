/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/charmbracelet/glamour"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("show called")
		fmt.Println(viper.Get("a"))

		in, _ := ioutil.ReadFile("README.md")

		r, _ := glamour.NewTermRenderer(
			// detect background color and pick either the default dark or light theme
			glamour.WithAutoStyle(),
			// wrap output at specific width
			glamour.WithWordWrap(40),
		)

		out, _ := r.Render(string(in))
		fmt.Print(out)
	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
