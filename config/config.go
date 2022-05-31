package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// LoadConfig reads configuration from ./config/config.json file.
func LoadConfig() (cfg *viper.Viper, err error) {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}
	fmt.Print("\n Got the config file .")
	return viper.GetViper(), nil
}

// func GetConfig(cfg *viper.Viper, key string) string {
// 	// return cfg.GetString("DB_HOST")
// 	data := cfg.
// }
