package main

import (
	"log"

	"github.com/spf13/viper"
)

type WorkerConfig struct {
	SchedulerAddr string `mapstructure:"scheduler_address"`
	WorkerCnt     int    `mapstructure:"worker_count"`
}

func LoadAppConfig() WorkerConfig {
	res := WorkerConfig{}

	log.Println("Loading worker configurations...")

	viper.SetDefault("scheduler_address", "localhost:50051")
	viper.SetDefault("worker_count", 4)

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
	log.Printf("worker count: %v", res.WorkerCnt)

	return res
}
