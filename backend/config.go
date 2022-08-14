package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	API_PORT         string  `mapstructure:"api_port"`
	DB_ADDRESS       string  `mapstructure:"db_address"`
	JOURNEYS_FOLDER  string  `mapstructure:"journeys_folder"`
	STATIONS_FILE    string  `mapstructure:"stations_file"`
	STMT_COUNT_QUERY int     `mapstructure:"stmt_count_query"`
	MIN_JOURNEY_DIST float64 `mapstructure:"min_journey_dist"`
	MIN_JOURNEY_TIME int     `mapstructure:"min_journey_time"`
}

var ApiConfig *Config

func LoadApiConfig() {
	fmt.Println("Loading configurations.")
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&ApiConfig)
	if err != nil {
		panic(err)
	}
}
