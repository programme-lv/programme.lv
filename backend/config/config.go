package config

import (
	"log"

	"github.com/spf13/viper"
)

type SchedulerConfig struct {
	APIPort       int    `mapstructure:"api_port"`
	SchedulerPort int    `mapstructure:"scheduler_port"`
	DBConnString  string `mapstructure:"db_conn_string"`
	TasksDir      string `mapstructure:"tasks_dir"`
	PasswordSalt  string `mapstructure:"password_salt"`
}

func LoadAppConfig() SchedulerConfig {
	res := SchedulerConfig{}

	log.Println("Loading worker configurations...")

	viper.SetDefault("api_port", 8080)
	viper.SetDefault("scheduler_port", 50051)
	viper.SetDefault("db_conn_string", "database.db")
	viper.SetDefault("tasks_dir", "/srv/deikstra/tasks")

	viper.SetConfigFile("config.toml")
	viper.SetConfigFile(".env")

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
	log.Printf("DB connection string: idk lmao")
	log.Printf("tasks folder: %v", res.TasksDir)

	return res
}
