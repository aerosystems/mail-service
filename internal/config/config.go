package config

import (
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

type Config struct {
	Mode              string `mapstructure:"MODE"`
	EmailProvider     string `mapstructure:"EMAIL_PROVIDER"`
	MailhogPort       int    `mapstructure:"MAILHOG_PORT"`
	MailhogDomain     string `mapstructure:"MAILHOG_DOMAIN"`
	MailhogHost       string `mapstructure:"MAILHOG_HOST"`
	MailhogUsername   string `mapstructure:"MAILHOG_USERNAME"`
	MailhogPassword   string `mapstructure:"MAILHOG_PASSWORD"`
	MailhogEncryption string `mapstructure:"MAILHOG_ENCRYPTION"`
	BrevoApiKey       string `mapstructure:"BREVO_API_KEY"`
	FeedBackEmail     string `mapstructure:"FEEDBACK_EMAIL"`
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
