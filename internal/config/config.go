package config

import (
	"github.com/spf13/viper"
	"gitlabee.chehejia.com/gopkg/lsego/pkg/cc"
	"gitlabee.chehejia.com/gopkg/lsego/pkg/log"
)

type Config struct {
	Debug bool   `yaml:"debug"`
	Addr  string `yaml:"addr"`

	DB *Database `yaml:"db,omitempty"`
}

func New() (*Config, error) {
	automaticEnv()
	if AppName != "" {
		log.Infof("config loading from apollo for the app %s", AppName)
		if err := cc.NewApollo(AppName, Env).Run(); err != nil {
			return nil, err
		}
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	if cfg.Addr == "" {
		cfg.Addr = ":8125"
	}

	return &cfg, nil
}
