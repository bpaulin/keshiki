/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
)

// restCmd represents the rest command
var restCmd = &cobra.Command{
	Use:   "rest",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rest called") // rest client
		client := resty.New()
		resp, err := client.R().
			EnableTrace().
			Get("https://bitbucket.org/")

		// Explore response object
		fmt.Println("Response Info:")
		fmt.Println("  Error      :", err)
		fmt.Println("  Status Code:", resp.StatusCode())
		fmt.Println("  Status     :", resp.Status())
		fmt.Println("  Proto      :", resp.Proto())
		fmt.Println("  Time       :", resp.Time())
		fmt.Println("  Received At:", resp.ReceivedAt())
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(restCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// restCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// restCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
