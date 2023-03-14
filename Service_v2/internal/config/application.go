package config

import (
	"github.com/spf13/viper"
)

type Application struct {
	Host    string
	Port    string
	Name    string
	Mode    string
	Domain  string
	IsHttps bool
	DbUrl   string
}

func InitApplication(cfg *viper.Viper) *Application {
	return &Application{
		Host:    cfg.GetString("host"),
		Port:    portDefault(cfg),
		Name:    cfg.GetString("name"),
		Mode:    cfg.GetString("mode"),
		Domain:  cfg.GetString("domain"),
		IsHttps: cfg.GetBool("ishttps"),
		DbUrl:   getDbUrl(),
	}
}

var ApplicationConfig = new(Application)

func portDefault(cfg *viper.Viper) string {
	if cfg.GetString("port") == "" {
		return "8000"
	} else {
		return cfg.GetString("port")
	}
}

func isHttpsDefault(cfg *viper.Viper) bool {
	if cfg.GetString("ishttps") == "" || cfg.GetBool("ishttps") == false {
		return false
	} else {
		return true
	}
}
