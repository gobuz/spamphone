package configuration

import (
	"github.com/spf13/viper"
)

var Viper *viper.Viper

// Init init the package for loading configurations
// pathConfig can be empty which will cause the configuration package find env for parameters
func Init(configFileName string,pathConfigs []string, getENV bool, postCalls ...func() error) error {
	Viper = viper.New()
	Viper.SetConfigType("json") // Force using json to load configuration json later from db
	if len(pathConfigs) != 0 && len(configFileName) != 0 {
		Viper.SetConfigName(configFileName)
		for _, pathConfig := range pathConfigs {
			if len(pathConfig) != 0 {
				Viper.AddConfigPath(pathConfig)
			}
		}
	}
	if getENV {
		Viper.AutomaticEnv()
	}
	err := Viper.ReadInConfig()
	if err != nil {
		return err
	}
	for _, postCall := range postCalls {
		err = postCall()
		if err != nil {
			return err
		}
	}
	return nil
}
