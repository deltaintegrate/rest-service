package tests

import (
	"rest-service/services"
	"testing"
)

func TestFetchAndFilterHolidays(t *testing.T) {
	holidays, err := services.FetchAndFilterHolidays("Religious", "2024-01-01", "2024-12-31")
	if err != nil {
		t.Fatalf("Failed to fetch holidays: %v", err)
	}

	if len(holidays) == 0 {
		t.Error("Expected non-zero holidays")
	}
}
