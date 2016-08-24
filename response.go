package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type ResponseObj struct {
	Err       int         `json:error`
	Timestamp time.Time   `json:timestamp`
	Path      string      `json:path`
	Result    interface{} `json:result`
}

func responseDispatcher(w http.ResponseWriter, r *http.Request, message ResponseObj, isErr bool) {
	response, err := json.Marshal(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if isErr == true {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Write(response)
}

func responseHandler(w http.ResponseWriter, r *http.Request, path string, result interface{}) {
	message := ResponseObj{0, time.Now(), path, result}
	responseDispatcher(w, r, message, false)
}

func errorHandler(w http.ResponseWriter, r *http.Request, path string, err int) {
	message := ResponseObj{err, time.Now(), path, ""}
	responseDispatcher(w, r, message, true)
}
