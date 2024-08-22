package handlers

import (
	"project/handlers/render"
	"project/views/home"
	"net/http"

	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

func HandleHome(w http.ResponseWriter, r *http.Request) error {
	if checkCompleteUserAuth(r) {
		user, err := gothic.CompleteUserAuth(w, r)
		if err != nil {
			return err
		}
		return render.Template(w, r, home.Index(user))
	}
	user, ok := userFallback(r)
	if !ok {
		return render.Template(w, r, home.Index(goth.User{}))
	}
	return render.Template(w, r, home.Index(user))
}