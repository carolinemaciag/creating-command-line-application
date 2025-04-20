package converter

import (
	"testing"
)

//Testing each important function
func TestRowToHouse_ValidRow(t *testing.T) {
	row := []string{"452600", "8.3252", "41", "880", "129", "322", "126"}
	_, err := rowToHouse(row)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestRowToHouse_InvalidValue(t *testing.T) {
	row := []string{"notANumber", "8.3252", "41", "880", "129", "322", "126"}
	_, err := rowToHouse(row)
	if err == nil {
		t.Error("Expected error for invalid value, got nil")
	}
}

func TestRowToHouse_InvalidIncome(t *testing.T) {
	row := []string{"452600", "notAFloat", "41", "880", "129", "322", "126"}
	_, err := rowToHouse(row)
	if err == nil {
		t.Error("Expected error for invalid income, got nil")
	}
}

func TestRowToHouse_InvalidAge(t *testing.T) {
	row := []string{"452600", "8.3252", "notANumber", "880", "129", "322", "126"}
	_, err := rowToHouse(row)
	if err == nil {
		t.Error("Expected error for invalid age, got nil")
	}
}

func TestRowToHouse_InvalidRooms(t *testing.T) {
	row := []string{"452600", "8.3252", "41", "badRooms", "129", "322", "126"}
	_, err := rowToHouse(row)
	if err == nil {
		t.Error("Expected error for invalid rooms, got nil")
	}
}

func TestRowToHouse_InvalidBedrooms(t *testing.T) {
	row := []string{"452600", "8.3252", "41", "880", "bad", "322", "126"}
	_, err := rowToHouse(row)
	if err == nil {
		t.Error("Expected error for invalid bedrooms, got nil")
	}
}

func TestRowToHouse_InvalidPopulation(t *testing.T) {
	row := []string{"452600", "8.3252", "41", "880", "129", "bad", "126"}
	_, err := rowToHouse(row)
	if err == nil {
		t.Error("Expected error for invalid population, got nil")
	}
}

func TestRowToHouse_InvalidHouseholds(t *testing.T) {
	row := []string{"452600", "8.3252", "41", "880", "129", "322", "bad"}
	_, err := rowToHouse(row)
	if err == nil {
		t.Error("Expected error for invalid households, got nil")
	}
}

func TestRowToHouse_TooFewFields(t *testing.T) {
	row := []string{"452600", "8.3252", "41"} // only 3 fields
	_, err := rowToHouse(row)
	if err == nil {
		t.Error("Expected error for too few fields, got nil")
	}
}

