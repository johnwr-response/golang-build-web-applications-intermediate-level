package main

import "net/http"

func (app *application) VirtualTerminal(_ http.ResponseWriter, _ *http.Request) {
	app.infoLog.Println("Hit the handler")
}
