package api

import (
	"net/http"

	"github.com/spf13/viper"
)

func (a *API) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Auth-Token")

		if token == "" {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		if token != viper.GetString("cloud.http.token") {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
