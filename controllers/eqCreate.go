package controllers

import (
	"net/http"

	"github.com/irishconstant/ui/helpers"
)

func (h *DecoratedHandler) equipmentCreate(w http.ResponseWriter, r *http.Request) { //

	helpers.ExecuteHTML("equipment", "create", w, nil)

}
