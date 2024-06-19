package main

import (
	"encoding/json"
	"net/http"
)

func manageHandler(w http.ResponseWriter, r *http.Request) {
	var req ManageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	handleManageRequest(req)
}
