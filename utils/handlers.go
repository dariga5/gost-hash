package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func StartHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello! Send file to address: POST /api/v1/sendfile")

}
func FileHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Wethod not allowed", http.StatusMethodNotAllowed)
		return
	}

	file, header, err := r.FormFile("file")

	log.Print(header)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

	dst, err := os.Create("./uplodes")

	if err != nil {
		return
	}

	io.Copy(dst, file)
	fmt.Fprintf(w, "Upload succesful")
}
