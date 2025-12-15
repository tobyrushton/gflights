package gflights_test

import (
	"time"

	"github.com/tobyrushton/gflights"
)

func lhrToJfkFlight(depTime, arrTime time.Time) gflights.Flight {
	return gflights.Flight{
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
		FlightCode:     gflights.FlightCode{AirlineCode: "BA", FlightNumber: "177"},
		AirlineName:    "British Airways",
		Legroom:        "31 inches",
	}
}

func jfkToLhrFlight(depTime, arrTime time.Time) gflights.Flight {
	return gflights.Flight{
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
		FlightCode:     gflights.FlightCode{AirlineCode: "BA", FlightNumber: "180"},
		AirlineName:    "British Airways",
		Legroom:        "31 inches",
	}
}
