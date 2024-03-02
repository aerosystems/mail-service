package config

import (
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

type Config struct {
	EmailProvider     string `mapstructure:"EMAIL_PROVIDER"`
	MailHogPort       int    `mapstructure:"MAILHOG_PORT"`
	MailHogDomain     string `mapstructure:"MAILHOG_DOMAIN"`
	MailHogDHost      string `mapstructure:"MAILHOG_HOST"`
	MailHogUsername   string `mapstructure:"MAILHOG_USERNAME"`
	MailHogPassword   string `mapstructure:"MAILHOG_PASSWORD"`
	MailHogEncryption string `mapstructure:"MAILHOG_ENCRYPTION"`
}

func NewConfig() *Config {
	var cfg Config
	viper.AutomaticEnv()
	executablePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	executableDir := filepath.Dir(executablePath)
	viper.SetConfigFile(filepath.Join(executableDir, ".env"))
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}
	return &cfg
}
