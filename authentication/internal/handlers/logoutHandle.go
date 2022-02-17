package handlers

import (
	"authentication/internal/store"
	jwthelper "authentication/pkg/jwt"
	"authentication/pkg/response"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// LogoutHandle ...
func LogoutHandle(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		_, err := jwthelper.ExtractTokenMetadata(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			s.Logger.Errorf("you are unauthorized. err: %w", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("you are unauthorized. err: %v", err)})
			return
		}

		c := http.Cookie{
			Name:     "Refresh-Token",
			Value:    "",
			HttpOnly: true,
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Access-Token", "")
		http.SetCookie(w, &c)
		json.NewEncoder(w).Encode(response.Info{Messsage: "Successfully logged out"})
	}
}