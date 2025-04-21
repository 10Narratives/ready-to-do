package config

import (
	"flag"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	General GeneralConfig `yaml:"general"`
	Logging LoggingConfig `yaml:"logging"`
}

// GeneralConfig represents common options for each node in cluster.
type GeneralConfig struct {
	NodeID string `yaml:"node_id" env-required:"true"`
	Env    string `yaml:"environment" env-required:"true"`
}

// LoggingConfig represents logging options.
type LoggingConfig struct {
	Level         string `yaml:"level" env-default:"info"`
	LogDir        string `yaml:"log_dir" env-required:"true"`
	EnableConsole bool   `yaml:"enable_console" env-default:"false"`
}

// MustLoad loads the application configuration using a path which is gotten by flag or .env file.
func MustLoad() *Config {
	configPath := fetchConfigPath()
	return MustLoadFromPath(configPath)
}

// MustLoadFromPath loads the application configuration from the specified file path.
// It panics if the file is missing, unreadable, or contains invalid configuration.
func MustLoadFromPath(configPath string) *Config {
	if configPath == "" {
		panic("cannot read configuration: " + configPath + " is empty.")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("cannot read configuration: file does not exists")
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("cannot read configuration by " + configPath + ": error is occurred " + err.Error())
	}

	return &cfg
}

// fetchConfigPath retrieves the configuration file path, either from a command-line flag or an environment variable.
// If neither is provided, it attempts to load the path from a .env file.
func fetchConfigPath() string {
	var path string

	flag.StringVar(&path, "config", "", "path to configuration file")
	flag.Parse()

	if path == "" {
		err := godotenv.Load()
		if err != nil {
			panic("flag is not set and .env cannot be loaded")
		}
		path = os.Getenv("CONFIG_PATH")
	}

	return path
}
