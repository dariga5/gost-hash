package utils

import (
	"net/http"
)

func StartHandle(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../static/index.html")

}
func FileHandler(w http.ResponseWriter, r *http.Request) {

}
