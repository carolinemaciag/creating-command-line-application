package converter

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

//Each house respresenting one record from the csv
type House struct {
	Value    int     `json:"value"`
	Income   float64 `json:"income"`
	Age      int     `json:"age"`
	Rooms    int     `json:"rooms"`
	Bedrooms int     `json:"bedrooms"`
	Pop      int     `json:"pop"`
	HH       int     `json:"hh"`
}

//Begin conversion to json lines
func ConvertCSVtoJSONLines(inputPath, outputPath string) error {
	records, err := readCSV(inputPath)
	if err != nil {
		return err
	}

	houses, err := parseHouses(records)
	if err != nil {
		return err
	}

	return writeJSONLines(outputPath, houses)
}


func readCSV(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	return reader.ReadAll()
}

//Converts csv into slices of house struct, skip the header row
func parseHouses(records [][]string) ([]House, error) {
	if len(records) < 2 {
		return nil, errors.New("no data rows found in CSV")
	}

	var houses []House

	for i, row := range records[1:] { // skip header
		if len(row) != 7 {
			log.Printf("Skipping row %d: wrong number of fields", i+2)
			continue
		}

		house, err := rowToHouse(row)
		if err != nil {
			return nil, fmt.Errorf("row %d: %v", i+2, err)
		}

		houses = append(houses, house)
	}

	return houses, nil
}

//Conversion of string data and if statements for potential errors
func rowToHouse(row []string) (House, error) {
	if len(row) != 7 {
		return House{}, fmt.Errorf("invalid number of fields: expected 7, got %d", len(row))
	}
	valueFloat, err := strconv.ParseFloat(row[0], 64)
	if err != nil {
		return House{}, fmt.Errorf("invalid value: %v", err)
	}
	income, err := strconv.ParseFloat(row[1], 64)
	if err != nil {
		return House{}, fmt.Errorf("invalid income: %v", err)
	}
	ageFloat, err := strconv.ParseFloat(row[2], 64)
	if err != nil {
		return House{}, fmt.Errorf("invalid age: %v", err)
	}
	roomsFloat, err := strconv.ParseFloat(row[3], 64)
	if err != nil {
		return House{}, fmt.Errorf("invalid rooms: %v", err)
	}
	bedroomsFloat, err := strconv.ParseFloat(row[4], 64)
	if err != nil {
		return House{}, fmt.Errorf("invalid bedrooms: %v", err)
	}
	popFloat, err := strconv.ParseFloat(row[5], 64)
	if err != nil {
		return House{}, fmt.Errorf("invalid pop: %v", err)
	}
	hhFloat, err := strconv.ParseFloat(row[6], 64)
	if err != nil {
		return House{}, fmt.Errorf("invalid hh: %v", err)
	}

	return House{
		Value:    int(valueFloat),
		Income:   income,
		Age:      int(ageFloat),
		Rooms:    int(roomsFloat),
		Bedrooms: int(bedroomsFloat),
		Pop:      int(popFloat),
		HH:       int(hhFloat),
	}, nil
}

//func to write each house into a json line and output
func writeJSONLines(path string, houses []House) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer file.Close()

	for _, house := range houses {
		line, err := json.Marshal(house)
		if err != nil {
			return fmt.Errorf("failed to marshal house: %v", err)
		}

		_, err = file.Write(append(line, '\n'))
		if err != nil {
			return fmt.Errorf("failed to write to output file: %v", err)
		}
	}

	return nil
}
