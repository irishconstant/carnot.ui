package controllers

import (
	"github.com/irishconstant/ui/helpers"
	"net/http"
)

func (h *DecoratedHandler) equipmentDelete(w http.ResponseWriter, r *http.Request) { //

	helpers.ExecuteHTML("equipment", "delete", w, nil)

}
