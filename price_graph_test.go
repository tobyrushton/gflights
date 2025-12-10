package gflights

import (
	"context"
	"testing"
	"time"

	"golang.org/x/text/currency"
	"golang.org/x/text/language"
)

func TestPriceGraph(t *testing.T) {
	session, err := New()
	if err != nil {
		t.Fatalf("failed to create session: %v", err)
	}

	daysDiff1 := 60
	daysDiff2 := 90
	expectedOffersCount := daysDiff2 - daysDiff1 + 1

	offers, err := session.GetPriceGraph(
		context.Background(),
		PriceGraphArgs{
			RangeStartDate: time.Now().AddDate(0, 0, daysDiff1),
			RangeEndDate:   time.Now().AddDate(0, 0, daysDiff2),
			TripLength:     7,
			SrcCities:      []string{"New York"},
			DstCities:      []string{"London"},
			Options: Options{
				Travelers: Travelers{Adults: 1},
				Currency:  currency.GBP,
				Class:     Economy,
				Lang:      language.English,
				TripType:  RoundTrip,
			},
		},
	)
	if err != nil {
		t.Fatalf("GetPriceGraph failed: %v", err)
	}
	if len(offers) != expectedOffersCount {
		t.Fatalf("expected %d offers, got %d", expectedOffersCount, len(offers))
	}
}
