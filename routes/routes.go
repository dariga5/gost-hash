package routes

import (
	ustls "gost-hash/server/utils"
	"net/http"
)

func InitRouter() *http.ServeMux {

	router := http.NewServeMux()

	router.HandleFunc("/home", ustls.StartHandle)
	router.HandleFunc("/api/v1/sendfile", ustls.FileHandler)

	return router
}
