package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/test/{input}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		input := vars["input"]
		fmt.Fprintf(w, "Hello, %s", input)
	})

	fmt.Println("Server running or port 80!")
	http.ListenAndServe(":80", router)
}
