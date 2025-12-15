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

func TestPriceGraph(t *testing.T) {
	session, err := gflights.New()
	if err != nil {
		t.Fatalf("failed to create session: %v", err)
	}

	daysDiff1 := 60
	daysDiff2 := 90
	expectedOffersCount := daysDiff2 - daysDiff1 + 1

	offers, err := session.GetPriceGraph(
		context.Background(),
		gflights.PriceGraphArgs{
			RangeStartDate: time.Now().AddDate(0, 0, daysDiff1),
			RangeEndDate:   time.Now().AddDate(0, 0, daysDiff2),
			TripLength:     7,
			SrcCities:      []string{"New York"},
			DstCities:      []string{"London"},
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
	if len(offers) != expectedOffersCount {
		t.Fatalf("expected %d offers, got %d", expectedOffersCount, len(offers))
	}
}

func TestPriceGraphOneWay(t *testing.T) {
	session, err := gflights.New()
	if err != nil {
		t.Fatalf("failed to create session: %v", err)
	}

	daysDiff1 := 60
	daysDiff2 := 90
	expectedOffersCount := daysDiff2 - daysDiff1 + 1

	offers, err := session.GetPriceGraph(
		context.Background(),
		gflights.PriceGraphArgs{
			RangeStartDate: time.Now().AddDate(0, 0, daysDiff1),
			RangeEndDate:   time.Now().AddDate(0, 0, daysDiff2),
			TripLength:     7,
			SrcCities:      []string{"New York"},
			DstCities:      []string{"London"},
			Options: gflights.Options{
				Travelers: gflights.Travelers{Adults: 1},
				Currency:  currency.GBP,
				Class:     gflights.Economy,
				Lang:      language.English,
				TripType:  gflights.OneWay,
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

func TestMockedGetPriceGraph(t *testing.T) {
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
		Body:       getReadCloser(t, "testdata/price_graph/return.txt"),
	}, nil)

	offers, err := session.GetPriceGraph(
		context.Background(),
		gflights.PriceGraphArgs{
			RangeStartDate: time.Date(2026, time.January, 5, 0, 0, 0, 0, time.UTC),
			RangeEndDate:   time.Date(2026, time.March, 5, 0, 0, 0, 0, time.UTC),
			TripLength:     7,
			SrcCities:      []string{"London"},
			DstCities:      []string{"Denver"},
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
	expectedOffersCount := 60
	if len(offers) != expectedOffersCount {
		t.Fatalf("expected %d offers, got %d", expectedOffersCount, len(offers))
	}
}
