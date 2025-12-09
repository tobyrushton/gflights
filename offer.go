package gflights

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/tobyrushton/gflights/internal/pb"
	"golang.org/x/text/currency"
	"google.golang.org/protobuf/proto"
)

func serialiseRoundTrip(token string, flights []Flight, price float64, currency currency.Unit) (string, error) {
	flightCodes := func() string {
		out := ""
		for _, flight := range flights {
			out += fmt.Sprintf("%s%s|", flight.FlightCode.AirlineCode, flight.FlightCode.FlightNumber)
		}
		return out[:len(out)-1]
	}()

	roundTrip := pb.RoundTrip{
		Token:        token,
		FlightNumber: flightCodes,
		Price: &pb.RoundTrip_Price{
			Amount:    int32(price) * 100,
			Currency:  currency.String(),
			Precision: 1,
		},
		Unknown:  28,
		UsdPrice: int32(price) * 100, // just set to the same currency for now
	}
	serBytes, err := proto.Marshal(&roundTrip)
	if err != nil {
		return "", fmt.Errorf("could not serialise round trip: %v", err)
	}
	return base64.RawURLEncoding.EncodeToString(serBytes), nil
}

func (o *RoundTripOffer) GetReturnFlights(ctx context.Context) ([]RoundTripFlight, error) {
	// not stricly required for the request.
	serRoundTrip, err := serialiseRoundTrip(o.token, o.Flight, o.Price, o.args.Options.Currency)
	if err != nil {
		return nil, err
	}

	offers, _, _, err := o.s.doGetOffers(ctx, requestFlightArgs{
		rawDataArgs: rawDataArgs{
			Args:            o.args,
			OutboundFlights: o.Flight,
		},
		FlightsSession: serRoundTrip,
	})
	if err != nil {
		return nil, err
	}

	finalOffers := make([]RoundTripFlight, 0, len(offers))
	for _, offer := range offers {
		finalOffers = append(finalOffers, RoundTripFlight{
			Price:  offer.Price,
			Flight: offer.Flight,
		})
	}
	return finalOffers, nil
}
