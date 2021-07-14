package helpers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// executeHTML инкапсулирует работу с шаблонами и генерацию html
func ExecuteHTML(folder string, page string, w http.ResponseWriter, param interface{}) {
	absPath, _ := filepath.Abs(fmt.Sprintf("../ui/views/%s/%s.html", folder, page))
	html, err := template.ParseFiles(absPath)
	Check(err)
	err = html.Execute(w, param)
	Check(err)
}

func MakeURLWithAttributes(origin string, params map[string]string) string {

	var paramPart string

	for key, value := range params {
		if value != "" {
			paramPart = paramPart + key + "=" + value + "&"
		}
	}
	result := "/" + origin + "?" + paramPart
	return result
}

func Check(err error) {
	if err != nil {
		fmt.Println("Ошибочка", err)
		log.Fatal(err)
	}
}
