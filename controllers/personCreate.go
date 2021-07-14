package controllers

import (
	"net/http"
	"strconv"

	"github.com/irishconstant/core/contract"
	"github.com/irishconstant/ui/coockies"
	"github.com/irishconstant/ui/helpers"
)

// PersonCreate обработчик доступен только авторизованным пользователям, прошедшим аутентификацию. Контроллируется middleware Auth
func (h *DecoratedHandler) personCreate(w http.ResponseWriter, r *http.Request) {
	// Работа с куками
	session, err := coockies.Store.Get(r, "cookie-name")
	helpers.Check(err)
	user := coockies.GetUser(session)
	//err = h.connection.GetUserAttributes(&user)
	helpers.Check(err)

	var userBook UserBook
	userBook.Users, err = h.connection.GetAllUsers()
	helpers.Check(err)

	if r.Method == http.MethodGet {
		currentInformation := sessionInformation{User: *user, Attribute: userBook}
		helpers.ExecuteHTML("person", "create", w, currentInformation)
	}

	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		familyName := r.FormValue("familyname")
		patronymicName := r.FormValue("patronymicname")
		sex, _ := strconv.ParseBool(r.FormValue("sex"))

		userID := r.FormValue("user")

		User, _ := h.connection.GetUser(userID)

		newPerson := contract.Person{
			Name:           name,
			FamilyName:     familyName,
			PatronymicName: patronymicName,
			Sex:            sex,
			User:           *User,
		}

		err = h.connection.CreatePerson(&newPerson)

		if err != nil {
			helpers.ExecuteHTML("person", "create", w, nil)
		}
		http.Redirect(w, r, "/person", http.StatusFound)

	}
}
