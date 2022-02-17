package server

import (
	pethandlers "user/internal/handlers/petHandlers"
	usershandlers "user/internal/handlers/usersHandlers"
	"user/internal/store"
)

// ConfigureRouter ...
func (s *Server) ConfigureRouter() {

	s.Router.Handle("GET", "/pets", pethandlers.AllPetsHandler(store.New(s.Config)))
	s.Router.Handle("GET", "/pets/:id", pethandlers.GetPetByID(store.New(s.Config)))
	s.Router.Handle("DELETE", "/pets/delete/:id", pethandlers.DeletePets(store.New(s.Config)))
	s.Router.Handle("POST", "/pets/new", pethandlers.NewPet(store.New(s.Config)))
	s.Router.Handle("POST", "/pets/update", pethandlers.UpdatePet(store.New(s.Config)))

	s.Router.Handle("GET", "/users", usershandlers.AllUsersHandler(store.New(s.Config)))
	s.Router.Handle("GET", "/users/:id", usershandlers.GetUserByID(store.New(s.Config)))
	s.Router.Handle("DELETE", "/users/delete/:id", usershandlers.DeleteUser(store.New(s.Config)))
	s.Router.Handle("POST", "/users/update", usershandlers.UpdateUser(store.New(s.Config)))
	s.Router.Handle("POST", "/users/new", usershandlers.NewUser(store.New(s.Config)))

}
