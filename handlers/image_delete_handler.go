package handlers

import (
	"go_discord_file_server/services"
	"net/http"
)

func ImageDeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := services.DeleteImage(r.URL.Query().Get("msgId"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
