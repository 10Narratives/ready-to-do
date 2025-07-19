package loader

import (
	"flag"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

// ConfigLoader loads configuration structs from files using cleanenv.
// Use MustLoad() to load from --config flag, or MustLoadFromFile(path) for a specific file.
// Panics if the file is missing or invalid.
type ConfigLoader[T any] struct{}

// ConfigLoader is a generic helper for loading configuration structs from YAML files using cleanenv.
// Use MustLoad() to load from a file specified by the --config flag, or MustLoadFromFile(path) to load from a specific file path.
// Both methods panic if the file is missing or invalid.
func (cl *ConfigLoader[T]) MustLoad() *T {
	var path string
	flag.StringVar(&path, "config", "", "path to configuration file")
	flag.Parse()

	return cl.MustLoadFromFile(path)
}

// MustLoadFromFile loads the config struct from the given file path using cleanenv.
// Panics if the file does not exist or is invalid.
func (cl *ConfigLoader[T]) MustLoadFromFile(path string) *T {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist: " + path)
	}

	var cfg T
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return &cfg
}
