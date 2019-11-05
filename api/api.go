package api

import (
	"encoding/json"
	"net/http"
	"netsoc/cloud/models"
	"netsoc/cloud/services/cloudcix"

	"github.com/Strum355/log"
	"github.com/go-chi/chi"
)

type API struct {
	CloudService cloudcix.CloudCIXService
}

func (a *API) Register(r chi.Router) {
	r.Get("/vm", func(w http.ResponseWriter, r *http.Request) {
		vms, err := a.CloudService.GetVMs()
		if err != nil {
			log.WithError(err).Error("Could not fetch VMs.")
			return
		}
		response := struct {
			Content []models.VM `json:"content"`
		}{vms}

		json.NewEncoder(w).Encode(response)
	})
}
