package config

import (
	databasecfg "github.com/10Narratives/ready-to-do/server/internal/config/database"
	transportcfg "github.com/10Narratives/ready-to-do/server/internal/config/transport"
)

type Config struct {
	Transport transportcfg.Config
	Database  databasecfg.Config
}
