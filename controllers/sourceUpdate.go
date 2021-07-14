package controllers

import (
	"net/http"
	"strconv"

	"github.com/irishconstant/ui/coockies"
	"github.com/irishconstant/ui/helpers"
)

func (h *DecoratedHandler) sourceUpdate(w http.ResponseWriter, r *http.Request) { //

	// Получаем текущий теплоисточник из параметров
	keySource, err := strconv.Atoi(r.URL.Query().Get("key"))
	helpers.Check(err)
	// Получаем текущий период из параметров
	/*
		var calcPeriod *ref.CalcPeriod
		period := r.URL.Query().Get("period")
		if period == "" {
			calcPeriod, err = h.connection.GetCurrentPeriod()
		} else {
			calcPeriodID, err := strconv.Atoi(period)
			calcPeriod, err = h.connection.GetCalcPeriod(calcPeriodID)
			if err != nil {
				fmt.Println("Передано ошибочное значение расчётного периода")
			}
		}
	*/

	Source, err := h.connection.GetSource(keySource, nil)
	helpers.Check(err)
	session, err := coockies.Store.Get(r, "cookie-name")
	helpers.Check(err)
	user := coockies.GetUser(session)

	if r.Method == http.MethodGet {
		helpers.Check(err)
		currentInformation := sessionInformation{User: *user, Attribute: Source}
		helpers.ExecuteHTML("source", "update", w, currentInformation)
	}
	/*
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
	*/
}
