package gflights_test

import (
	"context"
	"testing"
	"time"

	"github.com/tobyrushton/gflights"
	"golang.org/x/text/currency"
)

func TestExplore(t *testing.T) {
	session, err := gflights.New()
	if err != nil {
		t.Fatal(err)
	}

	args := gflights.ExploreArgs{
		DepartureDate: time.Now().Add(time.Hour * 24),
		ReturnDate:    time.Now().Add(time.Hour * 24 * 7),
		SrcCities:     []string{"LONDON"},
		Options: gflights.Options{
			Travelers: gflights.Travelers{Adults: 1},
			Currency:  currency.GBP,
			Stops:     gflights.NonStop,
			Class:     gflights.Economy,
			TripType:  gflights.RoundTrip,
		},
	}

	offers, err := session.GetExplore(context.Background(), args)
	if err != nil {
		t.Fatal(err)
	}

	if len(offers) == 0 {
		t.Fatal("expected at least one offer")
	}
}
