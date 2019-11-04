package config

import "github.com/spf13/viper"

func loadDefaults() {
	viper.SetDefault("cloud.production", false)
}
