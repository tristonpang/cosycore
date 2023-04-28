package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"cosycore/documents"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/test/{input}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		input := vars["input"]
		fmt.Fprintf(w, "Hello, %s", input)
	})

	router.HandleFunc("/documents", func(w http.ResponseWriter, r *http.Request) {
		docs := documents.GetAllDocuments()

		encoder := json.NewEncoder(w)

		w.Header().Set("Content-Type", "application/json")

		encoder.Encode(docs)
		// fmt.Fprint(w, docs)

	}).Methods("GET")

	router.HandleFunc("/documents", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		decoder := json.NewDecoder(r.Body)
		var document documents.Document

		err := decoder.Decode(&document)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}

		documents.AddDocument(document)
	}).Methods("POST")

	fmt.Println("Server running or port 80!")
	http.ListenAndServe(":80", router)
}
