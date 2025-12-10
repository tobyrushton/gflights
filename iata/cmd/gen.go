package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"time"
)

type airport struct {
	Iata string
	Tz   string
	City string
}

func getAirports(url string) (map[string]airport, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	airports := map[string]airport{}
	err = json.NewDecoder(resp.Body).Decode(&airports)
	if err != nil {
		return nil, err
	}
	return airports, err
}

func main() {
	url := "https://raw.githubusercontent.com/mwgg/Airports/refs/heads/master/airports.json"
	airports, err := getAirports(url)
	if err != nil {
		log.Fatal(err)
	}

	iataFileContent := fmt.Sprintf(
		`// Package iata contains IATA airport codes, which are supported by the Google Flights API, along with time zones.
// This package was generated using an airport list (which can be found at this address: [airports.json])
// and the Google Flights API.
//
// Command: go run ./iata/cmd/gen.go
//
// Generation date: %s
//
// [airports.json]: %s
package iata

type Location struct{ City, Tz string }

// IATATimeZone turns IATA airport codes into the time zone where the airport is located.
// If IATATimeZone can't find an IATA airport code, then it returns "Not supported IATA Code".
func IATATimeZone(iata string) Location {
	switch iata {
`, time.Now().Format(time.DateOnly), url)

	lines := []string{}
	checked := map[string]struct{}{}

	caseTmpl := `	case "%s":
		return Location{"%s", "%s"}
`

	for _, airport := range airports {
		if airport.Iata == "" || airport.Iata == "0" {
			continue
		}
		if _, exists := checked[airport.Iata]; exists {
			continue
		}
		checked[airport.Iata] = struct{}{}
		lines = append(lines, fmt.Sprintf(caseTmpl, airport.Iata, airport.City, airport.Tz))
	}

	sort.Strings(lines)

	for _, line := range lines {
		iataFileContent += line
	}

	iataFileContent += `	}
	return Location{"Not supported IATA Code", "Not supported IATA Code"}
}
`

	iataFile, err := os.Create("./iata/iata.go")
	if err != nil {
		log.Fatal(err)
	}
	defer iataFile.Close()

	_, err = iataFile.WriteString(iataFileContent)
	if err != nil {
		log.Fatal(err)
	}
}
