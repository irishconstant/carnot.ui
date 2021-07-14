package controllers

import (
	"github.com/irishconstant/ui/helpers"
	"net/http"
)

func (h *DecoratedHandler) equipment(w http.ResponseWriter, r *http.Request) { //

	helpers.ExecuteHTML("equipment", "list", w, nil)

}
