package server

import (
	"gateway/internal/client"
	"gateway/internal/handlers/auth"
	"gateway/internal/handlers/middleware"
)

// ConfigureRouter ...
func (s *Server) ConfigureRouter() {
	s.Router.Handle("POST", "/login", middleware.IsLoggedIn(auth.LoginHandle(client.AuthLoginService)))
	s.Router.Handle("POST", "/logout", auth.LogoutHandle(client.AuthLogoutService))
	s.Router.Handle("POST", "/registration", auth.RegistrationHandle(client.AuthRegistrationService))
	s.Router.Handle("POST", "/refresh", auth.RefreshHandle(client.AuthRefreshService))
}
