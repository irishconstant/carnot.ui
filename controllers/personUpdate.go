package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/irishconstant/core/contract"
	"github.com/irishconstant/ui/coockies"
	"github.com/irishconstant/ui/helpers"
)

func (h *DecoratedHandler) personUpdate(w http.ResponseWriter, r *http.Request) {
	keyPerson, err := strconv.Atoi(r.URL.Query().Get("key"))
	helpers.Check(err)
	Person, err := h.connection.GetPerson(keyPerson)
	helpers.Check(err)
	session, err := coockies.Store.Get(r, "cookie-name")
	helpers.Check(err)
	user := coockies.GetUser(session)

	if r.Method == http.MethodGet {
		Person.PossibleUsers, err = h.connection.GetAllUsers()
		helpers.Check(err)
		currentInformation := sessionInformation{User: *user, Attribute: Person}
		helpers.ExecuteHTML("person", "update", w, currentInformation)
	}

	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		familyName := r.FormValue("familyname")
		patronymicName := r.FormValue("patronymicname")
		dateBirth := r.FormValue("datebirth")
		dateDeath := r.FormValue("datedeath")
		sex, _ := strconv.ParseBool(r.FormValue("sex"))
		userLogin := r.FormValue("user")

		dateBirthG, _ := time.Parse("2006-01-02", dateBirth)
		dateDeathG, _ := time.Parse("2006-01-02", dateDeath)

		newUser, err := h.connection.GetUser(userLogin)
		helpers.Check(err)
		newPerson := contract.Person{
			Key:            Person.Key,
			Name:           name,
			FamilyName:     familyName,
			PatronymicName: patronymicName,
			Sex:            sex,
			DateBirth:      dateBirthG,
			DateDeath:      dateDeathG,
			User:           *newUser,
		}

		err = h.connection.UpdatePerson(&newPerson)

		if err != nil {
			helpers.ExecuteHTML("person", "update", w, nil)
		}
		http.Redirect(w, r, "/person", http.StatusFound)
	}

}
