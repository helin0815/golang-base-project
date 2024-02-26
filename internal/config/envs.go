package config

import (
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/viper"
	"gitlabee.chehejia.com/gopkg/lsego/pkg/log"
)

var (
	_, b, _, _ = runtime.Caller(0)
	RootPath   = filepath.Join(filepath.Dir(b), "../")
)
var (
	Env     string
	AppName string
)

func automaticEnv() {
	// 默认使用.env文件，但是如果有环境变量，环境变量优先级更高
	if err := godotenv.Load(filepath.Join(RootPath, ".env")); err != nil {
		log.Debugf("load .env: %s", err)
	}

	// 如果有.env.local文件，则.env.local优先级比前两者更高
	if err := godotenv.Overload(filepath.Join(RootPath, ".env.local")); err != nil {
		log.Debugf("load .env.local: %s", err)
	}

	viper.AutomaticEnv()
	Env = viper.GetString("LI_ENV")
	AppName = viper.GetString("CHJ_APP_NAME")

	log.Infof("Current Env: %v", Env)
}
