package databasecfg

import "github.com/10Narratives/ready-to-do/common/pkg/config/logging"

type Database struct {
	Host     string          `yaml:"host" env-required:"true" env-default:"localhost"`
	Port     int             `yaml:"port" env-required:"true" env-default:"5432"`
	User     string          `yaml:"user" env-required:"true" env-default:"postgres"`
	Password string          `yaml:"password" env-required:"true" env-default:"secret"`
	DBName   string          `yaml:"dbname" env-required:"true" env-default:"mydb"`
	SSLMode  string          `yaml:"sslmode" env-default:"disable"`
	Logging  logging.Logging `yaml:"logging"`
}
