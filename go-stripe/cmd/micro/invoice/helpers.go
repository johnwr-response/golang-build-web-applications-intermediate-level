package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
)

// writeJSON writes arbitrary data out as json
func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for k, v := range headers[0] {
			w.Header()[k] = v
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(out)
	return nil
}

// readJSON reads json from request body into data. We only accept a single json value in the body
func (app *application) readJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	// Sanity checks
	// don't handle any json larger than 1MB
	maxBytes := 1024 * 1024
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return err
	}
	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only have a single JSON value")
	}
	return nil
}

// badRequest sends a json response with status http.StatusBadRequest, describing the error
func (app *application) badRequest(w http.ResponseWriter, _ *http.Request, err error) error {
	var payload struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}
	payload.Error = true
	payload.Message = err.Error()

	out, err := json.MarshalIndent(payload, "", "\t")
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	_, _ = w.Write(out)
	return nil
}

func (app *application) CreateDirIfNotExist(path string) error {
	const mode = 0755
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, mode)
		if err != nil {
			app.errorLog.Println(err)
			return err
		}
	}
	return nil
}
