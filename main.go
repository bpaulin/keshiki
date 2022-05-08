/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"fmt"
	"log"

	"github.com/adrg/xdg"
	"github.com/bpaulin/keshiki/cmd"
	"github.com/spf13/viper"
)

func loadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	configFilePath, err := xdg.ConfigFile("keshiki")
	if err != nil {
		log.Fatal(err)
	}
	viper.AddConfigPath(configFilePath)
	viper.AddConfigPath(".")

	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}
func main() {
	loadConfig()
	cmd.Execute()
}
