package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Cerchie/go-api/pkg/mocks"
)

func GetAllHaikus(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(mocks.Haikus)
}
