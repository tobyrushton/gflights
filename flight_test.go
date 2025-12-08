package gflights

import (
	"context"
	"testing"
	"time"

	"golang.org/x/text/currency"
	"golang.org/x/text/language"
)

func TestGetOneWayOffers(t *testing.T) {
	session, err := New()
	if err != nil {
		t.Fatal(err)
	}

	offers, _, err := session.GetOffers(
		context.Background(),
		Args{
			DepartureDate: time.Now().AddDate(0, 6, 0),
			SrcCities:     []string{"New York"},
			DstAirports:   []string{"LHR"},
			Options: Options{
				Travelers: Travelers{Adults: 1},
				Currency:  currency.GBP,
				Stops:     NonStop,
				Class:     Economy,
				TripType:  OneWay,
				Lang:      language.English,
			},
		},
	)

	if err != nil {
		t.Fatalf("error getting flight offers: %s", err.Error())
	}

	if len(offers) == 0 {
		t.Fatal("expected at least one flight offer, got none")
	}
}

func TestGetRoundTripOffers(t *testing.T) {
	session, err := New()
	if err != nil {
		t.Fatal(err)
	}

	offers, _, err := session.GetRoundTripOffers(
		context.Background(),
		Args{
			DepartureDate: time.Now().AddDate(0, 6, 0),
			ReturnDate:    time.Now().AddDate(0, 6, 7),
			SrcCities:     []string{"New York"},
			DstAirports:   []string{"LHR"},
			Options: Options{
				Travelers: Travelers{Adults: 1},
				Currency:  currency.EUR,
				Stops:     AnyStops,
				Class:     Business,
				TripType:  RoundTrip,
				Lang:      language.French,
			},
		},
	)

	if err != nil {
		t.Fatalf("error getting flight offers: %s", err.Error())
	}

	if len(offers) == 0 {
		t.Fatal("expected at least one flight offer, got none")
	}
}
