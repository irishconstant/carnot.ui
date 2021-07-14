package controllers

import (
	"github.com/irishconstant/ui/helpers"
	"net/http"
)

func (h *DecoratedHandler) sourceCreate(w http.ResponseWriter, r *http.Request) { //

	helpers.ExecuteHTML("source", "edit", w, nil)

}
