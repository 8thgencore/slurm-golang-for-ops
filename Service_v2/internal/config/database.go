package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Database struct {
	DbType   string
	Host     string
	Port     int
	Name     string
	Username string
	Password string
}

func InitDatabase(cfg *viper.Viper) *Database {
	return &Database{
		DbType:   cfg.GetString("dbtype"),
		Host:     cfg.GetString("host"),
		Port:     cfg.GetInt("port"),
		Name:     cfg.GetString("name"),
		Username: cfg.GetString("username"),
		Password: cfg.GetString("password"),
	}
}

var DatabaseConfig = new(Database)

func getDbUrl() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		DatabaseConfig.Username,
		DatabaseConfig.Password,
		DatabaseConfig.Host,
		DatabaseConfig.Port,
		DatabaseConfig.Name,
	)
}
