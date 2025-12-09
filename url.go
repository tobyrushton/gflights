package gflights

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/tobyrushton/gflights/internal/pb"
	"google.golang.org/protobuf/proto"
)

func serialiseLocations(locations []string, locationType pb.Url_LocationType) []*pb.Url_Location {
	locationsRet := make([]*pb.Url_Location, 0, len(locations))
	for _, l := range locations {
		locationsRet = append(locationsRet, &pb.Url_Location{
			Type: locationType,
			Name: l,
		})
	}
	return locationsRet
}

func serialiseFlight(
	date time.Time,
	srcCities, srcAirports, dstCities, dstAirports []string,
	stops Stops,
) *pb.Url_Flight {
	return &pb.Url_Flight{
		Date:         date.Format(time.DateOnly),
		SrcLocations: append(serialiseLocations(srcCities, pb.Url_CITY), serialiseLocations(srcAirports, pb.Url_AIRPORT)...),
		DstLocations: append(serialiseLocations(dstCities, pb.Url_CITY), serialiseLocations(dstAirports, pb.Url_AIRPORT)...),
		Stops:        pb.Url_Stops(stops).Enum(),
	}
}

func serialiseFlights(args Args) []*pb.Url_Flight {
	if args.Options.TripType == OneWay {
		return []*pb.Url_Flight{
			serialiseFlight(args.DepartureDate, args.SrcCities, args.SrcAirports, args.DstCities, args.DstAirports, args.Options.Stops),
		}
	}
	return []*pb.Url_Flight{
		serialiseFlight(args.DepartureDate, args.SrcCities, args.SrcAirports, args.DstCities, args.DstAirports, args.Options.Stops),
		serialiseFlight(args.ReturnDate, args.DstCities, args.DstAirports, args.SrcCities, args.SrcAirports, args.Options.Stops),
	}
}

func serialiseTravelers(travelers Travelers) []pb.Url_Traveler {
	travelersRet := make([]pb.Url_Traveler, 0,
		travelers.Adults+travelers.Children+
			travelers.InfantsInSeat+travelers.InfantsOnLap)

	for i := 0; i < travelers.Adults; i++ {
		travelersRet = append(travelersRet, pb.Url_ADULT)
	}
	for i := 0; i < travelers.Children; i++ {
		travelersRet = append(travelersRet, pb.Url_CHILD)
	}
	for i := 0; i < travelers.InfantsInSeat; i++ {
		travelersRet = append(travelersRet, pb.Url_INFANT_IN_SEAT)
	}
	for i := 0; i < travelers.InfantsOnLap; i++ {
		travelersRet = append(travelersRet, pb.Url_INFANT_ON_LAP)
	}
	return travelersRet
}

func (s *Session) SerialiseURL(ctx context.Context, args Args) (string, error) {
	var err error

	if err = args.Validate(); err != nil {
		return "", err
	}

	args.SrcCities, err = s.abbrCities(ctx, args.SrcCities, args.Options.Lang)
	if err != nil {
		return "", fmt.Errorf("could not get abbreviated src cities: %v", err)
	}

	args.DstCities, err = s.abbrCities(ctx, args.DstCities, args.Options.Lang)
	if err != nil {
		return "", fmt.Errorf("could not get abbreviated dst cities: %v", err)
	}
	urlProto := &pb.Url{
		Flight:    serialiseFlights(args),
		Travelers: serialiseTravelers(args.Options.Travelers),
		Class:     pb.Url_Class(args.Options.Class),
		TripType:  pb.Url_TripType(args.Options.TripType),
	}

	tfs, err := proto.Marshal(urlProto)
	if err != nil {
		return "", fmt.Errorf("error during url serialization: %s", err)
	}

	return "https://www.google.com/travel/flights/search" +
		"?tfs=" + base64.RawURLEncoding.EncodeToString(tfs) +
		"&curr=" + args.Options.Currency.String() +
		"&hl=" + args.Options.Lang.String(), nil
}
