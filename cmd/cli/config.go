package main

import (
	"log"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v2"

	"github.com/megaboy101/cloudless/pkg/user"
)

type Config struct {
	User *user.User
}

func SaveConfig() {
	config := viper.AllSettings()

	bytes, err := yaml.Marshal(config)
	if err != nil {
		log.Fatalf("Error marshalling config file: %v", err)
	}

	home, err := homedir.Dir()
	if err != nil {
		log.Fatalf("Error fetching home directory: %v", err)
	}

	file, err := os.Create(home + "/.cloudless.yaml")
	if err != nil {
		log.Fatalf("Error creating config file: %v", err)
	}
	defer file.Close()

	file.WriteString(string(bytes))
}
