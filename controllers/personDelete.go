package controllers

import (
	"net/http"
	"strconv"

	"github.com/irishconstant/core/contract"
	"github.com/irishconstant/ui/coockies"
	"github.com/irishconstant/ui/helpers"
)

func (h *DecoratedHandler) personDelete(w http.ResponseWriter, r *http.Request) {
	keyPerson, err := strconv.Atoi(r.URL.Query().Get("key"))
	helpers.Check(err)
	Person, err := h.connection.GetPerson(keyPerson)
	helpers.Check(err)
	session, err := coockies.Store.Get(r, "cookie-name")
	helpers.Check(err)
	user := coockies.GetUser(session)
	//err = h.connection.GetUserAttributes(&user)
	helpers.Check(err)

	if r.Method == http.MethodGet {
		Person.PossibleUsers, err = h.connection.GetAllUsers()
		helpers.Check(err)
		currentInformation := sessionInformation{User: *user, Attribute: Person}
		helpers.ExecuteHTML("Person", "update", w, currentInformation)
	}

	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		familyName := r.FormValue("familyname")
		patronymicName := r.FormValue("patronymicname")
		userLogin := r.FormValue("user")

		user, err := h.connection.GetUser(userLogin)
		helpers.Check(err)
		newPerson := contract.Person{
			Key:            Person.Key,
			Name:           name,
			FamilyName:     familyName,
			PatronymicName: patronymicName,
			User:           *user,
		}

		err = h.connection.UpdatePerson(&newPerson)

		if err != nil {
			helpers.ExecuteHTML("person", "update", w, nil)
		}
		http.Redirect(w, r, "/person", http.StatusFound)
	}

}
