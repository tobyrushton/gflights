package gflights_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/tobyrushton/gflights"
	"github.com/tobyrushton/gflights/internal/fakes"
	"golang.org/x/text/language"
)

func TestMockedAbbr(t *testing.T) {
	d := &fakes.FakeHTTPDoer{}

	d.DoReturnsOnCall(0, getSetUpResponse(), nil)

	s, err := gflights.New(gflights.WithClient(d))
	if err != nil {
		t.Fatalf("error creating session: %v", err)
	}

	d.DoReturnsOnCall(1, &http.Response{
		StatusCode: http.StatusOK,
		Body:       getReadCloser(t, "testdata/location/abbr/london.txt"),
	}, nil)

	abbr, err := s.AbbrCity(context.Background(), "London", language.English)
	if err != nil {
		t.Fatalf("error getting city abbreviation: %v", err)
	}
	if abbr != "/m/04jpl" {
		t.Fatalf("expected /m/04jpl abbreviation, got %s", abbr)
	}

	d.DoReturnsOnCall(2, &http.Response{
		StatusCode: http.StatusOK,
		Body:       getReadCloser(t, "testdata/location/abbr/denver.txt"),
	}, nil)

	abbr, err = s.AbbrCity(context.Background(), "Denver", language.English)
	if err != nil {
		t.Fatalf("error getting city abbreviation: %v", err)
	}
	if abbr != "/m/02cl1" {
		t.Fatalf("expected /m/02cl1 abbreviation, got %s", abbr)
	}
}

func TestMockedAbbrClosestMatch(t *testing.T) {
	d := &fakes.FakeHTTPDoer{}

	d.DoReturnsOnCall(0, getSetUpResponse(), nil)

	s, err := gflights.New(gflights.WithClient(d))
	if err != nil {
		t.Fatalf("error creating session: %v", err)
	}

	d.DoReturnsOnCall(1, &http.Response{
		StatusCode: http.StatusOK,
		Body:       getReadCloser(t, "testdata/location/abbr/dallas.txt"),
	}, nil)

	// "Dallas-Fort Worth" doesn't match "Dallas" exactly, but the closest
	// result returned by the API is Dallas so it should be accepted.
	abbr, err := s.AbbrCity(context.Background(), "Dallas-Fort Worth", language.English)
	if err != nil {
		t.Fatalf("expected no error for closest match, got: %v", err)
	}
	if abbr != "/m/0f2rq" {
		t.Fatalf("expected /m/0f2rq abbreviation, got %s", abbr)
	}

	// Verify the result is cached under the original input key.
	abbr, err = s.AbbrCity(context.Background(), "Dallas-Fort Worth", language.English)
	if err != nil {
		t.Fatalf("expected cached result, got error: %v", err)
	}
	if abbr != "/m/0f2rq" {
		t.Fatalf("expected /m/0f2rq abbreviation from cache, got %s", abbr)
	}
}
