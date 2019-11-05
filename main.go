package main

import (
	"fmt"
	"net/http"
	"netsoc/cloud/api"
	"netsoc/cloud/config"
	"netsoc/cloud/services"
	"netsoc/cloud/services/cloudcix"

	"github.com/Strum355/log"
	"github.com/go-chi/chi"
	"github.com/spf13/viper"
)

func main() {
	// Load config
	config.Load()

	// Initialise logger
	if viper.GetBool("cloud.production") {
		log.InitJSONLogger(&log.Config{})
	} else {
		log.InitSimpleLogger(&log.Config{})
	}

	config.PrintSettings()

	// Initialise CloudCIX service
	cloudService := cloudcix.CloudCIXService{}
	err := cloudService.CreateService()
	if err != nil {
		log.WithError(err).Error("Could not initialise cloud service")
		return
	}

	// Initialise Consul service
	consulService := services.ConsulService{}

	// Register with Consul
	consulService.Setup()
	consulService.GetSharedSecret()
	log.WithFields(log.Fields{
		"token": viper.GetString("cloud.http.token"),
	}).Info("Token")

	// Initialise router
	log.Info("Initialising chi router")
	r := chi.NewRouter()

	// Initialise API
	log.Info("Registering API endpoints")
	api := api.API{CloudService: cloudService}
	api.Register(r)

	// Listen and serve HTTP
	log.WithFields(log.Fields{
		"port": viper.GetInt("cloud.http.port"),
	}).Info("Serving HTTP")
	err = http.ListenAndServe(":"+fmt.Sprint(viper.GetInt("cloud.http.port")), r)
	if err != nil {
		log.WithError(err).Error("Error serving HTTP")
	}
}
