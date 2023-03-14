package config

import "github.com/spf13/viper"

type External struct {
	GometrUrl string
}

func InitExternal(cfg *viper.Viper) *External {
	return &External{
		GometrUrl: cfg.GetString("gometrurl"),
	}
}

var ExternalConfig = new(External)
