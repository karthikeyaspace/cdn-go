package handler

import (
	"fmt"
	"net/http"
	
	"github.com/gorilla/mux"
)

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Printf("View endpoint hit with ID: %s\n", id)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Viewing content with ID: %s", id)))
}
