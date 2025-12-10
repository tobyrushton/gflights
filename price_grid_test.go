package gflights

import (
	"context"
	"testing"
	"time"

	"golang.org/x/text/currency"
)

func TestPriceGrid(t *testing.T) {
	session, err := New()
	if err != nil {
		t.Fatalf("Failed to create session: %v", err)
	}

	offers, err := session.GetPriceGrid(
		context.Background(),
		PriceGridArgs{
			StartDepartureRange: time.Now().AddDate(0, 1, 0),
			EndDepartureRange:   time.Now().AddDate(0, 1, 7),
			StartReturnRange:    time.Now().AddDate(0, 1, 14),
			EndReturnRange:      time.Now().AddDate(0, 1, 21),
			SrcCities:           []string{"London"},
			DstCities:           []string{"New York"},
			Options: Options{
				Travelers: Travelers{Adults: 1},
				Currency:  currency.GBP,
				TripType:  RoundTrip,
				Class:     Economy,
			},
		},
	)
	if err != nil {
		t.Fatalf("GetPriceGrid failed: %v", err)
	}

	if len(offers) == 0 {
		t.Fatalf("Expected at least one offer, got none")
	}

	if len(offers) != 64 {
		t.Fatalf("Expected 49 offers, got %d", len(offers))
	}
}
