package config

import (
	"fmt"
	log "service/pkg/logger"

	"github.com/spf13/viper"

	"io/ioutil"
	"os"
	"strings"
)

var cfgDatabase *viper.Viper
var cfgApplication *viper.Viper
var cfgExternal *viper.Viper

// Setup Config
func Setup(path string) {
	viper.SetConfigFile(path)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf(fmt.Sprintf("Read config file fail: %s", err.Error()))
	}

	// Replace environment variables
	err = viper.ReadConfig(strings.NewReader(os.ExpandEnv(string(content))))
	if err != nil {
		log.Fatalf(fmt.Sprintf("Parse config file fail: %s", err.Error()))
	}

	// Databse
	cfgDatabase = viper.Sub("settings.database")
	if cfgDatabase == nil {
		panic("config not found settings.database")
	}
	DatabaseConfig = InitDatabase(cfgDatabase)

	// Application
	cfgApplication = viper.Sub("settings.application")
	if cfgApplication == nil {
		panic("config not found settings.application")
	}
	ApplicationConfig = InitApplication(cfgApplication)

	// External
	cfgExternal = viper.Sub("settings.external")
	if cfgExternal == nil {
		panic("config not found settings.external")
	}
	ExternalConfig = InitExternal(cfgExternal)

}

func SetConfig(configPath string, key string, value interface{}) {
	viper.AddConfigPath(configPath)
	viper.Set(key, value)
	_ = viper.WriteConfig()
}
