package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Cerchie/go-api/pkg/mocks"
	"github.com/gorilla/mux"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
    // Read dynamic id parameter
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])

    // Iterate over all the mock books
    for _, haiku := range mocks.Haikus {
        if haiku.Id == id {
            // If ids are equal send book as a response
            w.WriteHeader(http.StatusOK)
            w.Header().Add("Content-Type", "application/json")
            json.NewEncoder(w).Encode(haiku)
            break
        }
    }
}