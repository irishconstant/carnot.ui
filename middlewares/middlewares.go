package middlewares

import (
	"net/http"
	"strings"

	"github.com/irishconstant/ui/coockies"
)

// authMiddleware выполняется для проверки аутентифицирован ли пользователь. TODO: сделать доступ к определенным разделам по ролям
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := coockies.Store.Get(r, "cookie-name")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		user := coockies.GetUser(session)

		if user.Authenticated == false {
			session.AddFlash("Доступ запрещён (пройдите авторизацию и аутентификацию)!")
			err = session.Save(r, w)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				http.Redirect(w, r, "/forbidden", http.StatusFound)
				return
			}
			http.Redirect(w, r, "/forbidden", http.StatusFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func CaselessMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.ToLower(r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
