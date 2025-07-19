package transportcfg

import "github.com/10Narratives/ready-to-do/common/pkg/config/logging"

// Transport holds the transport configuration.
type Transport struct {
	GRPC GRPC `yaml:"grpc"`
}

// GRPC holds gRPC server configuration.
type GRPC struct {
	Host           string          `yaml:"host" env-required:"true"`
	Port           int             `yaml:"port" env-required:"true"`
	Timeout        string          `yaml:"timeout" env-default:"4s"`
	MaxRecvMsgSize int             `yaml:"max_recv_msg_size" env-default:"4194304"`
	MaxSendMsgSize int             `yaml:"max_send_msg_size" env-default:"4194304"`
	Reflection     bool            `yaml:"reflection" env-default:"true"`
	TLS            TLS             `yaml:"tls"`
	Logging        logging.Logging `yaml:"logging"`
}

// TLS holds TLS settings for gRPC.
type TLS struct {
	Enabled  bool   `yaml:"enabled" env-default:"false"`
	CertFile string `yaml:"cert_file"`
	KeyFile  string `yaml:"key_file"`
}
