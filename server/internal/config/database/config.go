// Package databasecfg provides configuration structures for PostgreSQL connection.
// It supports YAML configuration with environment variable overrides.
package databasecfg

import "time"

// Config represents the root database configuration structure.
type Config struct {
	Database Database `yaml:"database"`
}

// Database contains PostgreSQL connection configuration.
type Database struct {
	Host     string `yaml:"host" env-default:"localhost"`
	Port     int    `yaml:"port" env-default:"5432"`
	User     string `yaml:"user" env-default:"postgres"`
	Password string `yaml:"password" env-required:"true"`
	DBName   string `yaml:"dbname" env-default:"app_db"`
	SSLMode  string `yaml:"sslmode" env-default:"verify-full"`

	Pool     Pool     `yaml:"pool"`
	Timeouts Timeouts `yaml:"timeouts"`
}

// Pool contains connection pool settings.
type Pool struct {
	MaxConns          int           `yaml:"max_conns" env-default:"20"`
	MinConns          int           `yaml:"min_conns" env-default:"2"`
	MaxConnLifetime   time.Duration `yaml:"max_conn_lifetime" env-default:"30m"`
	MaxConnIdleTime   time.Duration `yaml:"max_conn_idle_time" env-default:"5m"`
	HealthCheckPeriod time.Duration `yaml:"health_check_period" env-default:"1m"`
}

// Timeouts contains database operation timeout settings.
type Timeouts struct {
	Connect time.Duration `yaml:"connect" env-default:"5s"`
	Query   time.Duration `yaml:"query" env-default:"30s"`
	Exec    time.Duration `yaml:"exec" env-default:"30s"`
}
