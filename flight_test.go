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

func TestGetOneWayOffers(t *testing.T) {
	session, err := gflights.New()
	if err != nil {
		t.Fatal(err)
	}

	offers, _, err := session.GetOutboundOffers(
		context.Background(),
		gflights.Args{
			DepartureDate: time.Now().AddDate(0, 6, 0),
			SrcCities:     []string{"New York"},
			DstAirports:   []string{"LHR"},
			Options: gflights.Options{
				Travelers: gflights.Travelers{Adults: 1},
				Currency:  currency.GBP,
				Stops:     gflights.NonStop,
				Class:     gflights.Economy,
				TripType:  gflights.OneWay,
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
	session, err := gflights.New()
	if err != nil {
		t.Fatal(err)
	}

	args := gflights.Args{
		DepartureDate: time.Now().AddDate(0, 6, 0),
		ReturnDate:    time.Now().AddDate(0, 6, 7),
		SrcCities:     []string{"New York"},
		DstAirports:   []string{"LHR"},
		Options: gflights.Options{
			Travelers: gflights.Travelers{Adults: 1},
			Currency:  currency.GBP,
			Stops:     gflights.AnyStops,
			Class:     gflights.Economy,
			TripType:  gflights.RoundTrip,
			Lang:      language.English,
		},
	}

	offers, _, err := session.GetOutboundOffers(
		context.Background(),
		args,
	)

	if err != nil {
		t.Fatalf("error getting flight offers: %s", err.Error())
	}

	if len(offers) == 0 {
		t.Fatal("expected at least one flight offer, got none")
	}
}

func TestMockedGetOneWayOffers(t *testing.T) {
	d := &fakes.FakeHTTPDoer{}

	d.DoReturnsOnCall(0, &http.Response{
		StatusCode: http.StatusOK,
		Header: http.Header{
			"Set-Cookie": []string{},
		},
	}, nil)

	s, err := gflights.New(gflights.WithClient(d))
	if err != nil {
		t.Fatal(err)
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
		Body:       getReadCloser(t, "testdata/offers/one_way.txt"),
	}, nil)

	offers, priceRange, err := s.GetOutboundOffers(
		context.Background(),
		gflights.Args{
			DepartureDate: time.Date(2026, time.April, 21, 0, 0, 0, 0, time.UTC),
			SrcCities:     []string{"London"},
			DstCities:     []string{"Denver"},
			Options: gflights.Options{
				Travelers: gflights.Travelers{Adults: 1},
				Currency:  currency.GBP,
				Stops:     gflights.AnyStops,
				Class:     gflights.Economy,
				TripType:  gflights.OneWay,
			},
		},
	)

	if err != nil {
		t.Fatalf("error getting flight offers: %s", err.Error())
	}
	if len(offers) == 0 {
		t.Fatal("expected at least one flight offer, got none")
	}
	if priceRange == nil {
		t.Fatal("expected price range, got none")
	}
}

func TestMockedGetRoundTripOffers(t *testing.T) {
	d := &fakes.FakeHTTPDoer{}

	d.DoReturnsOnCall(0, &http.Response{
		StatusCode: http.StatusOK,
		Header: http.Header{
			"Set-Cookie": []string{},
		},
	}, nil)

	s, err := gflights.New(gflights.WithClient(d))
	if err != nil {
		t.Fatal(err)
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
		Body:       getReadCloser(t, "testdata/offers/round_trip_outbound.txt"),
	}, nil)

	offers, priceRange, err := s.GetOutboundOffers(
		context.Background(),
		gflights.Args{
			DepartureDate: time.Date(2026, time.April, 14, 0, 0, 0, 0, time.UTC),
			ReturnDate:    time.Date(2026, time.April, 21, 0, 0, 0, 0, time.UTC),
			SrcCities:     []string{"London"},
			DstCities:     []string{"Denver"},
			Options: gflights.Options{
				Travelers: gflights.Travelers{Adults: 1},
				Currency:  currency.GBP,
				Stops:     gflights.AnyStops,
				Class:     gflights.Economy,
				TripType:  gflights.RoundTrip,
			},
		},
	)

	if err != nil {
		t.Fatalf("error getting flight offers: %s", err.Error())
	}
	if len(offers) == 0 {
		t.Fatal("expected at least one flight offer, got none")
	}
	if priceRange == nil {
		t.Fatal("expected price range, got none")
	}

	var outboundOffer gflights.OutboundOffer

	for _, offer := range offers {
		if offer.Flight[0].FlightCode.FlightNumber == "26" && offer.Flight[0].FlightCode.AirlineCode == "UA" {
			outboundOffer = offer
			break
		}
	}

	if outboundOffer.Flight == nil {
		t.Fatal("expected to find specific flight offer, got none")
	}

	d.DoReturnsOnCall(4, &http.Response{
		StatusCode: http.StatusOK,
		Body:       getReadCloser(t, "testdata/offers/round_trip_return.txt"),
	}, nil)

	returnOffers, err := outboundOffer.GetReturnFlights(context.Background())
	if err != nil {
		t.Fatalf("error getting return flight offers: %s", err.Error())
	}

	if len(returnOffers) != 5 {
		t.Fatal("expected five return flight offers, got ", len(returnOffers))
	}
}
