package commoncfg

import (
	"flag"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

// ConfigLoader is a generic configuration loader for any type T.
// It provides methods to load configuration from files with panic on errors.
type ConfigLoader[T any] struct{}

func (cl *ConfigLoader[T]) MustLoad() *T {
	var path string
	flag.StringVar(&path, "config", "", "path to configuration file")
	flag.Parse()

	return cl.MustLoadFromFile(path)
}

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
