package controllers

import (
	"net/http"

	"github.com/irishconstant/ui/coockies"
	"github.com/irishconstant/ui/helpers"
)

//index обрабывате запросы к стартовой странице
func (h *DecoratedHandler) index(w http.ResponseWriter, r *http.Request) {
	session, err := coockies.Store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user := coockies.GetUser(session)
	/*
		if user.Authenticated {
			h.connection.GetUserAttributes(&user)
		}
	*/

	helpers.ExecuteHTML("index", "index", w, *user)
}
