package auth

import (
	"encoding/json"
	"fmt"
	"gateway/internal/client"
	"gateway/internal/client/auth"
	"gateway/pkg/response"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// RegistrationHandle ...
func RegistrationHandle(service *client.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		RegistrationService, err := auth.Registration(r.Context(), service, r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("error during response getting from auth service. err: %v", err)})
		}

		var user *user
		if err := json.Unmarshal(RegistrationService.Body, &user); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("can't parse JSON. err: %v", err)})
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response.Info{Messsage: "User created!"})

	}
}
