package handlers

import (
	"authentication/domain/model"
	"authentication/internal/store"
	jwthelper "authentication/pkg/jwt"
	"authentication/pkg/response"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Login ...
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginHandle checkes login and password and returns user if validation was passed
func LoginHandle(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		req := &Login{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Eror during JSON request decoding. Request body: %v, Err msg: %w", r.Body, err)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Eror during JSON request decoding. Request body: %v, Err msg: %v", r.Body, err)})
			return
		}

		err := s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Can't open DB. Err msg: %w", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Can't open DB. Err msg: %v", err)})
			return
		}
		user, err := s.User().FindByEmail(req.Email)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Eror during checking users email or password. Err msg: %w", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Eror during checking users email or password. Err msg: %v", err)})
			return
		}

		err = model.CheckPasswordHash(user.Password, req.Password)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Eror during checking users email or password. Err msg: %w", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Eror during checking users email or password. Err msg: %v", err)})
			return
		}

		tk, err := jwthelper.CreateToken(uint64(user.UserID), string(user.Role))
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			s.Logger.Errorf("Eror during createing tokens. Err msg: %w", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Eror during createing tokens. Err msg: %v", err)})
			return
		}

		c := http.Cookie{
			Name:     "Refresh-Token",
			Value:    tk.RefreshToken,
			HttpOnly: true,
		}

		http.SetCookie(w, &c)

		w.Header().Set("Access-Token", tk.AccessToken)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}
}
