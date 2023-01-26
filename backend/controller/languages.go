package controller

import (
	"encoding/json"
	"github.com/KrisjanisP/deikstra/service/models"
	"net/http"
)

// c.router.HandleFunc("/languages/list", c.listLanguages).Methods("GET", "OPTIONS")
func (c *Controller) listLanguages(w http.ResponseWriter, r *http.Request) {
	// CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	if r.Method == http.MethodOptions {
		return
	}

	var languages []models.Language
	err := c.database.Model(&models.Language{}).Find(&languages).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := json.Marshal(languages)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = w.Write(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
