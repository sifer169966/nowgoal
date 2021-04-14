package configs

import (
	"fmt"
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	ServerConfig
	DatabaseConfig
	Key
}
type Key struct {
	Public string
}
type ServerConfig struct {
	HTTPPort string
}

type DatabaseConfig struct {
	Postgres struct {
		Host     string
		Port     string
		Username string
		Password string
		DbName   string
	}
}

var config Config

func InitViper(path, env string) {
	//godotenv.Load(".env")
	switch env {
	case "local":
		viper.SetConfigName("local-config")
		break
	case "develop":
		godotenv.Load(".env")
		viper.SetConfigName("develop-config")
		break
	default:
		viper.SetConfigName("production-config")
		break
	}

	viper.AddConfigPath(path)
	viper.SetEnvPrefix("app_nowgoal_service")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file has changed: ", e.Name)
	})
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(config)
}

func GetViper() *Config {
	return &config
}
