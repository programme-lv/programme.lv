package main

import (
	"log"

	"github.com/spf13/viper"
)

type WorkerConfig struct {
	SchedulerAddr  string `mapstructure:"scheduler_address"`
	WorkerStrength int    `mapstructure:"worker_strength"`
	WorkerName     string `mapstructure:"worker_name"`
}

func LoadAppConfig() WorkerConfig {
	res := WorkerConfig{}

	log.Println("loading worker configurations...")

	viper.SetDefault("scheduler_address", "localhost:50051")
	viper.SetDefault("worker_strength", 4)
	viper.SetDefault("worker_name", "test")

	viper.SetConfigFile("config.toml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = viper.Unmarshal(&res)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("scheduler address: %v", res.SchedulerAddr)
	log.Printf("worker strength: %v", res.WorkerStrength)
	log.Printf("worker name: %v", res.WorkerName)

	return res
}
