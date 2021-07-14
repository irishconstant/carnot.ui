package controllers

import (
	"net/http"

	"github.com/irishconstant/ui/coockies"
	"github.com/irishconstant/ui/helpers"
)

// forbidden обрабатывает попытку получить доступ туда, куда нельзя
func (h *DecoratedHandler) forbidden(w http.ResponseWriter, r *http.Request) {
	session, err := coockies.Store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	flashMessages := session.Flashes()
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	helpers.ExecuteHTML("index", "forbidden", w, flashMessages)
}
