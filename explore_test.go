package gflights_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/tobyrushton/gflights"
	"github.com/tobyrushton/gflights/internal/fakes"
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
			Stops:     gflights.AnyStops,
			Class:     gflights.Economy,
			TripType:  gflights.RoundTrip,
		},
		Coordinates: gflights.ExploreCoordinates{ // SEA
			NorthLat: 28,
			EastLng:  141,
			SouthLat: -11,
			WestLng:  92,
		},
	}

	offers, err := session.GetExplore(context.Background(), args)
	if err != nil {
		t.Fatal(err)
	}

	if len(offers) == 0 {
		t.Fatal("expected at least one offer")
	}

	// check for some basic airports in SEA to be present in the results
	SEAAirports := map[string]struct{}{
		"KTI": {},
		"BKK": {},
		"HKG": {},
		"SGN": {},
		"KUL": {},
		"DPS": {},
	}

	for _, offer := range offers {
		delete(SEAAirports, offer.AirportCode)
	}

	if len(SEAAirports) > 0 {
		t.Fatalf("expected to find offers for the following SEA airports: %v", SEAAirports)
	}
}

func TestExploreOneWay(t *testing.T) {
	session, err := gflights.New()
	if err != nil {
		t.Fatal(err)
	}

	args := gflights.ExploreArgs{
		DepartureDate: time.Now().Add(time.Hour * 24),
		SrcCities:     []string{"LONDON"},
		Options: gflights.Options{
			Travelers: gflights.Travelers{Adults: 1},
			Currency:  currency.GBP,
			Stops:     gflights.AnyStops,
			Class:     gflights.Economy,
			TripType:  gflights.OneWay,
		},
		Coordinates: gflights.ExploreCoordinates{ // SEA
			NorthLat: 28,
			EastLng:  141,
			SouthLat: -11,
			WestLng:  92,
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

func TestMockedExploreOneWay(t *testing.T) {
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
		Body:       getReadCloser(t, "testdata/explore/one_way.txt"),
	}, nil)

	args := gflights.ExploreArgs{
		DepartureDate: time.Now().Add(time.Hour * 24),
		SrcCities:     []string{"LONDON"},
		Options: gflights.Options{
			Travelers: gflights.Travelers{Adults: 1},
			Currency:  currency.GBP,
			Stops:     gflights.AnyStops,
			Class:     gflights.Economy,
			TripType:  gflights.OneWay,
		},
		Coordinates: gflights.ExploreCoordinates{ // SEA
			NorthLat: 28,
			EastLng:  141,
			SouthLat: -11,
			WestLng:  92,
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

func TestMockedExplore(t *testing.T) {
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
		Body:       getReadCloser(t, "testdata/explore/round_trip.txt"),
	}, nil)

	args := gflights.ExploreArgs{
		DepartureDate: time.Now().Add(time.Hour * 24),
		SrcCities:     []string{"LONDON"},
		Options: gflights.Options{
			Travelers: gflights.Travelers{Adults: 1},
			Currency:  currency.GBP,
			Stops:     gflights.AnyStops,
			Class:     gflights.Economy,
			TripType:  gflights.RoundTrip,
		},
		Coordinates: gflights.ExploreCoordinates{ // SEA
			NorthLat: 28,
			EastLng:  141,
			SouthLat: -11,
			WestLng:  92,
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
