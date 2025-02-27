package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App   AppConifg
	Mongo MongoConfig
	Log   LogConfig
}
type AppConifg struct {
	Env  string
	Host string
	Port int
}
type MongoConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}
type LogConfig struct {
	Level string
}

func Load() Config {
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Cannot find config.json file. Err: %s", err.Error())
		panic(err)
	}

	return Config{
		App: AppConifg{
			Host: viper.GetString("app.host"),
			Port: viper.GetInt("app.port"),
			Env:  viper.GetString("app.environtment"),
		},
		Mongo: MongoConfig{
			Host:     viper.GetString("database.host"),
			Port:     viper.GetInt("database.port"),
			Username: viper.GetString("database.username"),
			Password: viper.GetString("database.password"),
		},
		Log: LogConfig{
			Level: viper.GetString("log.level"),
		},
	}
}
