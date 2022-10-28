package main

import (
	"fmt"
	"log"
	"os"

	"github.com/adrg/xdg"
	"github.com/spf13/viper"

	"github.com/bpaulin/keshiki/cmd"
)

func loadConfig() {
	// config format
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// config path
	configFilePath, err := xdg.ConfigFile("keshiki")
	if err == nil {
		log.Print(configFilePath)
		viper.AddConfigPath(configFilePath)
		os.Mkdir(configFilePath, os.ModePerm)
	}
	viper.AddConfigPath(".")

	// read or init config
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Print("no config file found")
			if err := viper.SafeWriteConfig(); err != nil {
				log.Print(err)
			}
		} else {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	}
}

func main() {
	loadConfig()
	cmd.Execute()
}
