package main

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func bodyParser(requestData []byte, r *http.Request, w http.ResponseWriter, path string) (arguments []string) {
	var parsedObj []string
	err := json.Unmarshal(requestData, &requestObj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	return parsedObj
}

func analyser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	arguments := bodyParser(requestData, r, w, ps.ByName("path"))
	var cmd string

	for _, argument := range arguments {
		cmd += " " + argument
	}
}
