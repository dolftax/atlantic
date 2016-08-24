package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func initConfig(path) {
}

func checkConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	path := ps.ByName("path")
	if err := isDir(path); err == true {
		if err = configExists(path); err == true {
			_ = changePwd(path)
			go analyze(w, r, path)
		} else {
			_ = changePwd(path)
			go initConfig(w, r, path)
		}
	} else {
		errorHandler(w, r, path, 1001) // Path is a file
	}
}
