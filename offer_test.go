package gflights

import (
	"context"
	"testing"
	"time"

	"golang.org/x/text/currency"
)

func lhrToJfkFlight(depTime, arrTime time.Time) Flight {
	return Flight{
		DepAirportCode: "LHR",
		DepAirportName: "London Heathrow",
		DepCity:        "London",
		ArrAirportCode: "JFK",
		ArrAirportName: "John F. Kennedy International",
		ArrCity:        "New York",
		DepTime:        depTime,
		ArrTime:        arrTime,
		Duration:       8 * time.Hour,
		Airplane:       "Boeing 777",
		FlightCode:     FlightCode{AirlineCode: "BA", FlightNumber: "177"},
		AirlineName:    "British Airways",
		Legroom:        "31 inches",
	}
}

func jfkToLhrFlight(depTime, arrTime time.Time) Flight {
	return Flight{
		DepAirportCode: "JFK",
		DepAirportName: "John F. Kennedy International",
		DepCity:        "New York",
		ArrAirportCode: "LHR",
		ArrAirportName: "London Heathrow",
		ArrCity:        "London",
		DepTime:        depTime,
		ArrTime:        arrTime,
		Duration:       7*time.Hour + 5*time.Minute,
		Airplane:       "Boeing 777",
		FlightCode:     FlightCode{AirlineCode: "BA", FlightNumber: "180"},
		AirlineName:    "British Airways",
		Legroom:        "31 inches",
	}
}

func TestGetReturnFlights(t *testing.T) {
	s, err := New()
	if err != nil {
		t.Fatal(err)
	}

	ukTZ, err := time.LoadLocation("Europe/London")
	if err != nil {
		t.Fatalf("could not load UK timezone: %v", err)
	}
	nyTZ, err := time.LoadLocation("America/New_York")
	if err != nil {
		t.Fatalf("could not load NY timezone: %v", err)
	}
	baseDate := time.Now().AddDate(0, 6, 0)
	depTime := time.Date(baseDate.Year(), baseDate.Month(), baseDate.Day(), 8, 20, 0, 0, ukTZ)
	arrTime := time.Date(baseDate.Year(), baseDate.Month(), baseDate.Day(), 11, 20, 0, 0, nyTZ)

	offer := RoundTripOffer{
		OneWayOffer: OneWayOffer{
			SimpleOffer: SimpleOffer{
				StartDate:  time.Now().AddDate(0, 6, 0),
				ReturnDate: time.Now().AddDate(0, 6, 7),
				Price:      496,
			},
			Flight:         []Flight{lhrToJfkFlight(depTime, arrTime)},
			SrcAirportCode: "LHR",
			DstAirportCode: "JFK",
			Duration:       8 * time.Hour,
		},
		s: s,
		args: Args{
			Options: Options{
				Currency: currency.GBP,
				Class:    Economy,
				TripType: RoundTrip,
				Travelers: Travelers{
					Adults: 1,
				},
			},
			DepartureDate: time.Now().AddDate(0, 6, 0),
			ReturnDate:    time.Now().AddDate(0, 6, 7),
			SrcCities:     []string{"London"},
			DstAirports:   []string{"JFK"},
		},
	}

	flights, err := offer.GetReturnFlights(context.Background())
	if err != nil {
		t.Fatalf("error getting return flights: %s", err.Error())
	}
	if len(flights) == 0 {
		t.Fatal("expected at least one return flight, got none")
	}
}
