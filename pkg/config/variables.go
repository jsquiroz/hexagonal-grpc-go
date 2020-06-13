package config

import (
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	config     Config
	onceConfig sync.Once
)

// Config provide global variables
type Config struct {
	GRPCPort   uint   `mapstructure:"grpc_port"`
	HTTPPort   uint   `mapstructure:"http_port"`
	DBName     string `mapstructure:"db_name"`
	DBPort     uint   `mapstructure:"db_port"`
	DBPassword string `mapstructure:"db_password"`
	DBUser     string `mapstructure:"db_user"`
	DBServer   string `mapstructure:"db_server"`
	CertPem    string `mapstructure:"cert_pem"`
	CertKey    string `mapstructure:"cert_key"`
}

// LoadVariables file configuration
func LoadVariables() Config {
	onceConfig.Do(loadConfig)
	return config
}

func loadConfig() {

	cfgFile := viper.GetString("config")

	if cfgFile == "" {
		log.Fatal("No such config path")
	}

	viper.SetConfigType("yml")
	viper.SetConfigFile(cfgFile)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		log.Infof("Using config file: %v", viper.ConfigFileUsed())
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("Could not load configuration file")
	}

}
