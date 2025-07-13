// Package transportcfg provides configuration structures for gRPC server and gateway.
// It supports YAML configuration with environment variable overrides.
package transportcfg

import "time"

// Config represents the root transport configuration structure.
type Config struct {
	GRPC    GRPC    `yaml:"grpc"`
	Gateway Gateway `yaml:"gateway"`
}

// Cors contains Cross-Origin Resource Sharing configuration.
type Cors struct {
	AllowedOrigins   []string      `yaml:"allowed_origins" env-required:"true"`
	AllowedMethods   []string      `yaml:"allowed_methods" env-required:"true"`
	AllowedHeaders   []string      `yaml:"allowed_headers" env-required:"true"`
	AllowCredentials bool          `yaml:"allow_credentials" env-default:"true"`
	MaxAge           time.Duration `yaml:"max_age" env-default:"12h"`
}

// Gateway contains HTTP gateway configuration for gRPC-JSON transcoding.
type Gateway struct {
	Host       string     `yaml:"host" env-required:"true"`
	Port       uint       `yaml:"port" env-required:"true"`
	HTTP       HTTP       `yaml:"http"`      // HTTP server settings
	Cors       Cors       `yaml:"cors"`      // CORS configuration
	Marshaler  Marshaler  `yaml:"marshaler"` // Response marshaling options
	GatewayTLS GatewayTLS `yaml:"tls"`       // TLS configuration
	Shutdown   Shutdown   `yaml:"shutdown"`  // Graceful shutdown settings
}

// GRPC contains gRPC server configuration.
type GRPC struct {
	Host                  string        `yaml:"host" env-required:"true"`
	Port                  uint          `yaml:"port" env-required:"true"`
	MaxConnectionAge      time.Duration `yaml:"max_connection_age" env-default:"30m"`
	MaxConnectionAgeGrace time.Duration `yaml:"max_connection_grace" env-default:"5m"`
	MaxConcurrentStreams  uint          `yaml:"max_concurrent_streams" env-default:"1000"`
	MaxRecvMsgSize        uint          `yaml:"max_recv_msg_size" env-default:"4194304"` // 4MB
	MaxSendMsgSize        uint          `yaml:"max_send_msg_size" env-default:"4194304"` // 4MB
	Keepalive             Keepalive     `yaml:"keepalive"`                               // Connection keepalive settings
	TLS                   GrpcTLS       `yaml:"tls"`                                     // TLS configuration
	Health                Health        `yaml:"health"`                                  // Health check settings
	Logging               Logging       `yaml:"logging"`                                 // Logging configuration
	Shutdown              Shutdown      `yaml:"shutdown"`                                // Graceful shutdown settings
}

// Health contains gRPC health check server configuration.
type Health struct {
	Enabled     bool   `yaml:"enabled" env-default:"true"`
	ServiceName string `yaml:"service_name" env-default:"ready_to_do"`
}

// HTTP contains HTTP server settings for the gateway.
type HTTP struct {
	ReadTimeout    time.Duration `yaml:"read_timeout" env-default:"15s"`
	WriteTimeout   time.Duration `yaml:"write_timeout" env-default:"15s"`
	IdleTimeout    time.Duration `yaml:"idle_timeout" env-default:"60s"`
	MaxHeaderBytes uint          `yaml:"max_header_bytes" env-default:"1048576"` // 1MB
}

// Keepalive contains gRPC keepalive enforcement policies.
type Keepalive struct {
	EnforcementPolicy EnforcementPolicy `yaml:"enforcement_policy"` // Client policy enforcement
	ServerParameters  ServerParameters  `yaml:"server_parameters"`  // Server keepalive settings
}

// EnforcementPolicy defines keepalive enforcement rules for clients.
type EnforcementPolicy struct {
	MinTime             time.Duration `yaml:"min_time" env-default:"10s"` // Minimum ping interval
	PermitWithoutStream bool          `yaml:"permit_without_stream" env-default:"true"`
}

// ServerParameters defines keepalive parameters for the server.
type ServerParameters struct {
	MaxConnectionIdle time.Duration `yaml:"max_connection_idle" env-default:"15m"`
	Time              time.Duration `yaml:"time" end-default:"2h"`     // Ping interval
	Timeout           time.Duration `yaml:"timeout" env-default:"20s"` // Ping timeout
}

// Logging contains logging configuration.
type Logging struct {
	Level  string `yaml:"level" env-default:"info"`    // Log level (debug, info, warn, error)
	Format string `yaml:"format" env-default:"json"`   // Log format (json, text)
	Output string `yaml:"output" env-default:"stdout"` // Output destination
}

// Marshaler controls JSON marshaling behavior.
type Marshaler struct {
	EmitDefaults bool `yaml:"emit_defaults" env-default:"false"` // Include zero values in JSON
	EnumsAsInts  bool `yaml:"enums_as_ints" env-default:"false"` // Serialize enums as integers
	OrigName     bool `yaml:"orig_name" env-default:"true"`      // Use original proto field names
}

// Shutdown contains graceful shutdown configuration.
type Shutdown struct {
	GracePeriod time.Duration `yaml:"grace_period" env-default:"30s"` // Wait for active connections
	Timeout     time.Duration `yaml:"timeout" env-default:"10s"`      // Force shutdown after
}

// GrpcTLS contains gRPC server TLS configuration.
type GrpcTLS struct {
	Enabled      bool   `yaml:"enabled" env-default:"true"`
	CertFile     string `yaml:"cert_file" env-required:"true"`      // Server certificate
	KeyFile      string `yaml:"key_file" env-required:"true"`       // Private key
	ClientCaFile string `yaml:"client_ca_file" env-required:"true"` // Client CA for mTLS
}

// GatewayTLS contains HTTP gateway TLS configuration.
type GatewayTLS struct {
	Enabled  bool   `yaml:"enabled" env-default:"true"`
	CertFile string `yaml:"cert_file" env-required:"true"` // Server certificate
	KeyFile  string `yaml:"key_file" env-required:"true"`  // Private key
}
