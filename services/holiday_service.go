package services

import (
	"encoding/json"
	"net/http"
	"rest-service/models"
	"rest-service/utils"
)

const url = "https://api.victorsanmartin.com/feriados/en.json"

func FetchAndFilterHolidays(filterType, startDate, endDate string) ([]models.Holiday, error) {
	resp, err := http.Get(url)
	if err != nil {
		utils.Log.Errorf("Error making GET request: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	var holidayResponse models.HolidayResponse
	if err := json.NewDecoder(resp.Body).Decode(&holidayResponse); err != nil {
		utils.Log.Errorf("Error decoding JSON response: %v", err)
		return nil, err
	}

	holidays := holidayResponse.Data
	filteredHolidays := filterHolidays(holidays, filterType, startDate, endDate)
	return filteredHolidays, nil
}

func filterHolidays(holidays []models.Holiday, filterType, startDate, endDate string) []models.Holiday {
	var filtered []models.Holiday
	for _, holiday := range holidays {
		if (filterType == "" || holiday.Type == filterType) &&
			(startDate == "" || holiday.Date >= startDate) &&
			(endDate == "" || holiday.Date <= endDate) {
			filtered = append(filtered, holiday)
		}
	}
	return filtered
}
