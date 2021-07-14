package controllers

import (
	"net/http"

	"github.com/irishconstant/core/auth"

	"github.com/irishconstant/ui/coockies"
	"github.com/irishconstant/ui/helpers"
)

// login обрабатывает попытку залогиниться
func (h *DecoratedHandler) login(w http.ResponseWriter, r *http.Request) {
	session, err := coockies.Store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {
		login := r.FormValue("login")
		password := r.FormValue("password")
		if !h.connection.CheckPassword(login, password) {
			if password == "" {
				session.AddFlash("Необходимо ввести пароль")
			}
			session.AddFlash("Неправильное имя пользователя или пароль")
			err = session.Save(r, w)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/forbidden", http.StatusFound)
			return
		}

		user := &auth.User{
			Key:           login,
			Authenticated: true,
		}
		h.connection.GetUserRoles(user)
		err = h.connection.GetUserAttributes(user)
		helpers.Check(err)
		session.Values["SystemUser"] = user

		err = session.Save(r, w)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/person", http.StatusFound)
	}
}
