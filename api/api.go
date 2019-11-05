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
	r.Use(a.authMiddleware)
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
	r.Get("/vm/{username}", func(w http.ResponseWriter, r *http.Request) {
		username := chi.URLParam(r, "username")
		vms, err := a.CloudService.GetVMs()
		if err != nil {
			log.WithError(err).Error("Could not fetch VMs.")
			return
		}
		user_vms := make([]models.VM, 0)
		for _, i := range vms {
			if i.Username == username {
				user_vms = append(user_vms, i)
			}
		}
		response := struct {
			Content []models.VM `json:"content"`
		}{user_vms}

		json.NewEncoder(w).Encode(response)
	})
}
