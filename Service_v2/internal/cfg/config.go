package cfg

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Cfg struct {
	Port             string
	DbName           string
	DbUser           string
	DbPass           string
	DbHost           string
	DbPort           string
	GometrServiceUrl string 
}

func LoadAndStoreConfig() Cfg {
	v := viper.New()
	v.SetEnvPrefix("SERV")
	v.SetDefault("PORT", "8086")
	v.SetDefault("DBUSER", "postgres")
	v.SetDefault("DBPASS", "postgres")
	v.SetDefault("DBHOST", "localhost")
	v.SetDefault("DBPORT", "5436")
	v.SetDefault("DBNAME", "db")
	v.SetDefault("GOMETRSERVICEURL", "http://localhost:8300")
	v.AutomaticEnv()

	var cfg Cfg

	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Panic(err)
	}

	return cfg
}

func (cfg *Cfg) GetDBString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbPort, cfg.DbName)
}

