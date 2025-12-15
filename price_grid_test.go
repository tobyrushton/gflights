package gflights_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/tobyrushton/gflights"
	"github.com/tobyrushton/gflights/internal/fakes"
	"golang.org/x/text/currency"
	"golang.org/x/text/language"
)

func TestPriceGrid(t *testing.T) {
	session, err := gflights.New()
	if err != nil {
		t.Fatalf("Failed to create session: %v", err)
	}

	offers, err := session.GetPriceGrid(
		context.Background(),
		gflights.PriceGridArgs{
			StartDepartureRange: time.Now().AddDate(0, 1, 0),
			EndDepartureRange:   time.Now().AddDate(0, 1, 7),
			StartReturnRange:    time.Now().AddDate(0, 1, 14),
			EndReturnRange:      time.Now().AddDate(0, 1, 21),
			SrcCities:           []string{"London"},
			DstCities:           []string{"New York"},
			Options: gflights.Options{
				Travelers: gflights.Travelers{Adults: 1},
				Currency:  currency.GBP,
				TripType:  gflights.RoundTrip,
				Class:     gflights.Economy,
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

func TestPriceGridLarge(t *testing.T) {
	session, err := gflights.New()
	if err != nil {
		t.Fatalf("Failed to create session: %v", err)
	}

	offers, err := session.GetPriceGrid(
		context.Background(),
		gflights.PriceGridArgs{
			StartDepartureRange: time.Now().AddDate(0, 1, 0),
			EndDepartureRange:   time.Now().AddDate(0, 1, 19),
			StartReturnRange:    time.Now().AddDate(0, 1, 21),
			EndReturnRange:      time.Now().AddDate(0, 1, 30),
			SrcCities:           []string{"London"},
			DstCities:           []string{"New York"},
			Options: gflights.Options{
				Travelers: gflights.Travelers{Adults: 1},
				Currency:  currency.GBP,
				TripType:  gflights.RoundTrip,
				Class:     gflights.Economy,
			},
		},
	)
	if err != nil {
		t.Fatalf("GetPriceGrid failed: %v", err)
	}

	if len(offers) == 0 {
		t.Fatalf("Expected at least one offer, got none")
	}

	if len(offers) != 200 {
		t.Fatalf("Expected 200 offers, got %d", len(offers))
	}
}

func TestPriceGridCrossOver(t *testing.T) {
	session, err := gflights.New()
	if err != nil {
		t.Fatalf("Failed to create session: %v", err)
	}

	offers, err := session.GetPriceGrid(
		context.Background(),
		gflights.PriceGridArgs{
			StartDepartureRange: time.Now().AddDate(0, 1, 0),
			EndDepartureRange:   time.Now().AddDate(0, 1, 10),
			StartReturnRange:    time.Now().AddDate(0, 1, 5),
			EndReturnRange:      time.Now().AddDate(0, 1, 15),
			SrcCities:           []string{"London"},
			DstCities:           []string{"New York"},
			Options: gflights.Options{
				Travelers: gflights.Travelers{Adults: 1},
				Currency:  currency.GBP,
				TripType:  gflights.RoundTrip,
				Class:     gflights.Economy,
			},
		},
	)
	if err != nil {
		t.Fatalf("GetPriceGrid failed: %v", err)
	}

	if len(offers) == 0 {
		t.Fatalf("Expected at least one offer, got none")
	}

	for _, offer := range offers {
		if offer.ReturnDate.Before(offer.DepartureDate) {
			t.Fatalf("Found offer with return date before departure date: Departure %s, Return %s",
				offer.DepartureDate.Format("2006-01-02"),
				offer.ReturnDate.Format("2006-01-02"))
		}
	}
}

func TestMockedGetPriceGrid(t *testing.T) {
	d := &fakes.FakeHTTPDoer{}

	d.DoReturnsOnCall(0, getSetUpResponse(), nil)

	session, err := gflights.New(gflights.WithClient(d))
	if err != nil {
		t.Fatalf("failed to create session: %v", err)
	}

	d.DoReturnsOnCall(1, &http.Response{
		StatusCode: http.StatusOK,
		Body:       getReadCloser(t, "testdata/location/abbr/london.txt"),
	}, nil)

	d.DoReturnsOnCall(2, &http.Response{
		StatusCode: http.StatusOK,
		Body:       getReadCloser(t, "testdata/location/abbr/denver.txt"),
	}, nil)

	d.DoReturnsOnCall(3, &http.Response{
		StatusCode: http.StatusOK,
		Body:       getReadCloser(t, "testdata/price_grid/return.txt"),
	}, nil)

	offers, err := session.GetPriceGrid(
		context.Background(),
		gflights.PriceGridArgs{
			StartDepartureRange: time.Date(2026, time.January, 9, 0, 0, 0, 0, time.UTC),
			EndDepartureRange:   time.Date(2026, time.January, 15, 0, 0, 0, 0, time.UTC),
			StartReturnRange:    time.Date(2026, time.January, 16, 0, 0, 0, 0, time.UTC),
			EndReturnRange:      time.Date(2026, time.January, 22, 0, 0, 0, 0, time.UTC),
			SrcCities:           []string{"London"},
			DstCities:           []string{"Denver"},
			Options: gflights.Options{
				Travelers: gflights.Travelers{Adults: 1},
				Currency:  currency.GBP,
				Class:     gflights.Economy,
				Lang:      language.English,
				TripType:  gflights.RoundTrip,
			},
		},
	)
	if err != nil {
		t.Fatalf("GetPriceGraph failed: %v", err)
	}
	expectedOffersCount := 49
	if len(offers) != expectedOffersCount {
		t.Fatalf("expected %d offers, got %d", expectedOffersCount, len(offers))
	}
}
