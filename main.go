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

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Print("no config file found")
		} else {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	}
}
func main() {
	loadConfig()
	cmd.Execute()
}
