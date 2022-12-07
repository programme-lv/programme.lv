package main

import (
	"log"

	"github.com/spf13/viper"
)

type WorkerConfig struct {
	APIPort      int    `mapstructure:"api_port"`
	WorkerPort   int    `mapstructure:"worker_port"`
	DBConnString string `mapstructure:"db_conn_string"`
}

func LoadAppConfig() WorkerConfig {
	res := WorkerConfig{}

	log.Println("Loading worker configurations...")

	viper.SetDefault("api_port", 8080)
	viper.SetDefault("worker_port", 50051)
	viper.SetDefault("db_conn_string", "data.db")

	viper.SetConfigFile("config.toml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = viper.Unmarshal(&res)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("scheduler's API port: %v", res.APIPort)
	log.Printf("scheduler's WORKER port: %v", res.WorkerPort)
	log.Printf("DB connection string: %v", res.DBConnString)

	return res
}
