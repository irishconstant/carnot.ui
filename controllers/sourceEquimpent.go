package controllers

import (
	"github.com/irishconstant/ui/helpers"
	"net/http"
)

func (h *DecoratedHandler) sourceEquipment(w http.ResponseWriter, r *http.Request) { //

	helpers.ExecuteHTML("source", "scheme", w, nil)

}
