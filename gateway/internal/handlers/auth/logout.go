package auth

import (
	"gateway/internal/client"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// LogouHandle ...
func LogouHandle(service *client.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	}
}
