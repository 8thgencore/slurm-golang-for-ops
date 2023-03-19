package config

import "github.com/spf13/viper"

type External struct {
	GometrURL string
}

func InitExternal(cfg *viper.Viper) *External {
	return &External{
		GometrURL: cfg.GetString("gometrurl"),
	}
}

var ExternalConfig = new(External)
