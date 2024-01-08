package utils

import (
	"net/http"
)

func StartHandle(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "")

}
func FileHandler(w http.ResponseWriter, r *http.Request) {

}
