package controllers

import (
	"encoding/json"
	"net/http"
	// Uvozite ostale potrebne pakete
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	// Implementacija
	json.NewEncoder(w).Encode("Ovde ide odgovor")
}
