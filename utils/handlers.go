package utils

import (
	"net/http"
)

func StartHandle(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("../static"))

}
func FileHandler(w http.ResponseWriter, r *http.Request) {

}
