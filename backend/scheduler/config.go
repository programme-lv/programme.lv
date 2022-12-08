package main

import (
	"log"

	"github.com/spf13/viper"
)

type SchedulerConfig struct {
	APIPort       int    `mapstructure:"api_port"`
	SchedulerPort int    `mapstructure:"scheduler_port"`
	DBConnString  string `mapstructure:"db_conn_string"`
}

func LoadAppConfig() SchedulerConfig {
	res := SchedulerConfig{}

	log.Println("Loading worker configurations...")

	viper.SetDefault("api_port", 8080)
	viper.SetDefault("scheduler_port", 50051)
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
	log.Printf("scheduler's port: %v", res.SchedulerPort)
	log.Printf("DB connection string: %v", res.DBConnString)

	return res
}
