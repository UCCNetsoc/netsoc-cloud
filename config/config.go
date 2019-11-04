package config

import (
	"encoding/json"
	"strings"

	"github.com/Strum355/log"
	"github.com/spf13/viper"
)

func Load() {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	loadDefaults()
	viper.AutomaticEnv()
}

func PrintSettings() {
	settings := viper.AllSettings()
	settings["cloud"].(map[string]interface{})["password"] = "[password]"
	settings["cloud"].(map[string]interface{})["api_key"] = "[api_key]"

	out, _ := json.MarshalIndent(settings, "", "\t")
	log.Debug("config:\n" + string(out))
}
