package handlers

import (
	"authentication/internal/store"
	jwthelper "authentication/pkg/jwt"
	"authentication/pkg/response"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/julienschmidt/httprouter"
)

// RefreshHandle ...
func RefreshHandle(s *store.Store) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		refreshToken := jwthelper.ExtractRefreshToken(r)

		token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				s.Logger.Errorf("Unexpected signing method. %v", token.Header["alg"])
				json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Unexpected signing method. %v", token.Header["alg"])})
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("REFRESH_SECRET")), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			s.Logger.Errorf("refresh token expired. err: %w", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("refresh token expired. err: %v", err)})
			return
		}

		if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			s.Logger.Errorf("can't parse token. err: %w", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("can't parse token. err: %v", err)})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			userID, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
			if err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				s.Logger.Errorf("eror while parsing token. err: %w", err)
				json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("eror while parsing token. err: %v", err)})
				return
			}

			role := fmt.Sprint(claims["role"])
			if err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				s.Logger.Errorf("eror while parsing token. err: %w", err)
				json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("eror while parsing token. err: %v", err)})
				return
			}

			tk, createErr := jwthelper.CreateToken(userID, role)
			if createErr != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				s.Logger.Errorf("can't create token. err: %w", err)
				json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("can't create token. err: %v", err)})
				return
			}

			c := http.Cookie{
				Name:     "Refresh-Token",
				Value:    tk.RefreshToken,
				HttpOnly: true,
			}

			http.SetCookie(w, &c)
			w.Header().Add("Access-Token", tk.AccessToken)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(response.Info{Messsage: "Successfully refreshed"})

		} else {
			w.WriteHeader(http.StatusUnauthorized)
			s.Logger.Error("refresh token is expired")
			json.NewEncoder(w).Encode(response.Error{Messsage: "refresh token is expired"})
		}

	}
}
