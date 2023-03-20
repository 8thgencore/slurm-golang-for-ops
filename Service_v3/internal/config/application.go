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
	IsHTTPS bool
	DbURL   string
}

func InitApplication(cfg *viper.Viper) *Application {
	return &Application{
		Host:    cfg.GetString("host"),
		Port:    portDefault(cfg),
		Name:    cfg.GetString("name"),
		Mode:    cfg.GetString("mode"),
		Domain:  cfg.GetString("domain"),
		IsHTTPS: cfg.GetBool("ishttps"),
		DbURL:   getDbURL(),
	}
}

var ApplicationConfig = new(Application)

func portDefault(cfg *viper.Viper) string {
	if cfg.GetString("port") == "" {
		return "8000"
	}

	return cfg.GetString("port")
}

//nolint:unused // in future
func isHTTPSDefault(cfg *viper.Viper) bool {
	if cfg.GetString("ishttps") == "" || cfg.GetBool("ishttps") == false {
		return false
	}

	return true
}
