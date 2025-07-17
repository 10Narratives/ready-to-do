package config

import (
	"github.com/10Narratives/ready-to-do/common/pkg/config/loader"
	"github.com/10Narratives/ready-to-do/common/pkg/config/logging"
	databasecfg "github.com/10Narratives/ready-to-do/server/internal/config/database"
	transportcfg "github.com/10Narratives/ready-to-do/server/internal/config/transport"
)

type Config struct {
	Transport transportcfg.Transport `yaml:"transport"`
	Database  databasecfg.Database   `yaml:"database"`
	Logging   logging.Logging        `yaml:"logging"`
}

var l = loader.ConfigLoader[Config]{}

func MustLoad() *Config {
	return l.MustLoad()
}

func MustLoadFromFile(path string) *Config {
	return l.MustLoadFromFile(path)
}
