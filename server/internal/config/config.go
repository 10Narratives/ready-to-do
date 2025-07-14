package config

import (
	commoncfg "github.com/10Narratives/ready-to-do/server/internal/config/common"
	databasecfg "github.com/10Narratives/ready-to-do/server/internal/config/database"
	transportcfg "github.com/10Narratives/ready-to-do/server/internal/config/transport"
)

type Config struct {
	Transport transportcfg.Config `yaml:"transport"`
	Database  databasecfg.Config  `yaml:"database"`
	Logging   commoncfg.Logging   `yaml:"logging"`
}

var loader = commoncfg.ConfigLoader[Config]{}

func MustLoad() *Config {
	return loader.MustLoad()
}

func MustLoadFromFile(path string) *Config {
	return loader.MustLoadFromFile(path)
}
