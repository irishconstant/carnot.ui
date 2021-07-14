package controllers

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/irishconstant/ui/coockies"
	"github.com/irishconstant/ui/helpers"
)

// Person обработчик доступен только авторизованным пользователям, прошедшим аутентификацию. Контроллируется middleware Auth
func (h *DecoratedHandler) person(w http.ResponseWriter, r *http.Request) { //

	if r.Method == http.MethodPost {
		// Получаем данные фильтров из формы и формируем параметры для вызова
		params := make(map[string]string)
		params["name"] = r.FormValue("name")
		params["familyname"] = r.FormValue("familyname")
		params["patrname"] = r.FormValue("patrname")
		params["sex"] = r.FormValue("sex")
		filteredAddress := helpers.MakeURLWithAttributes("person", params)
		// Переходим на этот урл
		http.Redirect(w, r, filteredAddress, http.StatusFound)
	}

	// Получаем текущую страницу из параметров
	key := r.URL.Query().Get("page")
	var page int
	if key != "" {
		page, _ = strconv.Atoi(key)
	} else {
		page = 1
	}

	// Получаем параметры фильтрации из урла
	name := r.URL.Query().Get("name")
	familyName := r.URL.Query().Get("familyname")
	patrName := r.URL.Query().Get("patrname")
	sex := r.URL.Query().Get("sex")

	// Работа с пользователями
	session, err := coockies.Store.Get(r, "cookie-name")
	helpers.Check(err)
	user := coockies.GetUser(session)
	helpers.Check(err)
	//err = h.connection.GetUserAttributes(&user)
	//helpers.Check(err)

	// Работа с ФЛ
	quantity, err := h.connection.GetPersonQuantityFiltered(*user, name, familyName, patrName, sex)
	helpers.Check(err)
	PersonBook := PersonBook{PersonCount: quantity}

	// Если необходима пагинация
	if PersonBook.PersonCount > h.pageSize {
		PersonsPerPage, err := h.connection.GetPersonsFiltered(*user, 1, page, h.pageSize, name, familyName, patrName, sex)
		helpers.Check(err)
		for _, value := range PersonsPerPage {
			PersonBook.Persons = append(PersonBook.Persons, *value)
		}
		PersonBook.CurrentPage = page

		// Создаем страницы для показа (1, одна слева от текущей, одна справа от текущей, последняя)
		// Инициализируем фильтры для кнопок пагинации, которые к нам ранее пришли в POST запросе
		if name != "" {
			name = "&name=" + name
		}
		if familyName != "" {
			familyName = "&familyname=" + familyName
		}
		if patrName != "" {
			patrName = "&patrname=" + patrName
		}
		if sex != "" {
			sex = "&sex=" + sex
		}
		PersonBook.Pages = MakePages(1, int(math.Ceil(float64(PersonBook.PersonCount)/float64(h.pageSize))), page)
		for key := range PersonBook.Pages {
			PersonBook.Pages[key].URL = fmt.Sprintf("/person?%s%s%s%s", name, familyName, patrName, sex)
		}
		currentInformation := sessionInformation{User: *user, Attribute: PersonBook}
		helpers.ExecuteHTML("person", "list", w, currentInformation)

	} else {
		Persons, _ := h.connection.GetPersonsFiltered(*user, 0, page, h.pageSize, name, familyName, patrName, sex)
		for _, value := range Persons {
			PersonBook.Persons = append(PersonBook.Persons, *value)
		}

		currentInformation := sessionInformation{User: *user, Attribute: PersonBook}

		helpers.ExecuteHTML("person", "list", w, currentInformation)
	}
}
