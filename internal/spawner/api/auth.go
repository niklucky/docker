package api

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Auth struct {
	token string
}

func NewAuth(token string) *Auth {
	if token == "" {
		log.Fatalln("You need to provide security token")
	}
	return &Auth{
		token: token,
	}
}

func (a *Auth) ByToken(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Get the Basic Authentication credentials
		token := r.Header.Get("Authorization")
		if token == a.token {
			// Delegate request to the given handle
			h(w, r, ps)
		} else {
			// Request Basic Authentication otherwise
			w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}
