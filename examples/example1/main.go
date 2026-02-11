package main

import (
	"context"
	"fmt"
	"time"

	"github.com/tobyrushton/gflights"
	"golang.org/x/text/currency"
)

func main() {
	s, err := gflights.New()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()

	offers, _, err := s.GetOutboundOffers(ctx, gflights.Args{
		SrcCities:     []string{"London"},
		DstCities:     []string{"New York"},
		DepartureDate: time.Now().Add(24 * time.Hour),
		ReturnDate:    time.Now().Add(7 * 24 * time.Hour),
		Options: gflights.Options{
			Currency: currency.GBP,
			Stops:    gflights.NonStop,
			Travelers: gflights.Travelers{
				Adults: 1,
			},
			TripType: gflights.RoundTrip,
			Class:    gflights.Economy,
		},
	})
	if err != nil {
		panic(err)
	}

	returnOffers, err := offers[0].GetReturnFlights(ctx)
	if err != nil {
		panic(err)
	}

	trip, err := offers[0].SelectReturnFlight(returnOffers[0])
	if err != nil {
		panic(err)
	}

	url, err := s.SerialiseBookingURL(ctx, trip)
	if err != nil {
		panic(err)
	}

	fmt.Println(url)
}
