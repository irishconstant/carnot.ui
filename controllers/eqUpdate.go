package controllers

import (
	"net/http"

	"github.com/irishconstant/ui/helpers"
)

func (h *DecoratedHandler) equipmentUpdate(w http.ResponseWriter, r *http.Request) { //

	helpers.ExecuteHTML("equipment", "update", w, nil)
}
