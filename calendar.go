package gflights

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type calendarArgs struct {
	RangeStartDate time.Time
	RangeEndDate   time.Time
	SrcCities      []string
	SrcAirports    []string
	DstCities      []string
	DstAirports    []string
	Options        Options
}

func (s *Session) getCalendarRawData(ctx context.Context, args calendarArgs) (string, error) {
	serSrcs, err := s.serialiseFlightLocations(ctx, args.SrcCities, args.SrcAirports, args.Options.Lang)
	if err != nil {
		return "", err
	}
	serDsts, err := s.serialiseFlightLocations(ctx, args.DstCities, args.DstAirports, args.Options.Lang)
	if err != nil {
		return "", err
	}

	serDate := args.RangeStartDate.Format("2006-01-02")
	serReturnDate := args.RangeEndDate.Format("2006-01-02")

	serTravelers := serialiseFlightTravelers(args.Options.Travelers)
	serStops := serialiseFlightStop(args.Options.Stops)

	rawData := ""

	rawData += fmt.Sprintf(`[null,null,%d,null,[],%d,%s,null,null,null,null,null,null,[`,
		args.Options.TripType, args.Options.Class, serTravelers)

	rawData += fmt.Sprintf(`[[[%s]],[[%s]],null,%s,null,null,\"%s\",null,null,null,null,null,null,null,3]`,
		serSrcs, serDsts, serStops, serDate)

	if args.Options.TripType == RoundTrip {
		rawData += fmt.Sprintf(`,[[[%s]],[[%s]],null,%s,null,null,\"%s\",null,null,null,null,null,null,null,1]`,
			serDsts, serSrcs, serStops, serReturnDate)
	}

	return rawData, nil
}

func priceCalendarSchema(startDate, returnDate *string, price *float64) *[]any {
	// [startDate,returnDate,[[null,price],""],1]
	return &[]any{startDate, returnDate, &[]any{&[]any{nil, price}}}
}

func getPriceCalendarSection(bytesToDecode []byte) ([]SimpleOffer, error) {
	offers := []SimpleOffer{}

	var err error

	rawOffers := []json.RawMessage{}

	if err = json.Unmarshal([]byte(bytesToDecode), &[]any{nil, &rawOffers}); err != nil {
		return nil, err
	}

	for _, o := range rawOffers {
		finalOffer := SimpleOffer{}

		startDate := ""
		returnDate := ""

		if err = json.Unmarshal(o, priceCalendarSchema(&startDate, &returnDate, &finalOffer.Price)); err != nil {
			continue
		}

		if finalOffer.DepartureDate, err = time.Parse("2006-01-02", startDate); err != nil {
			continue
		}
		if finalOffer.ReturnDate, err = time.Parse("2006-01-02", returnDate); err != nil && returnDate != "" {
			continue
		}
		if finalOffer.Price == 0 || returnDate != "" && finalOffer.ReturnDate.Before(finalOffer.DepartureDate) {
			continue
		}

		offers = append(offers, finalOffer)
	}

	return offers, nil
}
