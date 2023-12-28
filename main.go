package main

import (
	routes "gost-hash/server/routes"
	"log"
	"net/http"
)

const port = ":7089"

func main() {

	router := routes.InitRouter()

	server := &http.Server{
		Addr:    port,
		Handler: router,
	}

	log.Println("Заруск сервера http://127.0.0.1" + port)

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
