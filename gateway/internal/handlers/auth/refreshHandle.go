package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"gateway/internal/client"
	"gateway/internal/client/auth"
	"gateway/pkg/response"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// RefreshHandle ...
func RefreshHandle(service *client.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		refreshTknCookie, err := r.Cookie("Refresh-Token")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("error occured while reading cookie. err: %v", err)})
			return
		}

		refreshService, err := auth.Refresh(context.WithValue(r.Context(), client.RefreshTokenCtxKey, refreshTknCookie.Value), service, r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("error during response getting from auth service. err: %v", err)})
		}

		cookies := refreshService.Cookies
		for _, cookie := range cookies {
			http.SetCookie(w, cookie)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response.Info{Messsage: "Successfully refreshed"})

	}
}
