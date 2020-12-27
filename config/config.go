package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Create private data struct to hold config options.
type ConfigStruct struct {
	App struct {
		ENV                string `yaml:"env"`
		HostName           string `yaml:"hostname"`
		Port               string `yaml:"port"`
		TimeZone           string `yaml:"timezone"`
		BaseURL            string `yaml:"baseurl"`
		APIKeyDuration     string `yaml:"apikeyduration"`
		CSRFExpireDuration string `yaml:"csrfexpireduration"`
		BasePath           string `yaml:"basepath"`
	}

	DB struct {
		Driver            string `yaml:"driver"`
		Host              string `yaml:"host"`
		Port              string `yaml:"port"`
		Name              string `yaml:"name"`
		User              string `yaml:"user"`
		Password          string `yaml:"password"`
		Locale            string `yaml:"locale"`
		MaxOpenConnection int    `yaml:"maxopenconns"`
	}
}

// Create a new config instance.
var (
	Config *ConfigStruct
)

// Read the config file from the current directory and marshal
// into the conf config struct.
func getConf() *ConfigStruct {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.WatchConfig()
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Printf("%v", err)
	}

	conf := &ConfigStruct{}
	err = viper.Unmarshal(conf)
	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}

	return conf
}

// Initialization routine.
func init() {
	// Retrieve config options.
	Config = getConf()
}
