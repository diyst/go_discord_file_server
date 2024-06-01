package handlers

import (
	"encoding/json"
	"go_discord_file_server/services"
	"net/http"
)

type ImageGetResponse struct {
	Url string `json:"url"`
}

func ImageGetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.URL.Query()

	url, err := services.GetImage(r.URL.Query().Get("msgId"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := ImageGetResponse{Url: url}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
