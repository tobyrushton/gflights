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
