package documents

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var DOCUMENTS_PREFIX = "/documents"

func SetupDocumentsControllers(router *mux.Router) {
	router.HandleFunc(DOCUMENTS_PREFIX, func(w http.ResponseWriter, r *http.Request) {
		docs := GetAllDocuments()

		encoder := json.NewEncoder(w)

		w.Header().Set("Content-Type", "application/json")

		encoder.Encode(docs)
		// fmt.Fprint(w, docs)

	}).Methods("GET")

	router.HandleFunc(DOCUMENTS_PREFIX, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		decoder := json.NewDecoder(r.Body)
		var document Document

		err := decoder.Decode(&document)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}

		err = AddDocument(document)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}).Methods("POST")

	router.HandleFunc(DOCUMENTS_PREFIX, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		decoder := json.NewDecoder(r.Body)
		var document Document

		err := decoder.Decode(&document)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}

		err = UpdateDocument(document)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}).Methods("PUT")

	router.HandleFunc(DOCUMENTS_PREFIX, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		decoder := json.NewDecoder(r.Body)
		var document Document

		err := decoder.Decode(&document)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}

		err = DeleteDocument(document)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}).Methods("DELETE")

}
