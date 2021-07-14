package controllers

import (
	"net/http"

	"github.com/irishconstant/ui/coockies"
)

// logout обрабывает попытку разлогиниться
func (h *DecoratedHandler) logout(w http.ResponseWriter, r *http.Request) {
	session, err := coockies.Store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Options.MaxAge = -1
	err = session.Save(r, w)

	user := coockies.GetUser(session)
	user.Authenticated = false

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)

}
