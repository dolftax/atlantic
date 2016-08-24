package main

import (
	"os"
)

func isDir(path string) bool {
	fdir, err := os.Open(path)
	if err != nil {
		return false
	}
	defer fdir.Close()

	finfo, err := fdir.Stat()

	if err != nil {
		return false
	}

	switch mode := finfo.Mode(); {

	case mode.IsDir():
		return true
	case mode.IsRegular():
		return false
	}
	return false
}

func configExists(path string) bool {
	_, err := os.Stat(path + "/.codeclimate.yml")
	if err != nil {
		return false
	}
	return true
}

func changePwd(path string) bool {
	err := os.Chdir(path)
	if err != nil {
		return false
	}
	return true
}
