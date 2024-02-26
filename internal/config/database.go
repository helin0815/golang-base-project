package config

type Database struct {
	DSN          string `yaml:"dsn"`
	MaxOpenConns int    `yaml:"maxOpenConns,omitempty"`
	MaxIdleConns int    `yaml:"maxIdleConns,omitempty"`
}
