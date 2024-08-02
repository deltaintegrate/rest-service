package handlers

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"rest-service/services"
	"rest-service/utils"

	"github.com/gorilla/mux"
)

func RegisterHolidayRoutes(router *mux.Router) {
	router.HandleFunc("/holidays", GetHolidays).Methods("GET")
}

func GetHolidays(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	filterType := queryParams.Get("type")
	startDate := queryParams.Get("start_date")
	endDate := queryParams.Get("end_date")
	acceptHeader := r.Header.Get("Accept")

	holidays, err := services.FetchAndFilterHolidays(filterType, startDate, endDate)
	if err != nil {
		utils.Log.Errorf("Error fetching holidays: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if acceptHeader == "application/xml" {
		w.Header().Set("Content-Type", "application/xml")
		if err := xml.NewEncoder(w).Encode(holidays); err != nil {
			utils.Log.Errorf("Error encoding response to XML: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	} else {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(holidays); err != nil {
			utils.Log.Errorf("Error encoding response to JSON: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}
