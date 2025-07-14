package commoncfg

// Logging contains logging configuration.
type Logging struct {
	Level  string `yaml:"level" env-default:"info"`    // Log level (debug, info, warn, error)
	Format string `yaml:"format" env-default:"json"`   // Log format (json, text)
	Output string `yaml:"output" env-default:"stdout"` // Output destination
}
