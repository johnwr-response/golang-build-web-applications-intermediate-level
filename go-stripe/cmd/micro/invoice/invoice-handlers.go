package main

import (
	"fmt"
	"net/http"
)

func (app *application) CreateAndSendInvoice(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintf(w, "hello, %s", "world")
}
