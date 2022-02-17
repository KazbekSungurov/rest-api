package server

import (
	"gateway/internal/client"
	"gateway/internal/handlers/auth"
	"gateway/internal/handlers/middleware"
)

// ConfigureRouter ...
func (s *Server) ConfigureRouter() {
	s.Router.Handle("POST", "/login", middleware.IsLoggedIn(auth.LoginHandle(client.AuthLoginService)))
}
