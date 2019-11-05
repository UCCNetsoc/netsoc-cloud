package config

import "github.com/spf13/viper"

func loadDefaults() {
	viper.SetDefault("cloud.production", false)

	viper.SetDefault("cloud.http.port", 7070)

	viper.SetDefault("cloud.consul.host", "127.0.0.1:8500")
	viper.SetDefault("cloud.consul.token", "")
	viper.SetDefault("cloud.consul.serviceaddr", "127.0.0.1")
}
