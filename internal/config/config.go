package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"strings"
	"time"
)

type
(
	Config struct {
		Pg   PgConfig
		HTTP HTTPConfig
	}

	PgConfig struct {
		URI      string
		User     string
		Password string
		Host     string
		Port     int
		Name     string
		SSLMode  string
		Dialect  string
	}

	HTTPConfig struct {
		Port           int
		MaxHeaderBytes int
		ReadTimeout    time.Duration
		WriteTimeout   time.Duration
	}
)

func Init(configPath string) (*Config, error) {
	var config Config

	if err := loadFilesIntoViper(configPath); err != nil {
		return nil, err
	}

	if err := parseFiles(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func loadFilesIntoViper(path string) error {
	loadConfigFiles(path)
	loadEnvFile()

	return viper.ReadInConfig()
}

func parseFiles(cfg *Config) error {
	if err := parseConfigFiles(cfg); err != nil {
		return err
	}

	if err := parseEnvFile(cfg); err != nil {
		return err
	}

	return nil
}


func loadConfigFiles(path string) {
	configPath := "/home/pogremulllka/go/src/github.com/MuZaZaVr/notesService/config/"
	configName := strings.Split(path, "/")[1]	//setConfigName -> main

	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
}

func loadEnvFile() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error load .env file: %s", err)
	}
}

//@TODO Remove if section?
func parseConfigFiles(cfg *Config) error {

	// Parse Pg variables:
	if viper.IsSet("pg.databaseName") {
		if err := viper.UnmarshalKey("pg.databaseName", &cfg.Pg.Name); err != nil {
			return err
		}
	} else {
		log.Fatal("Pg.databaseName is config file has not specified")
	}

	if viper.IsSet("pg.databaseSslMode") {
		if err := viper.UnmarshalKey("pg.databaseSslMode", &cfg.Pg.SSLMode); err != nil {
			return err
		}
	} else {
		log.Fatal("Pg.databaseSslMode is config file has not specified")
	}

	if viper.IsSet("pg.databaseDialect") {
		if err := viper.UnmarshalKey("pg.databaseDialect", &cfg.Pg.Dialect); err != nil {
			return err
		}
	} else {
		log.Fatal("Pg.databaseDialect is config file has not specified")
	}

	// Parse Http variables:
	if viper.IsSet("http.port") {
		if err := viper.UnmarshalKey("http.port", &cfg.HTTP.Port); err != nil {
			return err
		}
	} else {
		log.Fatal("Http.port is config file has not specified")
	}

	if viper.IsSet("http.maxHeaderBytes") {
		if err := viper.UnmarshalKey("http.maxHeaderBytes", &cfg.HTTP.MaxHeaderBytes); err != nil {
			return err
		}
	} else {
		log.Fatal("Http.maxHeaderBytes is config file has not specified")
	}

	if viper.IsSet("http.readTimeout") {
		if err := viper.UnmarshalKey("http.readTimeout", &cfg.HTTP.ReadTimeout); err != nil {
			return err
		}
	} else {
		log.Fatal("Http.readTimeout is config file has not specified")
	}

	if viper.IsSet("http.writeTimeout") {
		if err := viper.UnmarshalKey("http.writeTimeout", &cfg.HTTP.WriteTimeout); err != nil {
			return err
		}
	} else {
		log.Fatal("Http.writeTimeout is config file has not specified")
	}

	return nil
}

func parseEnvFile(cfg *Config) error {
	viper.SetEnvPrefix("pg")

	if err := viper.BindEnv("user"); err != nil {
		return err
	}

	if err := viper.BindEnv("password"); err != nil {
		return err
	}

	if err := viper.BindEnv("host"); err != nil {
		return err
	}

	if err := viper.BindEnv("port"); err != nil {
		return err
	}

	setEnvValues(cfg)

	return nil
}

func setEnvValues(cfg *Config) {
	cfg.Pg.User = viper.GetString("user")
	cfg.Pg.Password = viper.GetString("password")
	cfg.Pg.Host = viper.GetString("host")
	cfg.Pg.Port = viper.GetInt("port")
}