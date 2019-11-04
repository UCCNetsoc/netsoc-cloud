package main

import (
	"netsoc/cloud/config"
	"netsoc/cloud/services/cloudcix"
	"os"

	"github.com/Strum355/log"
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
	cloud_service := cloudcix.CloudCIXService{}
	cloud_service.CreateService()

	// Fetch projects
	projects, err := cloud_service.GetProjects()
	if err != nil {
		log.WithError(err).Error("Could not fetch projects")
		os.Exit(1)
	}
	log.WithFields(log.Fields{
		"num_projects": len(projects),
	}).Info("Projects fetched.")

	// Fetch VMs
	vms, err := cloud_service.GetVMs()
	if err != nil {
		log.WithError(err).Error("Could not fetch vms")
		os.Exit(1)
	}
	log.WithFields(log.Fields{
		"num_vms": len(vms),
	}).Info("VMs fetched.")

}
