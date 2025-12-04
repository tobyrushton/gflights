package flights

import (
	"context"
	"testing"
	"time"

	"golang.org/x/text/currency"
	"golang.org/x/text/language"
)

func TestSerialiseURL1(t *testing.T) {
	session, err := New()
	if err != nil {
		t.Fatal(err)
	}

	expectedURL := "https://www.google.com/travel/flights/search?tfs=Gj4SCjIwMjYtMDMtMDMoAmoOCAISCi9tLzAzMHFiM3RqBwgBEgNTRk9yDAgCEggvbS8wNGpwbHIHCAESA0NERxo-EgoyMDI2LTAzLTExKAJqDAgCEggvbS8wNGpwbGoHCAESA0NER3IOCAISCi9tLzAzMHFiM3RyBwgBEgNTRk9CAQFIAZgBAQ&curr=GBP&hl=en"

	departuredate, _ := time.Parse("2006-01-02", "2026-03-03")
	returnDate, _ := time.Parse("2006-01-02", "2026-03-11")

	url, err := session.SerialiseURL(
		context.Background(),
		Args{
			departuredate,
			returnDate,
			[]string{"Los Angeles"},
			[]string{"SFO"},
			[]string{"London"},
			[]string{"CDG"},
			Options{
				Travelers: Travelers{Adults: 1},
				Currency:  currency.GBP,
				Stops:     OneStop,
				Class:     Economy,
				TripType:  RoundTrip,
				Lang:      language.English,
			},
		},
	)

	if err != nil {
		t.Fatalf("error during serialization: %s", err.Error())
	}

	if expectedURL != url {
		t.Fatalf("wrong serialised url, expected: %v serialised: %v", expectedURL, url)
	}
}

func TestSerialiseURL2(t *testing.T) {
	session, err := New()
	if err != nil {
		t.Fatal(err)
	}

	expectedURL := "https://www.google.com/travel/flights/search?tfs=GjMSCjIwMjYtMDYtMTAoA2oMCAISCC9tLzA0anBsagcIARIDU0ZPcgwIAhIIL20vMGYydjAaMxIKMjAyNi0wNi0yMCgDagwIAhIIL20vMGYydjByDAgCEggvbS8wNGpwbHIHCAESA1NGT0IEAQECA0gBmAEB&curr=GBP&hl=en"

	date, _ := time.Parse("2006-01-02", "2026-06-10")
	returnDate, _ := time.Parse("2006-01-02", "2026-06-20")

	url, err := session.SerialiseURL(
		context.Background(),
		Args{
			date,
			returnDate,
			[]string{"London"},
			[]string{"SFO"},
			[]string{"Miami"},
			[]string{},
			Options{
				Travelers: Travelers{Adults: 2, Children: 1, InfantsOnLap: 1},
				Currency:  currency.GBP,
				Stops:     TwoPlusStops,
				Class:     Economy,
				TripType:  RoundTrip,
				Lang:      language.English,
			},
		},
	)

	if err != nil {
		t.Fatalf("error during serialization: %s", err.Error())
	}

	if expectedURL != url {
		t.Fatalf("wrong serialised url, expected: %v serialised: %v", expectedURL, url)
	}
}
