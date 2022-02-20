package pethandlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"user/domain/model"
	"user/internal/store"
	"user/pkg/response"

	"github.com/julienschmidt/httprouter"
)

// UpdatePet ...
func UpdatePet(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		req := &model.PetDTO{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Eror during JSON request decoding. Request body: %v, Err msg: %w", r.Body, err)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Eror during JSON request decoding. Request body: %v, Err msg: %v", r.Body, err)})
			return
		}

		err := s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Can't open DB. Err msg:%v.", err)
		}

		user, _ := s.User().FindByID(req.OwnerID)

		p, err := s.Pet().FindPetID(req.PetID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			s.Logger.Errorf("Cant find pet. Err msg:%v.", err)
			return
		}

		if req.Name != "" {
			p.Name = req.Name
		}

		if req.Type != "" {
			p.Type = req.Type
		}

		if req.Weight != 0 {
			p.Weight = req.Weight
		}

		if req.Diseases != "" {
			p.Diseases = req.Diseases
		}

		if user != nil {
			if user.UserID != req.OwnerID {
				p.Owner = *user
			}
		}

		if req.PetPhotoURL != "" {
			p.PetPhotoURL = req.PetPhotoURL
		}

		err = p.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Data is not valid. Err msg:%v.", err)
			return
		}

		err = s.Pet().Update(p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Can't update user. Err msg:%v.", err)
			return
		}

		s.Logger.Info("Update pet with id = %d", p.PetID)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response.Info{Messsage: fmt.Sprintf("Update pet with id = %d", p.PetID)})
	}
}
