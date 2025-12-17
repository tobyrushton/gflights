# gflights

This project is a Go client library for the Google Flights API. It makes requests directly to the public Google Flight API.

## Acknowledgements

This package has been based off the work completed in [google-flights-api](https://github.com/krisukox/google-flights-api). While sharing some of the same methods, this package builds on it by adding return flights, price grids and booking urls.

## Installation

```
go get -u github.com/tobyrushton/gflights
```

## Usage

### Session
The session struct contains the majority of the API-related functions required.

```go
session, err := gflights.New()
```

### Outbound Offers
```go
offers, priceRange, err := session.GetOutboundOffer(
    context.Background(),
    gflights.Args{
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
)
```

### Return Offer
```go
// first the outbound offer must be requested as part of a round trip search.
outboundOffer := offers[0]

returnOffers, err := offer.GetReturnFlights(context.Background())
```

### Booking URL
```go
// you must have a trip selected
selectedTrip := outboundOffer.SelectReturnFlight(returnOffer[0])

bookingUrl, err := session.SerialiseBookingURL(
    context.Background(),
    selectedTrip,
)
```

## Notes

This package is still experimental and the API may change at any time.