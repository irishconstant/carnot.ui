module github.com/irishconstant/ui

go 1.16

require (
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/securecookie v1.1.1
	github.com/gorilla/sessions v1.2.1
	github.com/irishconstant/core v0.0.0-20210714142336-71f1fc290d1d
	github.com/irishconstant/db v0.0.0-20210714145351-6fe1e459661e
)

replace github.com/irishconstant/db => ../db
