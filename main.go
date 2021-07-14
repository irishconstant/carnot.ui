package main

import (
	"github.com/irishconstant/db"
	"github.com/irishconstant/ui/controllers"
)

func main() {
	dbc := db.InitDBConnection()
	defer dbc.CloseConnect()
	controllers.Router(dbc)
}
