package server

import (
	"net/http"
	"os"
)

func CheckRequest(r *http.Request) {
	// Sanatization Function. Basic protection from basic attacks.

}

func FileConfirm(filepath string) bool {
	// Checks if the requested file exists

	_, err := os.Stat(filepath)
	return os.IsNotExist(err)

}
